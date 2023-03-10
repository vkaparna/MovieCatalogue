package routes

import (
	"github.com/Kathan-Vakharia/movie-catalog/metadata/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/metadata", controllers.GetMoviesMetadata)
	router.GET("/metadata/:id", controllers.GetMovieMetadataByID)
	router.GET("metadata/title/:title", controllers.GetMovieMetadataByTitle)
}
