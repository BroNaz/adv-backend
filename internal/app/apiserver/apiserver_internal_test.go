package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_apiHello(t *testing.T) {
	s := New(NewConfig())
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api", nil)
	s.apiHello().ServeHTTP(response, request)
	assert.Equal(t, response.Body.String(), "HELLO API")
}
