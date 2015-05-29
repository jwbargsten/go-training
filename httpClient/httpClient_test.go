package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItSuccess(t *testing.T) {
	// Create a fake server and let it return a 200 OK with json response body
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `
{
   "args": {
      "arg1": "1",
      "arg2": "2"
   },
   "headers": {
      "Accept": "application/json",
      "Accept-Encoding": "gzip"
    },
    "origin": "145.53.70.8",
    "url": "http://httpbin.org/get?arg1=1&arg2=2"
}`)

	}))
	defer server.Close()

	// Perform the action against the fake server
	actualResponse, err := getIt(fmt.Sprintf("%s/get", server.URL))

	// Verify the response
	assert.NoError(t, err)
	assert.Equal(t, "1", actualResponse.Args.Arg1)
	assert.Equal(t, "2", actualResponse.Args.Arg2)
	assert.Equal(t, "application/json", actualResponse.Headers.Accept)
	assert.Equal(t, "gzip", actualResponse.Headers.AcceptEncoding)
	assert.Equal(t, "http://httpbin.org/get?arg1=1&arg2=2", actualResponse.Url)
}

func TestGetItFailure(t *testing.T) {
	// Create a fake server and let it return a 404 not-found error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// Perform the action
	actualResponse, err := getIt(fmt.Sprintf("%s/get", server.URL))

	// Verify the response
	assert.Error(t, err)
	assert.Nil(t, actualResponse)

}
