package inventory

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/db"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-service/proto/inventory"
	"github.com/micro/go-log"
)

// Sell 销存
func (s *service) Sell(bookId int64, userId int64) (id int64, err error) {

	// 获取数据库
	o := db.GetDB()
	tx, err := o.Begin()
	if err != nil {
		log.Logf("[Sell] 事务开启失败", err.Error())
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	querySQL := `SELECT id, book_id, unit_price, stock, version FROM inventory WHERE book_id = ?`

	inv := &proto.Inv{}

	updateSQL := `UPDATE inventory SET stock = ?, version = ?  WHERE book_id = ? AND version = ?`

	// 销存方法，通过version字段避免脏写
	var minusInv func() error
	minusInv = func() (errIn error) {

		// 查询
		errIn = o.QueryRow(querySQL, bookId).Scan(&inv.Id, &inv.BookId, &inv.UnitPrice, &inv.Stock, &inv.Version)
		if err != nil {
			log.Logf("[Sell] 查询数据失败，err：%s", err)
			return err
		}

		if inv.Stock < 1 {
			log.Logf("[Sell] 库存不足，err：%s", err)
			return err
		}

		r, errIn := o.Exec(updateSQL, inv.Stock-1, inv.Version+1, bookId, inv.Version)
		if errIn != nil {
			log.Logf("[Sell] 更新库存数据失败，err：%s", errIn)
			return
		}

		if affected, _ := r.RowsAffected(); affected == 0 {
			log.Logf("[Sell] 更新库存数据失败，版本号%d过期，即将重试", inv.Version)
			minusInv()
		}

		return
	}

	err = minusInv()
	if err != nil {
		log.Logf("[Sell] 销存失败，err：%s", err)
		return
	}

	insertSQL := `INSERT inventory_history (book_id, user_id, state) VALUE (?, ?, ?) `
	r, err := o.Exec(insertSQL, bookId, userId, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[Sell] 新增销存记录失败，err：%s", err)
		return
	}

	id, _ = r.LastInsertId()

	// 忽略error
	tx.Commit()

	return
}

// Confirm 确认销存
func (s *service) Confirm(id int64, state int) (err error) {

	updateSQL := `UPDATE inventory_history SET state = ? WHERE id = ?;`

	// 获取数据库
	o := db.GetDB()

	// 查询
	_, err = o.Exec(updateSQL, state, id)
	if err != nil {
		log.Logf("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
