package main

import (
	"brian-big/gophercise-quiz-game/files"
	"flag"
)

func main() {

	fileName := flag.String("name", "example.txt", "name of the file you want to create")
	content := flag.String("content", "", "text content you want to write in the file")

	flag.Parse()

	lines := []string{
		"Hello, let's write line by line",
		"Have you ever had a story you wanted so bad to tell?",
		"Well, I will help you tell the world your story",
	}
	files.CreateFile(*fileName, []byte(*content))

	files.WriteByLines(*fileName, lines)

	//test append Function
	files.AppendFile(*fileName, *content)
}
