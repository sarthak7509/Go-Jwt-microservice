package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/go-micro/todo-service/internals/database"
	"github.com/sarthak7509/go-micro/todo-service/routers"
)

func main() {
	srv := gin.Default()
	database.InitDb()
	routers.Router(srv)
	srv.Run(":8080")
}
