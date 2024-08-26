package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.Use(func(context *gin.Context) {
		println("这是第一个middleware")
	}, func(context *gin.Context) {
		println("这是第二个middleware")
	})

	server.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	// 参数路由 路由参数
	server.GET("user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello,"+name)
	})

	// 查询参数
	// GET /order?id=123
	server.GET("order", func(context *gin.Context) {
		id := context.Query("id")
		context.String(http.StatusOK, "订单是"+id)
	})

	// 通配符匹配
	server.GET("/views/*.html", func(context *gin.Context) {
		path := context.Param(".html")
		context.String(http.StatusOK, "匹配上的值是%s", path)
	})

	server.POST("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "hello login")
	})

	server.Run(":8080")
}
