package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operand1, operand2 string
	operator           string
	output             string
}

func parse_wire(unparsed_wires []string) map[string]bool {
	wire := make(map[string]bool)
	for _, unparsed_wire := range unparsed_wires {
		var name string
		var signal bool
		name = strings.Split(unparsed_wire, ": ")[0]
		if strings.Split(unparsed_wire, ": ")[1] == "1" {
			signal = true
		} else {
			signal = false
		}
		wire[name] = signal
	}
	return wire
}

func parse_instruction(unparsed_instructions []string) []Instruction {
	var instructions []Instruction
	for _, unparsed_instruction := range unparsed_instructions {
		var current_instruction Instruction
		unparsed_instruction := strings.Split(unparsed_instruction, " ")
		current_instruction.operand1 = unparsed_instruction[0]
		current_instruction.operand2 = unparsed_instruction[2]
		current_instruction.operator = unparsed_instruction[1]
		current_instruction.output = unparsed_instruction[4]
		instructions = append(instructions, current_instruction)
	}
	return instructions
}

func gates(operand1 bool, operand2 bool, gate string) bool {
	var result bool
	switch gate {
	case "AND":
		result = operand1 && operand2
	case "XOR":
		result = (operand1 || operand2) && !(operand1 && operand2)
	case "OR":
		result = operand1 || operand2
	}
	return result
}

func retrive(instructions []Instruction, output string) int {
	for i, instruction := range instructions {
		if output == instruction.output {
			return i
		}
	}
	return -1
}

func execute(wire map[string]bool, instructions []Instruction, index int) {
	operand1, do_operand1_exists := wire[instructions[index].operand1]
	operand2, do_operand2_exists := wire[instructions[index].operand2]
	if do_operand1_exists && do_operand2_exists {
		wire[instructions[index].output] = gates(operand1, operand2, instructions[index].operator)
	} else {
		if !do_operand1_exists {
			execute(wire, instructions, retrive(instructions, instructions[index].operand1))
		}
		if !do_operand2_exists {
			execute(wire, instructions, retrive(instructions, instructions[index].operand2))
		}
		operand1 = wire[instructions[index].operand1]
		operand2 = wire[instructions[index].operand2]
		wire[instructions[index].output] = gates(operand1, operand2, instructions[index].operator)
	}
}

func part1(wire map[string]bool, instructions []Instruction) int {
	result := 0
	for i := range instructions {
		execute(wire, instructions, i)
	}
	for key, value := range wire {
		if strings.HasPrefix(key, "z") && value {
			position, _ := strconv.Atoi(key[1:])
			result += int(math.Pow(2, float64(position)))
			fmt.Println(key, value)
		}
	}
	fmt.Println(len(wire))
	return result
}

func main() {
	file, _ := os.ReadFile("input.txt")
	unparsed_wires := strings.Split(strings.Split(string(file), "\n\n")[0], "\n")
	unparsed_instructions := strings.Split(strings.Split(string(file), "\n\n")[1], "\n")

	wire := parse_wire(unparsed_wires)
	instructions := parse_instruction(unparsed_instructions)
	// fmt.Println(wire)
	// fmt.Println(instructions)
	fmt.Printf("Part1 Ans: %d\n", part1(wire, instructions))
}
