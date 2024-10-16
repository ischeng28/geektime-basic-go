package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ischeng28/basic-go/webook/internal/repository"
	"github.com/ischeng28/basic-go/webook/internal/repository/dao"
	"github.com/ischeng28/basic-go/webook/internal/service"
	"github.com/ischeng28/basic-go/webook/internal/web"
	"github.com/ischeng28/basic-go/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()

	server := initWebServer()
	initUserHdl(db, server)

	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

func initWebServer() *gin.Engine {

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//是否允许带cookie之类的数据
		AllowCredentials: true,
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//开发环境
				return true
			}
			return strings.Contains(origin, "cheng.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	useJWT(server)
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
	//	[]byte("secret"), []byte("secret"))
	//if err != nil {
	//	panic(err)
	//}
	//server.Use(sessions.Sessions("ssid", store))
	//
	//server.Use(middleware.NewLoginMiddlewareBuilder().
	//	IgnorePaths("/users/signup").
	//	IgnorePaths("/users/login").Build())
	return server
}

func useJWT(server *gin.Engine) {
	login := middleware.LoginJWTMiddleWareBuilder{}
	server.Use(login.IgnorePaths("/users/signup").IgnorePaths("/users/login").Build())
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// 只会在初始化过程panic
		// panic相当于整个goroutine结束
		// 一旦初始化过程出错 就不要启动了
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
