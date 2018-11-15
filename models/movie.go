package models
import (
	"time"
	"github.com/jinzhu/gorm"
)
type MovieModel struct{}
//Movie ...
type Movie struct {
	gorm.Model
	UserID    int   
	Title     string `gorm:"UNIQUE"`
	Year   		int
	Rated 		string
	Released 	time.Time
	Genre 		string
	Director 	string
	Writer 		string
	Plot 			string
	Language 	string
	Country 	string
	Awards 		string
	Poster 		string
	Ratings 	JSONRaw 
	Metascore float32
	ImdbRating float32
	imdbVotes int
	imdbID	 	string
	Type 			string
	DVD  			string
	BoxOffice string
	Production string
	Website 	string
}
