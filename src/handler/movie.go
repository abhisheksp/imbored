package handler

import (
	"log"
	"strconv"
	"strings"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryanbradynd05/go-tmdb"
)

const (
	ImageBasePath = "http://image.tmdb.org/t/p/w300/"
)

type MovieResult struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	ImageURL string `json:"imageURL"`
}

func MovieHandler(c *gin.Context) {
	APIKey := os.Getenv("TMDB_API_KEY")
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
		recommendedMovies, err := t.GetMovieSimilar(movieID, options)
		if err != nil {
			log.Fatalf("error recommending movies: %s", err)
		}

		if len(recommendedMovies.Results) == 0 {
			c.JSON(200, "no recommendations founds")
			return
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
