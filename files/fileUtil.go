package files

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(FileName string, data []byte) os.File {
	file, err := os.Create(FileName)

	if err != nil {
		log.Fatal("An error occured: ", err.Error())
	}

	defer file.Close()

	if len(data) == 0 {
		content := "Empty file. Ready to write"
		data = []byte(content)
	}

	n, err := file.Write(data)
	if err != nil {
		log.Fatal("Could not write content to file ", err.Error())
	}
	log.Output(n, fmt.Sprintf("Content of size %d wrote to file", n))

	return *file
}

func WriteByLines(fileName string, data []string) os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("An error occured", err.Error())
	}
	defer file.Close()

	for _, line := range data {

		_, err := fmt.Fprintln(file, line)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	return *file

}
