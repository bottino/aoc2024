package day17

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Part1(input string) any {
	comp := readInput(input)
	return comp.runProgram()
}

func Part2(input string) any {
	comp := readInput(input)

	n := len(comp.program)

	var currRegA int
	maxTries := 8
	for j := 0; j < n; j++ {
		for i := 0; i < maxTries; i++ { // Bootstrap the first iteration
			tmpRegA := 8*currRegA + i
			cmp := Computer{
				regA:    tmpRegA,
				program: comp.program,
				output:  []int{},
			}

			expOut := comp.program[n-j-1 : n]
			cmp.runProgram()
			fmt.Println(expOut, cmp.output)

			if outputEqual(expOut, cmp.output) {
				currRegA = tmpRegA
				break
			}

			if i == maxTries-1 {
				return fmt.Errorf("Couldn't build correct result")
			}
		}
	}

	sol := currRegA
	fmt.Println(sol)

	comp.regA = sol
	return comp.runProgram()
}

func outputEqual(exp []int, actual []int) bool {
	if len(exp) != len(actual) {
		return false
	}

	for i := range len(exp) {
		if exp[i] != actual[i] {
			return false
		}
	}

	return true
}

// A structure that holds computer operations
type Computer struct {
	regA, regB, regC int
	program          []int
	output           []int
	pInstr           *int
}

func (c *Computer) lastOutput() int {
	if len(c.output) <= 0 {
		return -1
	}

	return c.output[len(c.output)-1]
}

func (c *Computer) Clone() Computer {
	return Computer{
		regA: c.regA, regB: c.regB, regC: c.regC,
		program: c.program,
		output:  []int{},
	}
}

func (c *Computer) combo(opcode int) int {
	switch opcode {
	case 0, 1, 2, 3:
		return opcode
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	case 7:
		panic("Reserved number 7")
	default:
		panic("Not an 3bit number")
	}
}

func (c *Computer) runProgram() string {
	c.pInstr = new(int)
	instructions := []func(int){c.adv, c.bxl, c.bst, c.jnz, c.bxc, c.out, c.bdv, c.cdv}
	for *c.pInstr < len(c.program) {
		p := *c.pInstr
		instructions[c.program[p]](c.program[p+1])
	}

	out := ""
	for i, o := range c.output {
		out += fmt.Sprintf("%d", o)
		if i < len(c.output)-1 {
			out += ","
		}
	}

	return out
}

// 0
func (c *Computer) adv(op int) {
	den := int(math.Pow(2, float64(c.combo(op))))
	c.regA = c.regA / den
	*c.pInstr += 2
}

// 1
func (c *Computer) bxl(op int) {
	c.regB = c.regB ^ op
	*c.pInstr += 2
}

// 2
func (c *Computer) bst(op int) {
	c.regB = c.combo(op) % 8
	*c.pInstr += 2
}

// 3
func (c *Computer) jnz(op int) {
	if c.regA == 0 {
		*c.pInstr += 2
		return
	}

	*c.pInstr = op
}

// 4
func (c *Computer) bxc(_ int) {
	c.regB = c.regB ^ c.regC
	*c.pInstr += 2
}

// 5
func (c *Computer) out(op int) {
	c.output = append(c.output, c.combo(op)%8)
	*c.pInstr += 2
}

// 6
func (c *Computer) bdv(op int) {
	den := int(math.Pow(2, float64(c.combo(op))))
	c.regB = c.regA / den
	*c.pInstr += 2
}

// 7
func (c *Computer) cdv(op int) {
	den := int(math.Pow(2, float64(c.combo(op))))
	c.regC = c.regA / den
	*c.pInstr += 2
}

func readInput(input string) Computer {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)
	nums := []int{}
	for _, m := range matches {
		n, _ := strconv.Atoi(m)
		nums = append(nums, n)
	}

	return Computer{
		regA:    nums[0],
		regB:    nums[1],
		regC:    nums[2],
		program: nums[3:],
		output:  []int{},
	}
}
