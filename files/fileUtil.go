package files

import (
	"log"
	"os"
)

func CreateFile(FileName string) os.File {
	file, err := os.Create(FileName)

	if err != nil {
		log.Fatal("An error occured: ", err.Error())
	}
	return *file
}
