package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ischeng28/basic-go/webook/internal/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRoutes(server)
	//u.RegisterRoutesV1(server.Group("/users"))
	server.Run(":8080")
}
