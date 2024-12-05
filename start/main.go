package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bottino/aoc2024/utils"
)

type DayDir struct {
	Name   string
	Number int
}

func main() {
	now := time.Now()
	var dayNumber int
	flag.IntVar(&dayNumber, "d", now.Day(), "Advent of code day")

	dayName := fmt.Sprintf("day%02d", dayNumber)

	err := os.Mkdir(dayName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory for day %d: %v", dayNumber, err)
	}

	filePath := dayName + "/" + dayName + ".go"
	tmplPath := "templates/solution.tmpl"
	tmplData := DayDir{Name: dayName, Number: dayNumber}
	utils.CreateFromTemplate(filePath, tmplPath, tmplData)
}
