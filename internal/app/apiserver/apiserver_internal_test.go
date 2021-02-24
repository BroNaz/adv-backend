package apiserver

var (
	testAPIHelp = "501 - Not Implemented"
)

/*
func TestAPIServer_apiHello(t *testing.T) {
	s := New(NewConfig())
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api", nil)
	s.apiHello().ServeHTTP(response, request)
	assert.Equal(t, response.Body.String(), testAPIHelp)
}
*/
