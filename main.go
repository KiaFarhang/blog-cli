package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	const errorFileRequired = "You must pass a Markdown file name to this command; e.g. `main test.md`"

	programArgs := os.Args
	if len(programArgs) == 1 {
		log.Fatal(errorFileRequired)
	}

	blogFile := os.Args[1]

	if isBlankString(blogFile) || !strings.HasSuffix(blogFile, ".md") {
		log.Fatal(errorFileRequired)
	}

	postContent, err := ioutil.ReadFile(blogFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content: %s\n", postContent)

	fmt.Println("Welcome to the blog CLI!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter article slug (e.g. my-slug-here): ")
	slug, _ := reader.ReadString('\n')
	fmt.Printf("Slug: %v", slug)
}

func isBlankString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
