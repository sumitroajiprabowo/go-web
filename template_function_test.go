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

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + " My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("function").Parse(`{{ .SayHello "John" }}`))
	t.ExecuteTemplate(w, "function", MyPage{Name: "Budi"})
}

func TestTemplateFunction(t *testing.T) {
	// handler
	handler := http.HandlerFunc(TemplateFunction)

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
	expected := "Hello John My Name is Budi"

	assert.Equal(t, expected, string(body))
	// if string(body) != expected {
	// 	t.Errorf("Expected %s, got %s", expected, string(body))
	// }

	fmt.Println(string(body))

	fmt.Println("TestTemplateFunction OK")
}
