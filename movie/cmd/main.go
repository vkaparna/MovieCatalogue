package main

import (
	"github.com/Kathan-Vakharia/movie-catalog/movie/pkg/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// func getMetaDatas(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, )
// }
// func getTodo(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)

// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not found"})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, todo)
// }

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run("localhost:9092")

}
