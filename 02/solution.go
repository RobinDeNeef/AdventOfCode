package main

import (
    "fmt"
    "strconv"
    "strings"
)

func stringsToInts(sliceOfStrings []string) []int {
    var sliceOfInts = []int{}

    for _, i := range sliceOfStrings {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        sliceOfInts = append(sliceOfInts, j)
    }

    return sliceOfInts
}

func runIntCode(intcode []int, noun int, verb int) []int {
    memory := make([]int, len(intcode))
	copy(memory, intcode)
    cursor := 0
    memory[1] = noun
    memory[2] = verb
    for{
        opcode := memory[cursor]
		switch opcode {
		case 1:
			i1Pos := memory[cursor+1]
			i2Pos := memory[cursor+2]
			outputPos := memory[cursor+3]
			memory[outputPos] = memory[i1Pos] + memory[i2Pos]
		case 2:
			i1Pos := memory[cursor+1]
			i2Pos := memory[cursor+2]
			outputPos := memory[cursor+3]
			memory[outputPos] = memory[i1Pos] * memory[i2Pos]
		case 99:
			return memory
		default:
			fmt.Println("Received bad opcode: ", opcode)
		}
		cursor += 4
    }


func main() {
    input := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,5,19,23,1,6,23,27,1,27,10,31,1,31,5,35,2,10,35,39,1,9,39,43,1,43,5,47,1,47,6,51,2,51,6,55,1,13,55,59,2,6,59,63,1,63,5,67,2,10,67,71,1,9,71,75,1,75,13,79,1,10,79,83,2,83,13,87,1,87,6,91,1,5,91,95,2,95,9,99,1,5,99,103,1,103,6,107,2,107,13,111,1,111,10,115,2,10,115,119,1,9,119,123,1,123,9,127,1,13,127,131,2,10,131,135,1,135,5,139,1,2,139,143,1,143,5,0,99,2,0,14,0"
    intcode := stringsToInts(strings.Split(input, ","))
    fmt.Println(intcode)
    // Part 1
    output := runIntCode(intcode, 12, 2)
    fmt.Println(output[0])
    fmt.Println(intcode)
    // Part 2
    for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
            output := runIntCode(intcode, i, j)
			if output[0] == 19690720 {
				fmt.Println(i, "and", j, "=>", 100*i+j)
			}
		}
	}

}

