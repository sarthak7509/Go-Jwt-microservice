package routers

import "github.com/gin-gonic/gin"

func Router(app *gin.Engine) {
	app.GET("/health", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "healthy"}) })
}
