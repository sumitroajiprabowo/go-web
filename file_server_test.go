package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// Create Test Function for FileServer
func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	// fileServer := FileServer(directory)
	fileServer := http.FileServer(directory)

	// Create mux
	mux := http.NewServeMux()

	// localhost:8080/resources/static/filename
	// mux.Handle("/static/", fileServer)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Create server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Start server
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}

}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	// Create mux
	mux := http.NewServeMux()

	// localhost:8080/resources/static/filename
	// mux.Handle("/static/", fileServer)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Create server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Start server
	err := server.ListenAndServe()
	if err != nil {
		t.Errorf("Server failed to start: %s", err)
	}

}
