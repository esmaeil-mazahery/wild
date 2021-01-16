package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

type UploadFileResponse struct {
	Url string
}

var filesDirectory string = changePathSeperator("/wild/files/")

func changePathSeperator(src string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(src, "/", "\\", -1)
	}
	return src
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	extension := filepath.Ext(handler.Filename)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(filesDirectory, "f-*"+extension)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	body := UploadFileResponse{
		Url: strings.TrimPrefix(tempFile.Name(), filesDirectory),
	}

	jsonBody, err := json.Marshal(body)

	// return that we have successfully uploaded our file!
	fmt.Fprint(w, string(jsonBody))
}
