package orders

import (
	"context"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/db"
	"github.com/micro/go-log"

	invS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/proto/inventory"
)

// New 新增订单
func (s *service) New(bookId int64, userId int64) (orderId int64, err error) {

	// 请求销存
	rsp, err := invClient.Sell(context.TODO(), &invS.Request{
		BookId: bookId, UserId: userId,
	})
	if err != nil {
		log.Logf("[New] Sell 调用库存服务时失败：%s", err.Error())
		return
	}

	// 获取数据库
	o := db.GetDB()
	insertSQL := `INSERT orders (user_id, book_id, inv_his_id, state) VALUE (?, ?, ?, ?)`

	r, err := o.Exec(insertSQL, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[New] 新增订单失败，err：%s", err)
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

// UpdateOrderState 更新订单状态
func (s *service) UpdateOrderState(orderId int64, state int) (err error) {

	updateSQL := `UPDATE orders SET state = ? WHERE id = ?`

	// 获取数据库
	o := db.GetDB()
	// 更新
	_, err = o.Exec(updateSQL, state, orderId)
	if err != nil {
		log.Logf("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
