package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// createTokenClaims Claims
func (s *service) createTokenClaims(subject *Subject) (m *standardClaims, err error) {

	now := time.Now()
	m = &standardClaims{
		SubjectID: subject.ID,
		Name:      subject.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(tokenExpiredDate).Unix(),
			NotBefore: now.Unix(),
			Id:        subject.ID,
			IssuedAt:  now.Unix(),
			Issuer:    "book.micro.mu",
			Subject:   subject.ID,
		},
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

// clearTokenFromCache 清空token
func (s *service) clearTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("[clearTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
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
