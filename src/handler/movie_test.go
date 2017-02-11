package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMovieHandler(t *testing.T) {
	r := gin.New()
	r.POST("/movies/:movies", MovieHandler)

	req, err := http.NewRequest("POST", "/movies/interstellar", nil)
	require.NoError(t, err, "no error expected creating request")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
}
