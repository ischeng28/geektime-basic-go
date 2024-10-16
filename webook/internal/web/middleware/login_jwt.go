package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ischeng28/basic-go/webook/internal/web"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddleWareBuilder struct {
	paths []string
}

func NewLoginJWTMiddleWareBuilder() *LoginJWTMiddleWareBuilder {
	return &LoginJWTMiddleWareBuilder{}
}

func (l *LoginJWTMiddleWareBuilder) IgnorePaths(path string) *LoginJWTMiddleWareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (m *LoginJWTMiddleWareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range m.paths {
			if path == ctx.Request.URL.Path {
				return
			}
		}
		// 根据约定，token 在 Authorization 头部
		// Bearer XXXX
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			// 没登录，没有 token, Authorization 这个头部都没有
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			// 没登录，Authorization 中的内容是乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			//	后期我们讲到了监控告警的时候 这个地方要埋点
			//	能够进来这个分支的 大概率是攻击者
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		expireTime := uc.ExpiresAt
		//	剩余时间小于50s就要刷新
		if expireTime.Sub(time.Now()) < time.Second*50 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 10))
			tokenStr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				//	这边不要中断 因为仅仅是过期时间没有刷新 用户是登录了的
				log.Println(err)
			}
		}
		ctx.Set("user", uc)
	}
}
