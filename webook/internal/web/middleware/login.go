package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (m *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range m.paths {
			if path == ctx.Request.URL.Path {
				return
			}
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			// 中断，不要往后执行，也就是不要执行后面的业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			MaxAge: 60,
		})
		now := time.Now().UnixMilli()

		// 说明还没有刷新过 刚登陆还没刷新
		if updateTime == nil {
			sess.Set("update_time", now)
			sess.Save()
			return
		}

		updateTimeVal, _ := updateTime.(int64)
		if now-updateTimeVal > 60*1000 {
			sess.Set("update_time", now)
			sess.Save()
			return
		}

	}
}
