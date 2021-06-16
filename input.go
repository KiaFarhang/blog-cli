package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type InputReader interface {
	GetInputFile() (string, error)
}

type CLIInputReader struct {
}

const errorFileRequired = "You must pass a Markdown file name to this command; e.g. `main test.md`"

func (c CLIInputReader) GetInputFile() (string, error) {

	programArgs := os.Args
	if len(programArgs) == 1 {
		log.Fatal(errorFileRequired)
	}

	blogFile := os.Args[1]

	postContent, err := ioutil.ReadFile(blogFile)

	return string(postContent), err
}

func isBlankString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
