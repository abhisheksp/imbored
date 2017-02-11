package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
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
	artistsStr := c.Param("artists")

	artists := strings.Split(artistsStr, ",")
	musicResults := []MusicResult{}

	for _, artist := range artists {
		results, err := spotify.Search(artist, spotify.SearchTypeArtist)
		if err != nil {
			c.JSON(200, gin.H{
				"error": "error finding artist",
			})
		}

		//most relevant search result
		artistID := results.Artists.Artists[0].ID

		relatedArtists, err := spotify.GetRelatedArtists(artistID)
		if err != nil {
			log.Fatal(err)
		}

		relatedArtists = relatedArtists[:2]
		for _, artist := range relatedArtists {
			topTracks, err := spotify.GetArtistsTopTracks(artist.ID, spotify.CountryUSA)
			if err != nil {
				log.Fatal(err)
			}
			trackResults := []TopTrack{}
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
