package model

import (
	"context"

	"github.com/micro-in-cn/tutorials/micro-sync/common"
	invS "github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/proto/inventory"
	"github.com/micro/go-micro/util/log"
)

// New 新增订单
func (s *service) New(bookId int64, userId int64) (orderId int64, err error) {
	// 请求销存
	rsp, err := invClient.Sell(context.TODO(), &invS.Request{
		BookId: bookId, UserId: userId,
	})
	if err != nil {
		log.Logf("[ERR] [New] Sell 调用库存服务时失败：%s", err.Error())
		return
	}

	// 获取数据库
	insertSQL := `INSERT orders (user_id, book_id, inv_his_id, state) VALUE (?, ?, ?, ?)`

	r, err := db.Exec(insertSQL, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[ERR] [New] 新增订单失败，err：%s", err)
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

// UpdateOrderState 更新订单状态
func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSQL := `UPDATE orders SET state = ? WHERE id = ?`

	// 更新
	_, err = db.Exec(updateSQL, state, orderId)
	if err != nil {
		log.Logf("[ERR] [Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
