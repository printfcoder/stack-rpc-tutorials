package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/db"
	invS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/proto/inventory"
	ordS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/proto/orders"
	payS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/payment"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/broker"
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
	_, err = tx.Exec(insertSQL, orderRsp.Order.BookId, orderRsp.Order.Id, orderRsp.Order.InvHistoryId, common.InventoryHistoryStateOut)
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
	body, _ := json.Marshal(&payS.Payments{
		OrderId: orderId,
		State:   common.InventoryHistoryStateOut,
	})
	msg := &broker.Message{
		Header: map[string]string{},
		Body:   body,
	}
	broker.Publish(common.TopicPaymentDone, msg)

	tx.Commit()

	return
}
