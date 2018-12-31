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

var opNames = []string{"addr","addi","mulr","muli","banr","bani","borr","bori","setr","seti","gtir","gtri","gtrr","eqir","eqri","eqrr"}

func opNameToOpCode(opName string) int {
	for pos, name := range opNames{
		if opName == name {
			return pos
		}
	}
	panic("no opname found")
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

var validInstruction = regexp.MustCompile(`^(.+)\s+(\d+)\s+(\d+)\s+(\d+)$`)
var setInstRegister = regexp.MustCompile(`^#ip (\d+)$`)

func exec(reg []int, i inst) []int{
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
	default:
		fmt.Println("FUCK YOU")
	}
	return reg
}

func readProgram() (int, []inst){
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []inst
	scanner.Scan()
	firstLine := scanner.Text()
	match := setInstRegister.FindStringSubmatch(firstLine);
	instructionRegister, _ := strconv.Atoi(match[1])
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) > 0 {
			str = scanner.Text()
			match = validInstruction.FindStringSubmatch(str);
			opcode := opNameToOpCode(match[1])
			op1, _ := strconv.Atoi(match[2])
			op2, _ := strconv.Atoi(match[3])
			res, _ := strconv.Atoi(match[4])
			instruction := inst{opcode, op1, op2, res}
			instructions = append(instructions, instruction)
		}
	}
	return instructionRegister, instructions
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

func memToString(mem []int) string {
	return fmt.Sprintf("%v",mem)
}

func main(){
	instReg, instructions := readProgram()
	instPointer := 0
	seen := make(map[int]bool)
	memory := make([]int, 6)
	i := 0
	for instPointer >= 0 && instPointer < len(instructions) {
		//fmt.Println(memory)
		memory[instReg] = instPointer
		memory = exec(memory, instructions[instPointer])
		instPointer = memory[instReg]
		instPointer++
		if instPointer == 28 {
			i++;
			fmt.Printf("(%d): %v\n",i,memory[5])
			if seen[memory[5]] {
				break
			} else {
				seen[memory[5]] = true
			}
		}

	}
	fmt.Printf("answer: %v\n",memory[5])
}
