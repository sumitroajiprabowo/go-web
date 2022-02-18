package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	// Get the name from the query string
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(400) // Bad Request
		fmt.Fprintf(w, "Name is required")
	} else {
		w.WriteHeader(200) // OK
		fmt.Fprintf(w, "Hello, %s!", name)
	}

	// if r.Method != "GET" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	w.Write([]byte("Method not allowed"))
	// } else {
	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprintf(w, "Hello, %s!", name)
	// }

}

func TestResponseCode(t *testing.T) {

	// handler
	handler := http.HandlerFunc(ResponseCode)

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=John", nil)

	// response recorder
	recorder := httptest.NewRecorder()

	// call the handler
	handler.ServeHTTP(recorder, request)

	// get the response
	response := recorder.Result()

	// read the body
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	// expected result
	expect := "Hello, John!"

	// compare
	if response.StatusCode == 200 {
		assert.Equal(t, expect, string(result))
		fmt.Println("TestResponseCode passed")
	} else {
		fmt.Println("TestResponseCode failed")
	}
}
