package externals

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/go-micro/todo-service/internals/models"
)

func GetAllTodo(ctx *gin.Context) {
	todos, err := models.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops server error",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, todos)
}

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Ooops error occured"})
		log.Println(err)
		return
	}

	err = todo.Save()

	if err != nil {
		ctx.JSON(400, gin.H{"message": "Ooops error occured"})
		log.Println(err)
		return
	}

	ctx.JSON(201, todo)
}
