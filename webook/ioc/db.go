package ioc

import (
	"github.com/ischeng28/basic-go/webook/config"
	"github.com/ischeng28/basic-go/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
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
