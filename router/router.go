package router

import "github.com/gin-gonic/gin"

// Router主要负责初始化pool注册池子
func RunRouter() *gin.Engine {
	r := gin.Default()
	r.Group("/api")
	return r
}
