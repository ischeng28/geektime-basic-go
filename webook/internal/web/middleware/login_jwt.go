package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	jwt2 "github.com/ischeng28/basic-go/webook/internal/web/jwt"
	"net/http"
)

type LoginJWTMiddleWareBuilder struct {
	paths   []string
	handler jwt2.Handler
}

func NewLoginJWTMiddleWareBuilder(hdl jwt2.Handler) *LoginJWTMiddleWareBuilder {
	return &LoginJWTMiddleWareBuilder{
		paths:   []string{"/users/signup", "/users/login", "/hello", "/oauth2/wechat/authurl", "/oauth2/wechat/callback"},
		handler: hdl,
	}
}

func (l *LoginJWTMiddleWareBuilder) IgnorePaths(paths []string) *LoginJWTMiddleWareBuilder {
	l.paths = append(l.paths, paths...)
	return l
}

func (m *LoginJWTMiddleWareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range m.paths {
			if path == ctx.Request.URL.Path {
				return
			}
		}
		extractToken := m.handler.ExtractToken(ctx)
		var uc jwt2.UserClaims
		token, err := jwt.ParseWithClaims(extractToken, &uc, func(token *jwt.Token) (interface{}, error) {
			return jwt2.JWTKey, nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			// 在这里发现 access_token 过期了，生成一个新的 access_token

			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 这里看
		err = m.handler.CheckSession(ctx, uc.Ssid)
		if err != nil {
			// token 无效或者 redis 有问题
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", uc)
	}
}
