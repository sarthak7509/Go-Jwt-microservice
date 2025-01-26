package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/go-micro/todo-service/externals"
)

func Router(app *gin.Engine) {
	app.GET("/health", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "healthy"}) })

	//Todo api's it 
	app.GET("/todos", externals.GetAllTodo)
}
