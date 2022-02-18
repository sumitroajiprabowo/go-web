package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TemplateActionIfWithMap(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/example_if.gohtml"))

	t.ExecuteTemplate(w, "example_if.gohtml", map[string]interface{}{
		"Name": "John",
	})

}

func TestTemplateActionIfWithMap(t *testing.T) {

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

	// person := map[string]interface{}{
	// 	"Name": "John",
	// }

	// if person["Name"] == "" {
	// 	//expected
	// 	expected := `<html><body><h1>Hello</h1></body></html>`
	// 	assert.Equal(t, expected, string(body))
	// } else {
	// 	//expected
	// 	expected := `<html><body><h1>My Name is John</h1></body></html>`
	// 	assert.Equal(t, expected, string(body))
	// }

	//expected
	expected := `<html><body><h1>My Name is John</h1></body></html>`

	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateIfWithStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/example_if.gohtml"))

	t.ExecuteTemplate(w, "example_if.gohtml", Person{
		Name: "John",
	})
}

func TestTemplateIfWithStruct(t *testing.T) {
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

	// //Parse struct person
	// person := Person{
	// 	Name: "John",
	// }

	// if person.Name == "John" {
	// 	//expected
	// 	expected := `<html><body><h1>My Name is John</h1></body></html>`
	// 	assert.Equal(t, expected, string(body))
	// } else {

	// 	//expected
	// 	expected := `<html><body><h1>Hello</h1></body></html>`
	// 	assert.Equal(t, expected, string(body))
	// }

	//expected
	expected := `<html><body><h1>My Name is John</h1></body></html>`
	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionOperator)

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

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Hoobies": []string{"Skiing", "Snowboarding", "Hiking"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionRange)

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

	// 	//expected
	// 	expected := `<html>
	//     <body>

	//         <p>0 Skiing</p>

	//         <p>1 Snowboarding</p>

	//         <p>2 Hiking</p>

	//     </body>
	// </html>`
	// 	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/with.gohtml"))

	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Name": "John",
		"Age":  30,
		"Address": map[string]interface{}{
			"Street": "123 Main St",
			"City":   "Anytown",
			"State":  "CA",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateActionWith)

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

	// 	//expected
	// 	expected := `<html>
	//     <body>

	//     Name: John<br>
	//     Age: 30<br>

	//     Street: 123 Main St<br>
	//     City: Anytown<br>
	//     State: CA<br>

	//     </body>
	// </html>`
	// 	assert.Equal(t, expected, string(body))

	fmt.Println(string(body))

	fmt.Println("Success")
}
