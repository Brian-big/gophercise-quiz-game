package main

import (
	"brian-big/gophercise-quiz-game/files"
	"flag"
)

func main() {

	fileName := flag.String("name", "example.txt", "name of the file you want to create")

	flag.Parse()

	files.CreateFile(*fileName)
}
