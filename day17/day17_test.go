package day17

import (
	_ "embed"
	"testing"
)

type TestData struct {
	comp      Computer
	expOutput string
	expReg    map[rune]int
}

func TestDay17(t *testing.T) {
	testData := []TestData{
		{
			comp:   Computer{regC: 9, program: []int{2, 6}},
			expReg: map[rune]int{'B': 1},
		},
		{
			comp:      Computer{regA: 10, program: []int{5, 0, 5, 1, 5, 4}},
			expOutput: "0,1,2",
		},
		{
			comp:      Computer{regA: 2024, program: []int{0, 1, 5, 4, 3, 0}},
			expOutput: "4,2,5,6,7,7,7,7,3,1,0",
			expReg:    map[rune]int{'A': 0},
		},
		{
			comp:   Computer{regB: 29, program: []int{1, 7}},
			expReg: map[rune]int{'B': 26},
		},
		{
			comp:   Computer{regB: 2024, regC: 43690, program: []int{4, 0}},
			expReg: map[rune]int{'B': 44354},
		},
		{
			comp:      Computer{regA: 117440, program: []int{0, 3, 5, 4, 3, 0}},
			expOutput: "0,3,5,4,3,0",
		},
	}

	for i, d := range testData {
		out := d.comp.runProgram()

		if d.expOutput != "" {
			if out != d.expOutput {
				t.Errorf("Test %d output mismatch. Expected: %s - Actual: %s", i, d.expOutput, out)
			}
		}

		if d.expReg != nil {
			var fail bool
			for k, v := range d.expReg {
				switch k {
				case 'A':
					fail = v != d.comp.regA
				case 'B':
					fail = v != d.comp.regB
				case 'C':
					fail = v != d.comp.regC
				}
			}

			if fail {
				t.Errorf(
					"Test %d failed: Expected: %v - Actual: {A:%d, B:%d, C:%d}",
					i, d.expReg, d.comp.regA, d.comp.regB, d.comp.regC,
				)
			}
		}
	}
}
