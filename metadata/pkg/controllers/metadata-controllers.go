package controllers

import (
	"net/http"

	"github.com/Kathan-Vakharia/movie-catalog/metadata/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetMoviesMetadata(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, models.GetAllMetadata())
}
func GetMovieMetadataByID(context *gin.Context) {
	id := context.Param("id")
	metadata, _ := models.GetMetaDataById(id)
	context.IndentedJSON(http.StatusOK, metadata)
}

func GetMovieMetadataByTitle(context *gin.Context) {
	title := context.Param("title")
	metadata, _ := models.GetMetaDataByTitle(title)
	context.IndentedJSON(http.StatusOK, metadata)
}
