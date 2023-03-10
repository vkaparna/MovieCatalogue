package routes

import (
	"github.com/Kathan-Vakharia/movie-catalog/movie/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/movie", controllers.GetMovies)
	router.GET("/movie/:id", controllers.GetMovieByID)
	// router.GET("metadata/title/:title", controllers.GetMovieMetadataByTitle)
}
