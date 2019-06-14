package orders

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/db"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/proto/orders"
	"github.com/micro/go-log"
)

// GetOrder 获取订单
func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	// 获取数据库
	o := db.GetDB()
	// 查询
	err = o.QueryRow("SELECT id, user_id, book_id, inv_his_id, state FROM orders WHERE id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Logf("[GetOrder] 查询数据失败，err：%s", err)
		return
	}

	return
}
