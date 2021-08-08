package main

import (
	"dbLazy/conndb"
	"dbLazy/service"

	"github.com/gin-gonic/gin"
)

var (
	db      = conndb.Connection_db()
	handler = service.NewHandler(db)
)

func main() {
	r := gin.Default()

	db.AutoMigrate(&service.Product{})

	r.POST("/api/products", handler.InsertData)
	r.GET("/api/products/frontend", handler.GetAllData)
	r.GET("/api/products/backend", handler.GetQuery)

	r.Run()
}
