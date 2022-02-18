package golangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := new(http.Cookie)
	cookie.Name = "Golang-Cookie-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	cookie.HttpOnly = true

	// Create a cookie
	// cookie := &http.Cookie{
	// 	Name:     "Golang-Cookie-Name",
	// 	Value:    r.URL.Query().Get("name"),
	// 	HttpOnly: true,
	// 	Path:     "/",
	// }

	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Successfully set cookie "+cookie.Value)

}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Golang-Cookie-Name")
	if err != nil {
		fmt.Fprintln(w, "Cookie not set")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/set", SetCookie)
	mux.HandleFunc("/get", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}

	// request := httptest.NewRequest("GET", "http://localhost:8080/set", nil)
	// response := httptest.NewRecorder()
	// mux.ServeHTTP(response, request)
	// fmt.Println(response.Body.String())

}

func TestSetCookie(t *testing.T) {

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/set?name=Danu", nil)

	// recorder
	recorder := httptest.NewRecorder()

	//set cookie
	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s \n", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {

	//request
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// add cookie
	request.AddCookie(&http.Cookie{Name: "Golang-Cookie-Name", Value: "Budi"})

	// process add cookie
	// cookie := new(http.Cookie)
	// cookie.Name = "Golang-Cookie-Name"
	// cookie.Value = "Budi"
	// cookie.Path = "/"

	//recorder
	recorder := httptest.NewRecorder()

	//get cookie
	GetCookie(recorder, request)

	//read response
	result, err := ioutil.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
	}

	expect := "Hello Budi"

	assert.Equal(t, expect, string(result))

	fmt.Println(string(result))

	fmt.Println("TestGetCookie passed")

}
