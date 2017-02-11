package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanbradynd05/go-tmdb"
	"log"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

const (
	APIKey        = "82bb8117f63f31501095596db0260115"
	ImageBasePath = "http://image.tmdb.org/t/p/w300/"
)

type MovieResult struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	ImageURL string `json:"imageURL"`
}

func MovieHandler(c *gin.Context) {
	t := tmdb.Init(APIKey)

	movieStr := c.Param("movies")
	movies := strings.Split(movieStr, ",")

	movieResults := []MovieResult{}
	for _, m := range movies {

		//Get Liked Movie ID
		options := map[string]string{
			"page":          "1",
			"include_adult": "false",
		}
		likedMovie, err := t.SearchMovie(m, options)
		if err != nil {
			log.Fatalf("error searching movie: %s", err)
		}

		//most relevant result
		movieID := likedMovie.Results[0].ID

		options = map[string]string{"page": "1"}
		recommendedMovies, err := t.GetMovieRecommended(movieID, options)
		if err != nil {
			log.Fatalf("error recommending movies: %s", err)
		}

		reducedResults := recommendedMovies.Results[:6]
		for _, movie := range reducedResults {
			m := MovieResult{
				ID:       strconv.Itoa(movie.ID),
				Name:     movie.OriginalTitle,
				ImageURL: ImageBasePath + movie.PosterPath,
			}

			movieResults = append(movieResults, m)
		}
	}

	rand.Seed(time.Now().UnixNano())
	Shuffle(movieResults)
	c.JSON(200, movieResults)
}

func Shuffle(a []MovieResult) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}