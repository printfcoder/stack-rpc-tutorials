package handler

import (
	"context"

	"github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/greeter"
)

type Handler struct {
}

func (h *Handler) Hi(ctx context.Context, request *greeter.Request, response *greeter.Response) error {
	response.Msg = "Hello! " + request.Name
	return nil
}
