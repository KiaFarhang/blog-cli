package main

import (
	"fmt"
	"log"
)

func main() {

	reader := CLIInputReader{}

	postContent, err := reader.GetInputFile()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content: %s\n", postContent)

	// fmt.Println("Welcome to the blog CLI!")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter article slug (e.g. my-slug-here): ")
	// slug, _ := reader.ReadString('\n')
	// fmt.Printf("Slug: %v", slug)
}
