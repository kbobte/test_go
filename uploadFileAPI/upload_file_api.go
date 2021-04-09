package uploadFileAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// UploadFile godoc
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limit 10MB
	r.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	fmt.Println("File Info:")
	fmt.Println("File Name:", handler.Filename)
	fmt.Println("File Size:", handler.Size)
	fmt.Println("File Type:", handler.Header.Get("Content-Type"))

	tempFile, err2 := ioutil.TempFile("uploads", "upload-*.jpg")
	if err2 != nil {
		fmt.Println("err2: ", err2)
	}
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println("err3: ", err3)
	}

	tempFile.Write(fileBytes)
	fmt.Println("Done")
}
