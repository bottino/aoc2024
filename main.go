package main

// Generate days.go, creating the list of days
//go:generate go run ./gen

import (
	_ "embed"
	"flag"
	"fmt"
	"maps"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"
)

type aocDay struct {
	Number int
	Part1  func(string) any
	Part2  func(string) any
}

// Only run from the root of the repo
func (day aocDay) GetInputPath(useExample bool) string {
	dayName := fmt.Sprintf("day%02d", day.Number)
	if useExample {
		return "inputs/" + dayName + "/example.txt"
	} else {
		return "inputs/" + dayName + "/input.txt"
	}
}

func getEnvWithDefault(varName string, defValue int) int {
	env := os.Getenv(varName)
	if env == "" {
		return defValue
	}

	value, err := strconv.Atoi(env)
	if err != nil {
		fmt.Errorf("Couldn't convert %s variable value %s to string. Using default %d", varName, env, defValue)
		return defValue
	}

	return value
}

func main() {
	// Flags
	var (
		dayNumber  int
		part       int
		useExample bool
		verbose    bool
		runAllDays bool
	)
	defaultDay := getEnvWithDefault("AOC_DAY", time.Now().Day())
	defaultPart := getEnvWithDefault("AOC_PART", 0)
	flag.IntVar(&dayNumber, "d", defaultDay, "Advent of code day")
	flag.IntVar(&part, "p", defaultPart, "The part of the puzzle")
	flag.BoolVar(&useExample, "e", false, "Use the example as input")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.BoolVar(&runAllDays, "all", false, "Run all days")
	flag.Parse()

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	var aocSolutions []AocSolution
	if !runAllDays {
		aocSolutions = []AocSolution{{Day: dayNumber, Part: part, UseExample: useExample}}
	} else {
		for _, d := range slices.Sorted(maps.Keys(days)) {
			for _, part := range []int{1, 2} {
				aocSolutions = append(aocSolutions, AocSolution{Day: d, Part: part, UseExample: false})
			}
		}
	}

	for _, sol := range aocSolutions {
		sol.Solve()

		if verbose {
			fmt.Println(&sol)
		} else {
			fmt.Println(sol.Solution)
		}
	}
}

type AocSolution struct {
	Day        int
	Part       int
	Time       float64
	Solution   any
	UseExample bool
}

func (s *AocSolution) String() string {
	var exampleStr string
	if s.UseExample {
		exampleStr = " (example)"
	}
	return fmt.Sprintf(
		"Day %02d, part %d%s. Solution: %v, Time elapsed %.2f ms",
		s.Day, s.Part, exampleStr, s.Solution, s.Time,
	)
}

func (s *AocSolution) Solve() {
	day, ok := days[s.Day]
	if ok == false {
		fmt.Printf("Day %d does not exist", s.Day)
		return
	}

	path := day.GetInputPath(s.UseExample)
	inputBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}
	input := strings.TrimRight(string(inputBytes), "\n")

	start := time.Now()
	if s.Part == 1 {
		s.Solution = day.Part1(input)
	} else {
		s.Solution = day.Part2(input)
	}

	end := time.Now()
	s.Time = float64(end.Sub(start).Microseconds()) / 1000
}
