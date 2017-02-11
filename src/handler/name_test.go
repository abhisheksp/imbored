package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

)

func TestTestHandler(t *testing.T) {
	r := gin.New()
	r.POST("/testpath/:testparam", TestHandler)

	req, err := http.NewRequest("POST", "/testpath/exampleparam", nil)
	require.NoError(t, err, "no error expected creating request")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	expected := `{"param":"exampleparam","status":"posted"}` + "\n"

	assert.Contains(t, res.Body.String(), expected)
}