package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create a new handler function for the server
func SayHello(w http.ResponseWriter, r *http.Request) {

	// Get the name from the query string
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Hello World")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}

}

func TestQueryParam(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(SayHello)

	// Create a new request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=John", nil)

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

	expect := "Hello, John!"

	assert.Equal(t, expect, string(result))

	fmt.Println("TestQueryParam passed")

}

// Create new handler function for the server with multiple parameter
func MultipleParameter(w http.ResponseWriter, r *http.Request) {

	// Get the name from the query string
	name := r.URL.Query().Get("name")

	// Get the age from the query string
	age := r.URL.Query().Get("age")

	// Get the
	fmt.Fprintf(w, "Hello, %s! You are %s years old.", name, age)
}

func TestMultipleParameter(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(MultipleParameter)

	// Create a new request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=John&age=20", nil)

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

	expect := "Hello, John! You are 20 years old."

	assert.Equal(t, expect, string(result))

	fmt.Println("TestMultipleParameter passed")

}

func MultipleValueParameter(w http.ResponseWriter, r *http.Request) {

	// get the query string
	query := r.URL.Query()

	// get the name from the query string
	names := query["name"]

	// get the age from the query string
	age := r.URL.Query().Get("age")

	// Loop through the names
	fmt.Fprint(w, strings.Join(names, " ")+" are "+age+" years old.") // print the result
}

//create test multiple value query parameter
func TestMultipleValueParameter(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(MultipleValueParameter)

	// Create a new request
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=John&name=Jane&name=Mary&age=20", nil)

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

	expect := "John Jane Mary are 20 years old."

	assert.Equal(t, expect, string(result))

	fmt.Println("TestMultipleValueParameter passed")

}
