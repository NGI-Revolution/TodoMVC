package main

import (
	"project/todomvc/main.go/db"
	"project/todomvc/main.go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitInt()
	//db.DB.AutoMigrate(&model.Todomvc{})
	r := gin.Default()
	InitRouter(r)

	r.Run(":8000")
}

func InitRouter(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.POST("add", handler.Add)
	v1.POST("delete", handler.Delete)
	v1.POST("update", handler.Update)
	v1.POST("find", handler.Find)
}
