package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the blog CLI!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter article slug (e.g. my-slug-here): ")
	slug, _ := reader.ReadString('\n')
	fmt.Printf("Slug: %v", slug)
}

func isBlankString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
