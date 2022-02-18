package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {

	// Create a form file
	err := myTemplates.ExecuteTemplate(w, "upload_form.gohtml", nil)

	if err != nil {
		panic(err)
	}

}

func Upload(w http.ResponseWriter, r *http.Request) {

	/*
	 Get the file from the request
	 Default file size limit is 32MB
	 If you want to change the limit, you can use the following code:

	 r.ParseMultipartForm(32 << 20)

	 Refer to the following link for more information:
	 https://golang.org/pkg/net/http/#Request.ParseMultipartForm

	*/
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	// // Close the file when the function returns
	defer file.Close()

	// Create destination target file upload path and file name
	fileDestination, err := os.Create("./tmp/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}

	// // Close the file when the function returns
	defer fileDestination.Close()

	// Copy the file to the destination
	_, err = io.Copy(fileDestination, file)

	if err != nil {
		panic(err)
	}

	//Create name variable to store the file name
	name := r.PostFormValue("name")

	//Create a map to store the data and pass it to the template
	myTemplates.ExecuteTemplate(w, "upload_success.gohtml",
		map[string]interface{}{
			"Name": name,
			"File": "/static/" + fileHeader.Filename,
		})

}

func TestUploadForm(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("tmp/"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		t.Error(err)
	}

}

//go:embed tmp/test.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {

	// Create a new multipart reader based on the request body
	body := new(bytes.Buffer)

	// Create a multipart writer and add the file to the body
	writer := multipart.NewWriter(body)

	// Create a new file to upload and write the contents to the file
	err := writer.WriteField("name", "Golang Upload File Test")

	if err != nil {
		panic(err)
	}

	// Create a new file to upload and write the contents to the file
	file, err := writer.CreateFormFile("file", "example.png")

	if err != nil {
		panic(err)
	}

	// Write the file contents to the file
	file.Write(uploadFileTest)

	// Close the writer
	writer.Close()

	// Create a new request with the content type set to the one we created
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)

	// Set the content type to multipart/form-data
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new recorder
	recorder := httptest.NewRecorder()

	// Serve the request
	Upload(recorder, request)

	// Check the response status code
	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	// Check the response body
	fmt.Println(string(bodyResponse))
}
