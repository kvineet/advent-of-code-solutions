package main

import (
	"fmt"
	"os"

	"github.com/niklasfasching/go-org/org"
)

func main() {
	filePath := "test.org"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening org file:", err)
		return
	}
	// Parse content as org file
	doc := org.New().Parse(file, file.Name())
	if doc.Error != nil {
		fmt.Println("Error parsing org file:", err)
		return
	}

	// Access parsed org elements
	for _, element := range doc.Outline.Headline.Status {
		// Process each org element
		fmt.Println(element)
	}
}
