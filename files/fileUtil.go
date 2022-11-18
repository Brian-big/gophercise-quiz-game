package files

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(FileName string, content string) os.File {
	file, err := os.Create(FileName)

	if err != nil {
		log.Fatal("An error occured: ", err.Error())
	}

	defer file.Close()

	if len(content) == 0 {
		content = "Empty file. Ready to write"
	}

	n, err := file.WriteString(content)
	if err != nil {
		log.Fatal("Could not write content to file ", err.Error())
	}
	log.Output(n, fmt.Sprintf("Content of size %d wrote to file", n))

	return *file
}
