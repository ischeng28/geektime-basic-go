package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ischeng28/basic-go/webook/internal/web"
	ijwt "github.com/ischeng28/basic-go/webook/internal/web/jwt"
	"github.com/ischeng28/basic-go/webook/internal/web/middleware"
	"github.com/ischeng28/basic-go/webook/pkg/ginx/middleware/prometheus"
	"github.com/ischeng28/basic-go/webook/pkg/logger"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc,
	artHdl *web.ArticleHandler,
	userHdl *web.UserHandler,
	wechatHdl *web.OAuth2WechatHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	wechatHdl.RegisterRoutes(server)
	artHdl.RegisterRoutes(server)
	userHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares(redisClient redis.Cmdable, hdl ijwt.Handler, l logger.LoggerV1) []gin.HandlerFunc {
	pb := &prometheus.Builder{
		Namespace: "geektime_daming",
		Subsystem: "webook",
		Name:      "gin_http",
		Help:      "统计 GIN 的HTTP接口数据",
	}
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
			ExposeHeaders: []string{"x-jwt-token", "x-refresh-token"},
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
		pb.BuildResponseTime(),
		pb.BuildActiveRequest(),
		//ratelimit.NewBuilder(limiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 1000)).Build(),
		//middleware.NewLogMiddlewareBuilder(func(ctx context.Context, al middleware.AccessLog) {
		//	l.Debug("", logger.Field{Key: "req", Val: al})
		//}).AllowReqBody().AllowRespBody().Build(),
		middleware.NewLoginJWTMiddlewareBuilder(hdl).CheckLogin(),
	}
}
