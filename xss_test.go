package golangweb

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Demonstrates how to use the template.HTML type to escape HTML.
Solution for security XSS Scripting
For example, the following code:
*/
func TemplateAutoEscaper(w http.ResponseWriter, r *http.Request) {

	//create a new template
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Auto Escaper",
		"Body":  "<p>This is a post</p>",
	})

}

func TestTemplateAutoEscaper(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateAutoEscaper)

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

	// // expected
	// expected := `<html><body><h1>Auto Escaper</h1><p>This is a post</p></body></html>`

	// // assert
	// assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

}

func TestTemplateAutoEscapeServer(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscaper),
	}

	err := server.ListenAndServe()

	if err != nil {
		t.Error("TestTemplateAutoEscapeServer: ", err)
	}

}

/*
Demonstrates how to use the template.HTML type to disable escape HTML.
If you want to disable escape HTML, you can use the template.HTML type.
or you can use the template.HTMLEscaper type.

If you want to disable escape CSS, you can use the template.CSS type.
or you can use the template.HTMLEscaper type.

If you want to disable escape JS, you can use the template.JS type.
or you can use the template.HTMLEscaper type.

Reference: https://golang.org/pkg/html/template/#HTMLEscaper
For example, the following code:
*/
func TemplateAutoEscaperDisable(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Auto Escaper",
		"Body":  template.HTML("<p>This is a post</p>"),
	})

}

func TestTemplateAutoEscaperDisable(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateAutoEscaperDisable)

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

	// // expected
	// expected := `<html><body><h1>Auto Escaper</h1><p>This is a post</p></body></html>`

	// // assert
	// assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

}

/*
Demonstrates XSS Scripting
in this example, we use the template.HTML to disable escape HTML.
and i use url query to pass the parameter.

Example call URL:
http://localhost:8080/?body=<p>Hahaha</p>

For example, the following code:
*/
func TemplateExampleXSS(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Auto Escaper",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})

}

func TestTemplateExampleXSS(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateExampleXSS)

	//request
	request := httptest.NewRequest("GET", "http://localhost:8080?body=<p>('Hahaha')</p>", nil)

	// recorder
	recorder := httptest.NewRecorder()

	// serve
	handler.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// check body
	body, _ := ioutil.ReadAll(response.Body)

	// // expected
	// expected := `<html><body><h1>Auto Escaper</h1><p>This is a post</p></body></html>`

	// // assert
	// assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

}

func TestTemplateExampleXSSServer(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateExampleXSS),
	}

	err := server.ListenAndServe()

	if err != nil {
		t.Error("TestTemplateExampleXSSServer: ", err)
	}
}
