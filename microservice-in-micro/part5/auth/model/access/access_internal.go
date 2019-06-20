package access

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// createTokenClaims Claims
func (s *service) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "book.micro.mu",
		Subject:   subject.ID,
	}

	return
}

// saveTokenToCache 保存token到缓存
func (s *service) saveTokenToCache(subject *Subject, val string) (err error) {
	//保存
	if err = ca.Set(tokenIDKeyPrefix+subject.ID, val, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveTokenToCache] 保存token到缓存发生错误，err:" + err.Error())
	}
	return
}

// delTokenFromCache 清空token
func (s *service) delTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("[delTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
	}
	return
}

// getTokenFromCache 从缓存获取token
func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {
	// 获取
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.ID).Result()
	if err != nil {
		return token, fmt.Errorf("[getTokenFromCache] token不存在 %s", err)
	}

	return string(tokenCached), nil
}

// parseToken 解析token
func (s *service) parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})

	// jwt 框架自带了一些检测，如过期，发布者错误等
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
	}

	// 检测合法
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] 不合法的token")
	}

	return mapClaimToJwClaim(claims), nil
}

// 把jwt的claim转成claims
func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}

	return jC
}
