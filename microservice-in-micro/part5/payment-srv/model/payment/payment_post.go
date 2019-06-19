package payment

import (
	"context"
	"fmt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic/common"
	invS "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/inventory-srv/proto/inventory"
	ordS "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/orders-srv/proto/orders"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/db"
	"github.com/micro/go-micro/util/log"
)

// PayOrder 支付订单
func (s *service) PayOrder(orderId int64) (err error) {
	// 获取支付单
	orderRsp, err := ordSClient.GetOrder(context.TODO(), &ordS.Request{
		OrderId: orderId,
	})
	if err != nil {
		log.Logf("[PayOrder] 查询 订单信息失败，orderId：%d, %s", orderId, err)
		return
	}

	// 订单不存在
	if orderRsp == nil || !orderRsp.Success || orderRsp.Order == nil {
		err = fmt.Errorf("[PayOrder] 支付单不存在")
		log.Logf("[PayOrder] 查询 订单信息失败，orderId：%d, %s", orderId, err)
		return
	}

	// 订单已支付
	if orderRsp.Order.State == common.InventoryHistoryStateOut {
		err = fmt.Errorf("[PayOrder] 订单已支付")
		log.Logf("[PayOrder] 查询 订单已支付，orderId：%d", orderId)
		return
	}

	// 获取数据库并开启事务
	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Logf("[PayOrder] 事务开启失败", err.Error())
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 插入新记录
	insertSQL := `INSERT INTO payment (user_id, book_id, order_id, inv_his_id, state) VALUE (?, ?, ?, ?, ?)`
	_, err = tx.Exec(insertSQL, orderRsp.Order.UserId, orderRsp.Order.BookId, orderRsp.Order.Id, orderRsp.Order.InvHistoryId, common.InventoryHistoryStateOut)
	if err != nil {
		log.Logf("[New] 新增支付单失败，%v, err：%s", orderRsp.Order, err)
		return
	}

	// 确认出库
	invRsp, err := invClient.Confirm(context.TODO(), &invS.Request{
		HistoryId: orderRsp.Order.InvHistoryId,
	})
	if err != nil || invRsp == nil || !invRsp.Success {
		err = fmt.Errorf("[PayOrder] 确认出库失败，%s", err)
		log.Logf("%s", err)
		return
	}

	// 广播支付成功
	s.sendPayDoneEvt(orderId, common.InventoryHistoryStateOut)

	tx.Commit()
	return
}
