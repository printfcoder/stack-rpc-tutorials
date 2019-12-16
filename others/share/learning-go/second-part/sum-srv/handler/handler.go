package handler

import (
	"context"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/sum-srv/service"
)

type handler struct {
}

func (h handler) GetSum(ctx context.Context, req *sum.SumRequest, rsp *sum.SumResponse) error {
	inputs := make([]int64, 0)
	var i int64 = 0
	for ; i <= req.Input; i++ {
		inputs = append(inputs, i)
	}

	rsp.Output = service.GetSum(inputs...)

	return nil
}

func Handler() sum.SumHandler {
	return handler{}
}
