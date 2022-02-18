package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Example of redirect")

}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)

}

func RedirectOut(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://google.com", http.StatusTemporaryRedirect)

}

func TestRedirect(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		t.Error(err)
	}

}
