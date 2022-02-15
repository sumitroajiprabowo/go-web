package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	content := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "Content-Type: %s", content)
}

// create test server request header
func TestRequestHeader(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(RequestHeader)

	// Create a new request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)

	// Add a header to the request
	request.Header.Add("Content-Type", "application/json")

	// Create a new response recorder
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(recorder, request)

	// Get the response
	response := recorder.Result()

	// Read the body
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	// expected result
	expect := "Content-Type: application/json"

	// compare
	assert.Equal(t, expect, string(result))

	// print test result
	fmt.Println("TestRequestHeader passed")

}

// create new handler function with response header
func ResponseHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("X-Custom-Header", "belajar golang web")
	fmt.Fprintf(w, "Hello World")
}

// create test server response header
func TestResponseHeader(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(ResponseHeader)

	// Create a new request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)

	// Create a new response recorder
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(recorder, request)

	// Get the response get from key X-Custom-Header
	response := recorder.Header().Get("X-Custom-Header")

	// expected result
	expect := "belajar golang web"

	// print response
	fmt.Println(response)

	// compare
	assert.Equal(t, expect, response)

	// print test result
	fmt.Println("TestResponseHeader passed")

}
