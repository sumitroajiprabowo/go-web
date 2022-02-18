package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

/*
Example of template global function
Reference : https://golang.org/pkg/text/template/#FuncMap
https://github.com/golang/go/blob/master/src/text/template/funcs.go
*/

func TemplateGlobalFunction(w http.ResponseWriter, r *http.Request) {

	// create template with global function
	t := template.Must(template.New("function").Parse(`{{ len .Name }}`))

	// execute template and write to response
	t.ExecuteTemplate(w, "function", MyPage{Name: "Budi"})
}

func TestTemplateGlobalFunction(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateGlobalFunction)

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

	// expected
	expected := "4"

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("TestTemplateFunction OK")
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {

	// create template with global function
	t := template.New("function").Funcs(template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
		"title": func(s string) string {
			return strings.Title(s)
		},
		"trim": func(s string) string {
			return strings.Trim(s, " ")
		},
		"trimLeft": func(s string) string {
			return strings.TrimLeft(s, " ")
		},
		"trimRight": func(s string) string {
			return strings.TrimRight(s, " ")
		},
		"trimSpace": func(s string) string {
			return strings.TrimSpace(s)
		},
		"trimPrefix": func(s, prefix string) string {
			return strings.TrimPrefix(s, prefix)
		},
		"trimSuffix": func(s, suffix string) string {
			return strings.TrimSuffix(s, suffix)
		},
		"repeat": func(s string, count int) string {
			return strings.Repeat(s, count)
		},
		"hasPrefix": func(s, prefix string) bool {
			return strings.HasPrefix(s, prefix)
		},
		"hasSuffix": func(s, suffix string) bool {
			return strings.HasSuffix(s, suffix)
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(s, substr)
		},
	})

	// parse template using global function
	t = template.Must(t.Parse(`{{ upper .Name }} {{ lower .Name }} {{ title .Name }} {{ trim .Name }} {{ trimLeft .Name }} {{ trimRight .Name }} {{ trimSpace .Name }} {{ trimPrefix .Name "B" }} {{ trimSuffix .Name "i" }} {{ repeat .Name 3 }} {{ hasPrefix .Name "B" }} {{ hasSuffix .Name "i" }} {{ contains .Name "Budi" }}`))

	// execute template and write to response
	t.ExecuteTemplate(w, "function", MyPage{Name: "Budi"})
}

func TestTemplateFunctionMap(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateFunctionCreateGlobal)

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

	// expected
	expected := "BUDI budi Budi Budi Budi Budi Budi udi Bud BudiBudiBudi true true true"

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("TestTemplateFunctionMap OK")
}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {

	// create template
	t := template.New("function")

	// create function and add to template function map
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	/*
		Create a template with a function with a pipeline
		Reference : https://golang.org/pkg/text/template/#FuncMap
	*/
	// t := template.New("function").Funcs(map[string]interface{}{
	// 	"sayHello": func(name string) string {
	// 		return "Hello " + name
	// 	},
	// 	"upper": func(name string) string {
	// 		return strings.ToUpper(name)
	// 	},
	// })

	// parse template using template function map
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	// execute template and write to response
	t.ExecuteTemplate(w, "function", MyPage{Name: "Budi"})

}

func TestTemplateFunctionPipeline(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateFunctionPipeline)

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

	// expected
	expected := "HELLO BUDI"

	// assert result with expected
	assert.Equal(t, expected, string(body))

	// print result
	fmt.Println(string(body))

	fmt.Println("TestTemplateFunctionPipeline OK")
}
