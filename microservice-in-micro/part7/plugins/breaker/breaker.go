package breaker

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	statusCode "github.com/micro-in-cn/tutorials/microservice-in-micro/part7/plugins/breaker/http"
)

//BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		hystrix.Do(name, func() error {
			sct := &statusCode.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)

			if sct.Status >= http.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}, func(e error) error {
			if e == hystrix.ErrCircuitOpen {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("请稍后重试"))
			}

			return e
		})
	})
}
