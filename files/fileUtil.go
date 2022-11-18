package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
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

// function to append content to a given text
func AppendFile(fileName string, data string) os.File {
	file, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_WRONLY,
		fs.ModeAppend)
	if err != nil {
		log.Fatal("Could not open file:", err.Error())
	}
	defer file.Close()

	if len(data) == 0 {
		content := "Appended this line as default"
		data = content
	}

	_, writeErr := fmt.Fprintln(file, data)

	if writeErr != nil {
		log.Fatal(writeErr.Error())
	}

	return *file
}

func ReadFileInfo(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("=======>File info:")
	fmt.Printf("|=> file Name: %s\n|=> Modified: %s\n|=> Mode: %s\n|=> Size: %d bytes\n======================", fileInfo.Name(), fileInfo.ModTime().Format("15:04:03"), fileInfo.Mode(), fileInfo.Size())
}

func ReadFileContents(fileName string) {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(string(contents))
}

func main() {

	fileName := flag.String("name", "example.txt", "name of the file you want to create")
	content := flag.String("content", "", "text content you want to write in the file")

	flag.Parse()

	lines := []string{
		"Hello, let's write line by line",
		"Have you ever had a story you wanted so bad to tell?",
		"Well, I will help you tell the world your story",
	}
	CreateFile(*fileName, []byte(*content))

	WriteByLines(*fileName, lines)

	//test append Function
	AppendFile(*fileName, *content)

	// readFile
	ReadFileInfo(*fileName)
	fmt.Println("\n*************File Contents*****************")
	ReadFileContents(*fileName)
	fmt.Println("***********End File Contents***************")

}
