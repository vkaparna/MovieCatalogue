package routes

import (
	"github.com/Kathan-Vakharia/movie-catalog/ratings/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("ratings/", controllers.GetMovieRatings)
	router.GET("/ratings/:id", controllers.GetMovieRatingsByID)

}
