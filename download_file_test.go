package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {

	// Get the file name from the request
	file := r.URL.Query().Get("download")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}

	// Open the file and download it
	w.Header().Add("Content-Disposition", "attachment; filename="+file)

	http.ServeFile(w, r, "./resources/"+file)

}

/*
Create test cases for download file using fuction DonwloadFile
you can see with the following in your browser
http://localhost:8080/download?download=example.png
*/
func TestDownloadFile(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
