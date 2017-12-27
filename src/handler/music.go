package handler

import (
	"context"
	"log"
	"strings"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type MusicResult struct {
	Name      string     `json:"name"`
	ID        string     `json:"id"`
	TopTracks []TopTrack `json:"topTracks"`
	ImageURL  string     `json:"imageURL"`
}

type TopTrack struct {
	Name     string `json:"name"`
	URI      string `json:"uri"`
	ImageURL string `json:"imageURL"`
}

func MusicHandler(c *gin.Context) {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	client := spotify.Authenticator{}.NewClient(token)

	artistsStr := c.Param("artists")
	artists := strings.Split(artistsStr, ",")
	var musicResults []MusicResult

	for _, artist := range artists {
		results, err := client.Search(artist, spotify.SearchTypeArtist)
		if err != nil {
			c.JSON(200, gin.H{
				"error": "error finding artist",
			})
		}

		//most relevant search result
		artistID := results.Artists.Artists[0].ID
		relatedArtists, err := client.GetRelatedArtists(artistID)
		if err != nil {
			log.Fatal(err)
		}

		relatedArtists = relatedArtists[:2]
		for _, artist := range relatedArtists {
			topTracks, err := client.GetArtistsTopTracks(artist.ID, spotify.CountryUSA)
			if err != nil {
				log.Fatal(err)
			}
			var trackResults []TopTrack
			topTracks = topTracks[:5]
			for _, topTrack := range topTracks {
				name := topTrack.Name
				URI := string(topTrack.URI)
				imageIndex := len(topTrack.Album.Images)
				imageURL := topTrack.Album.Images[imageIndex-1]
				t := TopTrack{
					Name:     name,
					URI:      URI,
					ImageURL: imageURL.URL,
				}

				trackResults = append(trackResults, t)
			}

			imageLen := len(artist.Images)
			r := MusicResult{
				Name:      artist.Name,
				ID:        string(artist.ID),
				TopTracks: trackResults,
				ImageURL:  artist.Images[imageLen-1].URL,
			}
			musicResults = append(musicResults, r)
		}
	}

	c.JSON(200, musicResults)
}
