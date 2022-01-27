package helper

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateTmpFile(content string) *os.File {
	data := []byte(content)
	tmpfile, err := ioutil.TempFile("/tmp", "tdd.sh-")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func RemoveTmpFile(tmpfile *os.File) {
	os.Remove(tmpfile.Name())
}
