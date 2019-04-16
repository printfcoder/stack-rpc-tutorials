package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/config"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"
)

// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// standardClaims token 标准的Claims
type standardClaims struct {
	SubjectID string `json:"subjectId,omitempty"`
	Name      string `json:"name,omitempty"`
	jwt.StandardClaims
}

// MakeAccessToken 生成token并保存到redis
func (s *service) MakeAccessToken(subject *Subject) (ret string, err error) {

	m, err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	// 创建
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// GetCachedAccessToken 获取token
func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}
