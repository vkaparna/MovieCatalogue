package models

import (
	"github.com/Kathan-Vakharia/movie-catalog/ratings/pkg/config"
	"github.com/jmoiron/sqlx"
)

type Ratings struct {
	User_ID      int `db:"User_ID" json:"user_id"`
	Movie_ID     int `db:"Movie_ID" json:"movie_id"`
	Rating_Value int `db:"Rating_Value" json:"rating_value"`
}

type AggregatedRating struct {
	Movie_ID     int     `db:"Movie_ID" json:"movie_id"`
	Rating_Value float64 `db:"Rating_Value" json:"rating_value"`
}

var db *sqlx.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetRatingsById(id string) (*AggregatedRating, *sqlx.DB) {
	//var ratings Ratings
	var aggRating AggregatedRating

	err := db.Get(&aggRating, "SELECT Movie_ID, avg(Rating_Value) as Rating_Value FROM ratings WHERE Movie_ID=$1", id)
	if err != nil {
		panic(err)
	}
	return &aggRating, db
}

func GetRatings() []AggregatedRating {
	var ratings = []AggregatedRating{}

	err := db.Select(&ratings, "SELECT MOVIE_ID, avg(Rating_Value) as Rating_Value FROM ratings GROUP BY MOVIE_ID")
	if err != nil {
		panic(err)
	}
	return ratings
}
