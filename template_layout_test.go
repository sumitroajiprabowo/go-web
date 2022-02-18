package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/header.gohtml", "templates/layout.gohtml", "templates/footer.gohtml"))

	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title": "Layout",
		"Name":  "John",
	})

}

func TestTemplateLayout(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateLayout)

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

	fmt.Println("TestTemplateLayout OK")

}

func TemplateLayoutDefine(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/header_define.gohtml", "templates/layout_define.gohtml", "templates/footer_define.gohtml"))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Layout Define",
		"Name":  "John",
	})

}

func TestTemplateLayoutDefine(t *testing.T) {

	// handler
	handler := http.HandlerFunc(TemplateLayoutDefine)

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

	fmt.Println("TestTemplateLayoutDefine OK")

}
