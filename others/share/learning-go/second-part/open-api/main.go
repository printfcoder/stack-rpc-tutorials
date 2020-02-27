package main

import (
	"context"
	"encoding/json"
	"strconv"

	proto "github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/api"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/prime"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
)

var (
	sumClient   sum.SumService
	primeClient prime.PrimeService
)

type Open struct {
}

func (open *Open) Fetch(ctx context.Context, req *api.Request, rsp *api.Response) error {
	sumInputStr := req.Get["sum"].Values[0]
	primeInputStr := req.Get["prime"].Values[0]

	sumInput, _ := strconv.ParseInt(sumInputStr, 10, 10)
	sumReq := &sum.SumRequest{
		Input: sumInput,
	}
	sumRsp, err := sumClient.GetSum(context.Background(), sumReq)
	if err != nil {
		panic(err)
	}

	primeInput, _ := strconv.ParseInt(primeInputStr, 10, 10)
	primeReq := &prime.PrimeRequest{
		Input: primeInput,
	}
	primeRsp, err := primeClient.GetPrime(context.Background(), primeReq)
	if err != nil {
		panic(err)
	}

	ret, _ := json.Marshal(map[string]interface{}{
		"sum":   sumRsp.Output,
		"prime": primeRsp.Output,
	})

	rsp.Body = string(ret)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.api.open"),
	)

	service.Init(micro.Action(func(c *cli.Context) error {
		sumClient = sum.NewSumService("go.micro.learning.srv.sum", service.Client())
		primeClient = prime.NewPrimeService("go.micro.learning.srv.prime", service.Client())
		return nil
	}))

	_ = proto.RegisterOpenHandler(service.Server(), new(Open))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
