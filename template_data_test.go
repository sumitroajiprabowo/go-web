package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {

	// Create a new template with some data
	t := template.Must(template.ParseFiles("templates/name_data.gohtml")) // ParseFiles returns a template, parses the named files and returns a *Template. If an error occurs, parsing stops and the returned *Template is nil.

	// Create a map to pass to our template.
	t.ExecuteTemplate(w, "name_data.gohtml", map[string]interface{}{ // ExecuteTemplate applies a parsed template to the specified data object. If an error occurs while executing the template, the method returns an error.
		"Name": "John", // Name is a string.
		"Age":  30,     // Age is an int.
		// "Address": map[string]interface{}{
		// 	"Street":  "123 Main St",
		// 	"City":    "Anytown",
		// 	"State":   "CA",
		// 	"Zip":     "90210",
		// 	"Country": "USA",
		// },
	})
}

func TestTemplateDataMap(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateDataMap)

	//request
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// recorder
	recorder := httptest.NewRecorder()

	// serve
	handler.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// check body
	body, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := `<html><body><h1>My Name is John and i have 30 years old.</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

type Address struct {
	Street  string
	City    string
	State   string
	Zip     string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {

	// Create a new template with some data
	t := template.Must(template.ParseFiles("templates/example_data.gohtml")) // ParseFiles returns a template, parses the named files and returns a *Template. If an error occurs, parsing stops and the returned *Template is nil.

	// Create a map to pass to our template.
	t.ExecuteTemplate(w, "example_data.gohtml", Person{ // ExecuteTemplate applies a parsed template to the specified data object. If an error occurs while executing the template, the method returns an error.
		Name: "John", // Name is a string.
		Age:  30,     // Age is an int.
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Zip:     "90210",
			Country: "USA",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateDataStruct)

	//request
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// recorder
	recorder := httptest.NewRecorder()

	// serve
	handler.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// check body
	body, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := `<html><body><h1>My Name is John from Anytown and i have 30 years old.</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}
