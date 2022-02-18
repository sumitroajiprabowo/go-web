package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create a new handler function for the server
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func TestHttp(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil) // create a request
	recorder := httptest.NewRecorder()                                        // create a response recorder

	HelloHandler(recorder, request) // call the handler
	response := recorder.Result()   // get the result

	result, err := ioutil.ReadAll(response.Body) // read the body
	if err != nil {
		t.Errorf("Error reading response body: %s", err) // handle error
	}

	expect := "Hello World" // expected result

	assert.Equal(t, expect, string(result)) // compare

}
