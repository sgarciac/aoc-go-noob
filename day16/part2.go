package main

import (
	"regexp"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type inst struct {
	opcode int
	op1 int
	op2 int
	res int
}
type test struct {
	before []int
	inst inst
	after []int
}

var addr = 0
var addi = 1
var mulr = 2
var muli = 3
var banr = 4
var bani = 5
var borr = 6
var bori = 7
var setr = 8
var seti = 9
var gtir = 10
var gtri = 11
var gtrr = 12
var eqir = 13
var eqri = 14
var eqrr = 15

func (this inst) String() string {
	return fmt.Sprintf("<%d,%d,%d,%d>",this.opcode,this.op1,this.op2,this.res)
}

func (this test) String() string {
	return fmt.Sprintf("%v -> %s -> %v",this.before, this.inst, this.after)
}

var validBefore = regexp.MustCompile(`^Before:\s+\[(\d)+,\s+(\d)+,\s+(\d)+,\s+(\d)+\]$`)
var validInstruction = regexp.MustCompile(`^(\d)+\s+(\d)+\s+(\d)+\s+(\d)+$`)
var validAfter = regexp.MustCompile(`^After: +\[(\d)+, +(\d)+, +(\d)+, +(\d)+\]$`)

func exec(reg []int, i inst) {
	switch i.opcode{
	case addr:
		reg[i.res] = reg[i.op1] + reg[i.op2]

	case addi:
		reg[i.res] = reg[i.op1] + i.op2

	case mulr:
		reg[i.res] = reg[i.op1] * reg[i.op2]

	case muli:
		reg[i.res] = reg[i.op1] * i.op2

	case banr:
		reg[i.res] = reg[i.op1] & reg[i.op2]

	case bani:
		reg[i.res] = reg[i.op1] & i.op2

	case borr:
		reg[i.res] = reg[i.op1] | reg[i.op2]

	case bori:
		reg[i.res] = reg[i.op1] | i.op2

	case setr:
		reg[i.res] = reg[i.op1]

	case seti:
		reg[i.res] = i.op1

	case gtir:
		if i.op1 > reg[i.op2] {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	case gtri:
		if reg[i.op1] > i.op2 {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	case gtrr:
		if reg[i.op1] > reg[i.op2] {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	case eqir:
		if i.op1 == reg[i.op2] {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	case eqri:
		if reg[i.op1] == i.op2 {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	case eqrr:
		if reg[i.op1] == reg[i.op2] {
			reg[i.res] = 1
		} else {
			reg[i.res] = 0
		}

	}
}

func readTests() []test{
	scanner := bufio.NewScanner(os.Stdin)
	var tests []test
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) > 0 {
			// before
			match := validBefore.FindStringSubmatch(str);
			beforeRegs := make([]int,4)
			for i := 0; i < 4; i++{
				beforeRegs[i],_ = strconv.Atoi(match[i+1])
			}

			// instruction
			scanner.Scan()
			str = scanner.Text()
			match = validInstruction.FindStringSubmatch(str);
			opcode, _ := strconv.Atoi(match[1])
			op1, _ := strconv.Atoi(match[2])
			op2, _ := strconv.Atoi(match[3])
			res, _ := strconv.Atoi(match[4])
			instruction := inst{opcode, op1, op2, res}

			// after
			scanner.Scan()
			str = scanner.Text()
			match = validAfter.FindStringSubmatch(str);
			afterRegs := make([]int,4)
			for i := 0; i < 4; i++{
				afterRegs[i],_ = strconv.Atoi(match[i+1])
			}

			test:= test{beforeRegs, instruction, afterRegs}
			tests = append(tests,test)
		}
	}
	return tests
}

func equalMem(reg1,reg2 []int) bool{
	if len(reg1) != len(reg2) {
		return false
	} else {
		for i := 0; i < len(reg1); i++ {
			if reg1[i] != reg2[i] {
				return false
			}
		}
	}
	return true
}

func copyMem(reg []int) []int {
	duplicate := make([]int, 4)
	copy(duplicate, reg)
	return duplicate
}

func checkTest(test test) bool {
	dup := copyMem(test.before)
	exec(dup, test.inst)
	return equalMem(test.after,dup)
}

func main(){
	tests := readTests()
	var eliminatedCodes [16][16]bool // possibleCodes[i][j] = true if code i can NOT be op j
	for _, test := range tests {
		originalOpcode := test.inst.opcode
		for i := 0; i < 16; i++ {
			test.inst.opcode = i
			if !checkTest(test){
				if originalOpcode == 5 {
					fmt.Printf("not passed %v\n",test)
				}
				eliminatedCodes[originalOpcode][i] = true
			}
		}
	}
	//
	fmt.Println(eliminatedCodes)
	advanced := true
	for advanced {
		advanced = false
		for i := 0; i < 16; i++ {
			total := 0
			foundIndex := -1
			for j := 0; j < 16; j++ {
				if !eliminatedCodes[i][j] {
					foundIndex = j
					total++
				}
			}
			if (total == 0) {
				fmt.Println("PANIC")
				return
			}
			if (total == 1){
				advanced = true
				fmt.Printf("found that code %d is instruction %d\n",i,foundIndex)
				for k := 0; k < 16; k++ {
					if (k != i) {
						eliminatedCodes[k][foundIndex] = true
					}
				}
				fmt.Println(eliminatedCodes)
			}

		}
	}
	

}
