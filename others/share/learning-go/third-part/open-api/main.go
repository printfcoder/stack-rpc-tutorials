package main

import (
	"context"
	"encoding/json"

	proto "github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/api"
	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
)

type Open struct {
}

func (open *Open) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	name := req.Get["name"].Values[0]

	ret, _ := json.Marshal(map[string]interface{}{
		"Hello! ": name,
	})

	rsp.Body = string(ret)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.api.open"),
	)

	_ = proto.RegisterOpenHandler(service.Server(), new(Open))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
