package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	metadata_path = "http://localhost:8090"
	ratings_path  = "http://localhost:8091"
)

type Metadata struct {
	ID          int    `db:"ID" json:"id"`
	API_ID      string `db:"API_ID" json:"api_id"`
	Title       string `db:"Title" json:"title"`
	PosterPath  string `db:"Poster_Path" json:"poster_path"`
	Description string `db:"Description" json:"description"`
	ReleaseDate string `db:"Release_Date" json:"release_date"`
}
type AggregatedRating struct {
	Movie_ID     int     `db:"Movie_ID" json:"movie_id"`
	Rating_Value float64 `db:"Rating_Value" json:"rating_value"`
}

type Movie struct {
	Metadata Metadata         `json:"metadata"`
	Rating   AggregatedRating `json:"agg_rating"`
}

func GetMovies(context *gin.Context) {
	//var movies []Movie

	context.IndentedJSON(http.StatusOK, nil)

}
func GetMovieByID(context *gin.Context) {
	id := context.Param("id")
	//collecting metadata
	resp, _ := http.Get(metadata_path + "/metadata/" + id)
	body, _ := ioutil.ReadAll(resp.Body)
	var metadata Metadata
	json.Unmarshal(body, &metadata)

	resp1, _ := http.Get(ratings_path + "/ratings/" + id)
	body, _ = ioutil.ReadAll(resp1.Body)
	var rating AggregatedRating
	json.Unmarshal(body, &rating)

	var movie = Movie{Metadata: metadata, Rating: rating}

	context.IndentedJSON(http.StatusOK, movie)

}

// func GetMovieMetadataByTitle(context *gin.Context) {

// }
