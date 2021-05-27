// Walk - finding hash256 of files in given dir
// without goroutines
package main

import (
	//"archive/zip"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath" // выбирает / или \ в завис-ти от системы
	//"github.com/mholt/archiver"// - for archives
	//"github.com/iafan/cwalk" // - for goroutines
)

/*func isArchive(path string) bool{
	return false
}*/

func myWalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
		/*} else if isArchive(path) {
		err = z.Walk("/Users/matt/Desktop/test.zip", func(f archiver.File) error {
			if zfh, ok := f.Header.(zip.FileHeader); ok{
				fmt.Print(info.Size(), " ") // file size
				hash25(path)
			}
			return nil
		})*/
	} else {
		fmt.Print(info.Size(), " ") // file size
		hash25(path)
		return nil
	}
}

func hash25(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x ", h.Sum(nil))
	fmt.Println(path)
}

func main() {
	const root = "I:\\Download\\Apk Huawei" // dir root
	if err := filepath.Walk(root, myWalkFunc); err != nil {
		//if err := cwalk.Walk(root, myWalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
