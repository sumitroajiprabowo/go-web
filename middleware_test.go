package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging request")
	l.Handler.ServeHTTP(w, r)
	fmt.Println("Logging response")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (e *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()
	e.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Middleware"))
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Test"))
	})

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("Error")
	})

	LogMiddleware := &LogMiddleware{mux}
	ErrorHandler := &ErrorHandler{LogMiddleware}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
