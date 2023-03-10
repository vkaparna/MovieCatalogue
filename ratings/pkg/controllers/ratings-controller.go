package controllers

import (
	"net/http"

	"github.com/Kathan-Vakharia/movie-catalog/ratings/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetMovieRatingsByID(context *gin.Context) {
	id := context.Param("id")
	rating, _ := models.GetRatingsById(id)
	context.IndentedJSON(http.StatusOK, rating)
}

func GetMovieRatings(context *gin.Context) {
	ratings := models.GetRatings()
	context.IndentedJSON(http.StatusOK, ratings)
}
