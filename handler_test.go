package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func TestHandlerServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}
