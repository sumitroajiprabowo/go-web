package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//create variabel with global template
//go:embed templates/*.gohtml
var mytemplates embed.FS

//create variable with template
var myTemplates = template.Must(template.ParseFS(mytemplates, "templates/*.gohtml"))

//create function with template
func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "example.gohtml", "Hello From Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateCaching)

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
	expected := `<html><body><h1>Hello From Template Caching</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("TestTemplateCaching OK")
}

func TemplateActionIfWithMapCaching(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "example_if.gohtml", map[string]interface{}{
		"Name": "John",
	})

}

func TestTemplateActionIfWithMapCaching(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionIfWithMapCaching)

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
	expected := `<html><body><h1>My Name is John</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("TestTemplateActionIfWithMapCaching OK")
}

func TemplateIfWithStructCaching(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "example_if.gohtml", Person{
		Name: "John",
	})
}

func TestTemplateIfWithStructCaching(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateIfWithStruct)

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
	expected := `<html><body><h1>My Name is John</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateActionOperatorCaching(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 90,
	})
}

func TestTemplateActionOperatorCaching(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionOperatorCaching)

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
	expected := `<html><body><h1>Good</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateActionRangeCaching(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Hoobies": []string{"Skiing", "Snowboarding", "Hiking"},
	})
}

func TestTemplateActionRangeCaching(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionRangeCaching)

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

	fmt.Println(string(body))

	fmt.Println("Success")
}
