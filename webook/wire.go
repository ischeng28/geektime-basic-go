//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ischeng28/basic-go/webook/internal/repository"
	"github.com/ischeng28/basic-go/webook/internal/repository/cache"
	"github.com/ischeng28/basic-go/webook/internal/repository/dao"
	"github.com/ischeng28/basic-go/webook/internal/service"
	"github.com/ischeng28/basic-go/webook/internal/web"
	ijwt "github.com/ischeng28/basic-go/webook/internal/web/jwt"
	"github.com/ischeng28/basic-go/webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitRedis, ioc.InitDB,
		// DAO 部分
		dao.NewUserDAO,

		// cache 部分
		//cache.NewCodeCache,
		cache.NewUserCache,

		// repository 部分
		repository.NewUserRepository,
		//repository.NewCodeRepository,

		// Service 部分
		//ioc.InitSMSService,
		service.NewUserService,

		//service.NewCodeService,

		// handler 部分
		web.NewUserHandler,
		web.NewOAuth2WechatHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		ijwt.NewRedisJWTHandler,
		ioc.InitWechatService,
	)
	return gin.Default()
}
