package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"text/template"
)

type DayDir struct {
	Name   string
	Number int
}

func newDayDir(dirName string) DayDir {
	dayDir := DayDir{Name: dirName}
	fmt.Sscanf(dirName, "day%02d", &dayDir.Number)
	return dayDir
}

func sortByName(a DayDir, b DayDir) int {
	return strings.Compare(a.Name, b.Name)
}

func main() {
	files, _ := os.ReadDir(".")
	var dayDirs []DayDir
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "day") && file.IsDir() {
			dayDirs = append(dayDirs, newDayDir(file.Name()))
		}
	}

	slices.SortFunc(dayDirs, sortByName)
	var tmplFile = "templates/days.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		fmt.Printf("Error reading template: %v", err)
	}

	fout, err := os.Create("days.go")
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
	}
	defer fout.Close()

	err = tmpl.Execute(fout, dayDirs)
	if err != nil {
		fmt.Printf("Error creating file from template: %v", err)
	}
}

