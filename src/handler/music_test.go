package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

)

func TestMusicHandler(t *testing.T) {
	r := gin.New()
	r.POST("/music/:artist", MusicHandler)

	req, err := http.NewRequest("POST", "/music/linkin park", nil)
	require.NoError(t, err, "no error expected creating request")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	expected := `{"music":"linkin park"}` + "\n"

	assert.Equal(t, res.Body.String(), expected)
}