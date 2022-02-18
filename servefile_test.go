package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/index.html")
	} else {
		http.ServeFile(w, r, "./resources/not_found.html")
	}
}

func TestServeFile(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}

//go:embed resources/index.html
var resourcesOK string

//go:embed resources/not_found.html
var resourcesNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourcesOK)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}
