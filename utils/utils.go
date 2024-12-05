package utils

import (
	"fmt"
	"os"
	"text/template"
)

func CreateFromTemplate(filePath string, tmplPath string, tmplData any) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("Error reading template: %v", err)
		return
	}

	fout, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return
	}
	defer fout.Close()

	err = tmpl.Execute(fout, tmplData)
	if err != nil {
		fmt.Printf("Error creating file from template: %v", err)
		return
	}
}
