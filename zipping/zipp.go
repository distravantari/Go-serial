package zipping

import (
	"JPRO/modules"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/alexmullins/zip"
)

func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func Zipit() {
	files := []string{modules.FileLoc + "/Raw.txt", modules.FileLoc + "/Rep.log", modules.FileLoc + "/Max.xlsx"}
	tie := time.Now().Format("2006-01-02")
	hour := strconv.Itoa(time.Now().Hour())
	min := strconv.Itoa(time.Now().Minute())
	sec := strconv.Itoa(time.Now().Second())
	// output := "backup" + tie + "T" + hour + "." + min + "." + sec + ".zip"
	output := "backup" + tie + "T" + hour + "." + min + "." + sec

	err := ZipFiles(output, files)

	if err != nil {
		fmt.Println("wow ", err)
	}

	fmt.Println("Zipped File: " + output)
}

func ZipFiles(filename string, files []string) error {
	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)

	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}

	return nil
}
