package zipping

// func TTest() {
// 	contents := []byte("Hello World")
// 	filename:= "max.zip"
// 	// write a password zip
// 	newfile, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer newfile.Close()

// 	zipw := zip.NewWriter(newfile)
// 	zipw := zip.NewWriter(raw)
// 	w, err := zipw.Encrypt(, "golang")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = io.Copy(w, bytes.NewReader(contents))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	zipw.Close()
// }
