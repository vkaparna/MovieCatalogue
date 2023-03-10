package models

import (
	"github.com/Kathan-Vakharia/movie-catalog/metadata/pkg/config"
	"github.com/jmoiron/sqlx"
)

type Metadata struct {
	ID          int    `db:"ID" json:"id"`
	API_ID      string `db:"API_ID" json:"api_id"`
	Title       string `db:"Title" json:"title"`
	PosterPath  string `db:"Poster_Path" json:"poster_path"`
	Description string `db:"Description" json:"description"`
	ReleaseDate string `db:"Release_Date" json:"release_date"`
}

var db *sqlx.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllMetadata() []Metadata {
	var moviesmetadata = []Metadata{}

	err := db.Select(&moviesmetadata, "SELECT * FROM metadata")
	if err != nil {
		panic(err)
	}
	return moviesmetadata

}
func GetMetaDataById(id string) (*Metadata, *sqlx.DB) {
	var metadata Metadata

	err := db.Get(&metadata, "SELECT * FROM metadata WHERE ID=$1", id)
	if err != nil {
		panic(err)
	}
	return &metadata, db
}

func GetMetaDataByTitle(title string) (*Metadata, *sqlx.DB) {
	var metadata Metadata

	err := db.Get(&metadata, "SELECT * FROM metadata WHERE title=$1", title)

	if err != nil {
		panic(err)
	}
	return &metadata, db
}
