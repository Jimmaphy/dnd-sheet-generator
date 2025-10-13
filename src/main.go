package main

import (
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/services"
)

func main() {
	templateService, err := services.NewTemplateService("usage.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err := templateService.GetTemplateContent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(content)
}
