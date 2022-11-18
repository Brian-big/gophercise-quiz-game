package main

import (
	"brian-big/gophercise-quiz-game/files"
	"flag"
)

func main() {

	fileName := flag.String("name", "example.txt", "name of the file you want to create")
	content := flag.String("content", "", "text content you want to write in the file")

	flag.Parse()

	files.CreateFile(*fileName, []byte(*content))
}
