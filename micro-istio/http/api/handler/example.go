package handler

import (
	"context"
	"encoding/json"

	apiClient "github.com/micro-in-cn/tutorials/micro-istio/http/api/client"
	example "github.com/micro-in-cn/tutorials/micro-istio/http/srv/proto/example"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
)

type Example struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Call is called by the API as /http/example/call with post body {"name": "foo"}
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	exampleClient, ok := apiClient.ExampleFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.sample.example.call", "example client not found")
	}

	// make request
	log.Logf("post data:%v", req.GetPost())
	r := example.Request{}
	if req.Method == "POST" {
		if err := json.Unmarshal([]byte(req.Body), &r); err != nil {
			return errors.InternalServerError("go.micro.api.sample.example.call", err.Error())
		}

	} else {
		name := extractValue(req.Get["name"])
		r.Name = name
	}
	response, err := exampleClient.Call(ctx, &r)

	if err != nil {
		return errors.InternalServerError("go.micro.api.sample.example.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
