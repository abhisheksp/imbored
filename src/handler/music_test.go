package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestMusicHandler(t *testing.T) {
	r := gin.New()
	r.POST("/music/:artists", MusicHandler)

	req, err := http.NewRequest("POST", "/music/linkin park, metallica", nil)
	require.NoError(t, err, "no error expected creating request")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	status := res.Code
	assert.Equal(t, http.StatusOK, status)
}
