package handler

import (
	"context"
	"net/http"

	auth "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/auth/proto/auth"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/plugins/session"
	log "github.com/micro/go-micro/v2/logger"
)

// AuthWrapper 认证wrapper
func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)
		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}

		sess := session.GetSession(w, r)
		if sess.ID != "" {
			// 检测是否通过验证
			if sess.Values["valid"] != nil {
				h.ServeHTTP(w, r)
				return
			} else {
				userId := sess.Values["userId"].(int64)
				if userId != 0 {
					rsp, err := authClient.GetCachedAccessToken(context.TODO(), &auth.Request{
						UserId: userId,
					})
					if err != nil {
						log.Errorf("[AuthWrapper]，err：%s", err)
						http.Error(w, "非法请求", 400)
						return
					}

					// token不一致
					if rsp.Token != ck.Value {
						log.Errorf("[AuthWrapper]，token不一致")
						http.Error(w, "非法请求", 400)
						return
					}
				} else {
					log.Errorf("[AuthWrapper]，session不合法，无用户id")
					http.Error(w, "非法请求", 400)
					return
				}
			}
		} else {
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
