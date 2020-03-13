package learning

import (
	"context"
	"fmt"

	"github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/learning"
)

type Handler struct {
}

func (h *Handler) Hi(ctx context.Context, request *learning.Request, response *learning.Response) error {
	response.Msg = fmt.Sprintf("Hello! %s. Welcome to Learning!", request.Name)
	return nil
}
