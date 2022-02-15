package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//create a new handler function for the server with form post
func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)

}

func TestFormPost(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(FormPost)

	// request body
	bodyRequest := strings.NewReader("firstName=John&lastName=Doe")

	// Create a new request
	request := httptest.NewRequest("POST", "http://localhost:8080/hello", bodyRequest)
	// request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/hello", bodyRequest)

	// create request header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	expect := "Hello, John Doe"

	assert.Equal(t, expect, string(result))

	fmt.Println("TestFormPost passed")

}

//example of http request with header
func PostForm(w http.ResponseWriter, r *http.Request) {

	firstName := r.PostFormValue("firstName")
	lastName := r.PostFormValue("lastName")

	fmt.Fprintf(w, "Welcome, %s %s", firstName, lastName)

}

//create test server Post Form
func TestPostForm(t *testing.T) {

	// Create a new handler function for the server
	handler := http.HandlerFunc(PostForm)

	// request body
	bodyRequest := strings.NewReader("firstName=John&lastName=Doe")

	// Create a new request
	request := httptest.NewRequest("POST", "http://localhost:8080/hello", bodyRequest)

	// create request header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	expect := "Welcome, John Doe"

	assert.Equal(t, expect, string(result))

	fmt.Println("TestPostForm passed")

}
