package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ischeng28/basic-go/webook/internal/web"
	"github.com/ischeng28/basic-go/webook/internal/web/middleware"
	"github.com/ischeng28/basic-go/webook/pkg/ginx/middleware/ratelimit"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler, wechatHdl *web.OAuth2WechatHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	wechatHdl.RegisterRoutes(server)
	userHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			println("这是我的 Middleware")
		},
		cors.New(cors.Config{
			//AllowAllOrigins:  true,
			AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
			AllowCredentials: true,

			AllowHeaders: []string{"Content-Type", "Authorization"},
			// 这个是允许前端访问你的后端响应中带的头部
			ExposeHeaders: []string{"x-jwt-token"},
			//AllowHeaders: []string{"content-type"},
			//AllowMethods: []string{"POST"},
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					//if strings.Contains(origin, "localhost") {
					return true
				}
				return strings.Contains(origin, "live.webook.com")
			},
			MaxAge: 12 * time.Hour,
		}),
		ratelimit.NewBuilder(redisClient, time.Second, 2).Build(),
		middleware.NewLoginJWTMiddleWareBuilder().IgnorePaths([]string{"/users/signup", "/users/login", "/hello", "/oauth2/wechat/authurl", "/oauth2/wechat/callback"}).Build(),
	}
}
