package main

import (
	"github.com/wahyuharyo/prakerja/controllers"

	"github.com/wahyuharyo/prakerja/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	r.GET("/api/pesanan", controllers.Index)
	r.GET("/api/pesanan/:id", controllers.Show)
	r.POST("/api/pesanan", controllers.Create)
	r.PUT("/api/pesanan/:id", controllers.Update)
	r.DELETE("/api/pesanan", controllers.Delete)

	r.Run(":9090")
}