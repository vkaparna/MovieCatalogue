package main

import (
	"github.com/Kathan-Vakharia/movie-catalog/ratings/pkg/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run("localhost:8091")
}
