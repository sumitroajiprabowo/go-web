package golangweb

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {

	templateText := `<html><body><h1>{{.}}</h1></body></html>`

	// t, err := template.New("simple").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("simple").Parse(templateText))

	err := t.ExecuteTemplate(w, "simple", "Hello World")
	if err != nil {
		panic(err)
	}

}

func TestSimpleHtml(t *testing.T) {

	// request httptest
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// recorder httptest
	recorder := httptest.NewRecorder()

	// handler
	handler := http.HandlerFunc(SimpleHTML)

	// serve
	handler.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// check body
	body, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := `<html><body><h1>Hello World</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func SimpleHTMLTemplate(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/example.gohtml"))

	err := t.ExecuteTemplate(w, "example.gohtml", "Hello From HTML Template")
	if err != nil {
		panic(err)
	}

}

func TestSimpleHTMLTemplate(t *testing.T) {

	// handler
	handler := http.HandlerFunc(SimpleHTMLTemplate)

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
	expected := `<html><body><h1>Hello From HTML Template</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := t.ExecuteTemplate(w, "example.gohtml", "Hello From Template Directory")
	if err != nil {
		panic(err)
	}

}

func TestTemplateDirectory(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateDirectory)

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
	expected := `<html><body><h1>Hello From Template Directory</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	err := t.ExecuteTemplate(w, "example.gohtml", "Hello From Template Embed")
	if err != nil {
		panic(err)
	}

}

func TestTemplateEmbed(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateEmbed)

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
	expected := `<html><body><h1>Hello From Template Embed</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}
