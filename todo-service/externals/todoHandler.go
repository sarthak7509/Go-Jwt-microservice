package externals

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/go-micro/todo-service/internals/models"
)


func GetAllTodo(ctx *gin.Context){
	todos, err:= models.GetAllTodo()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Oops server error",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, todos)
}

func CreateTodo(ctx *gin.Context){
	
}