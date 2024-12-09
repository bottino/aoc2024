package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/bottino/aoc2024/utils"
)

type DayDir struct {
	Name   string
	Number int
}

func main() {
	envDay, _ := strconv.Atoi(os.Getenv("AOC_DAY"))
	var dayNumber int
	flag.IntVar(&dayNumber, "d", envDay, "Advent of code day")
	flag.Parse()

	dayName := fmt.Sprintf("day%02d", dayNumber)

	// Exist if dir already exists
	if stat, err := os.Stat(dayName); err == nil && stat.IsDir() {
		fmt.Printf("Directory for day %d already created.\n", dayNumber)
		return
	}

	err := os.Mkdir(dayName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory for day %d: %v", dayNumber, err)
		return
	}

	filePath := dayName + "/" + dayName + ".go"
	tmplPath := "templates/solution.tmpl"
	tmplData := DayDir{Name: dayName, Number: dayNumber}
	utils.CreateFromTemplate(filePath, tmplPath, tmplData)
}
