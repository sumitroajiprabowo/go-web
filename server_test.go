package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the home page!"))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I'm an image"))
	})
	mux.HandleFunc("/img/lg/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I'm an image large"))
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}

func TestRequest(t *testing.T) {
	// r, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	// if err != nil {
	// 	t.Errorf("Error creating request: %s", err)
	// }
	// if r.Method != "GET" {
	// 	t.Errorf("Expected GET but received %s", r.Method)
	// }
	// if r.URL.Path != "/hello" {
	// 	t.Errorf("Expected /hello but received %s", r.URL.Path)
	// }
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// // request method
		// if r.Method != "GET" {
		// 	t.Errorf("Expected GET but received %s", r.Method)
		// }
		// // request path
		// if r.URL.Path != "/hello" {
		// 	t.Errorf("Expected /hello but received %s", r.URL.Path)
		// }
		// // request headers
		// if r.Header.Get("User-Agent") != "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36" {
		// 	t.Errorf("Expected User-Agent Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36 but received %s", r.Header.Get("User-Agent"))
		// }
		// // request body
		// if r.Body == nil {
		// 	t.Errorf("Expected body but received nil")
		// }
		// // response headers
		// w.Header().Set("Content-Type", "text/plain")
		// // response body
		// w.Write([]byte("Hello World"))

		fmt.Fprint(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)

	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}
}
