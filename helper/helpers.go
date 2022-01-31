package helper

import (
	uuid "github.com/nu7hatch/gouuid"
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

func RemoveTmpFile(file *os.File) {
	os.Remove(file.Name())
}

func RemoveTmpFileByName(file string) {
	os.Remove(file)
}

func TmpFileName() string {
	id, _ := uuid.NewV4()
	return "/tmp/tdd.sh-" + id.String()
}
