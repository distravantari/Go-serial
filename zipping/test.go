package zipping

import (
	"bytes"
	"io"
	"log"

	"github.com/alexmullins/zip"
)

func TTest() {
	contents := []byte("Hello World")

	// write a password zip
	raw := new(bytes.Buffer)
	zipw := zip.NewWriter(raw)
	w, err := zipw.Encrypt("max.zip", "golang")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Close()
}
