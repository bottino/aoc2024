package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bottino/aoc2024/utils"
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

	utils.CreateFromTemplate("days.go", "templates/days.tmpl", dayDirs)
}
