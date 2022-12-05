package main

import (
    "bufio"
    "fmt"
    "log"
    "strings"
    "strconv"
    "os"
)

func readData(filename string) ([]string, error) {
    file, err := os.Open(filename)

    if err != nil { 
        log.Fatalf("failed opening files: %s", err)
        return nil, err 
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var txtLines []string 

    for scanner.Scan() {
        txtLines = append(txtLines, scanner.Text())
    }

    file.Close()

    return txtLines, nil
}

func printSlice(stackSlice [][]string) {
	fmt.Println("-----")
	for _, line := range stackSlice{
		fmt.Println("[" + strings.Join(line, ",") + "]")
	}
	fmt.Println("-----")
}

func parseStacks(stacks []string) [][]string{
	var stackSlice [][]string

	for i := len(stacks) -1; i >= 0; i-- {
		stackId := 0
		for charId := 1; charId < len(stacks[i]); charId += 4 {
			if i == len(stacks) - 1{
				stackSlice = append(stackSlice, make([]string, 0))
			}
			char := string(stacks[i][charId])
			if char != " " {
				stackSlice[stackId] = append(stackSlice[stackId], char)
			}
			stackId++
		}
	}

	return stackSlice
}

func popSlice(stringSlice []string, popSize int) ([]string, []string) {
	return stringSlice[:len(stringSlice)-popSize], stringSlice[len(stringSlice) - popSize:]
}


func parseCommand(command string) (int, int, int) {
	commandParts := strings.Split(command, " ")
	qty, _ := strconv.Atoi(commandParts[1])
	from, _ := strconv.Atoi(commandParts[3])
	to, _ := strconv.Atoi(commandParts[5])
	return qty, from, to
}

func executeCommand(stackSlice [][]string, qty int, from int, to int) [][]string {
	for i := 0; i < qty; i++ {
		newStack, poppedItems := popSlice(stackSlice[from-1], 1)
		stackSlice[from-1] = newStack
		stackSlice[to-1] = append(stackSlice[to-1], poppedItems...)
	}
	return stackSlice
}

func executeCrateMover9001(stackSlice [][]string, qty int, from int, to int) [][]string {
	newStack, poppedItems := popSlice(stackSlice[from-1], qty)
	stackSlice[from-1] = newStack
	stackSlice[to-1] = append(stackSlice[to-1], poppedItems...)

	return stackSlice
}

func getTopOfStacks(stackSlice [][]string) string {
	topStack := ""
	for _, stack := range stackSlice {
		topStack = topStack + stack[len(stack)-1]
	}
	return topStack
}

func part1(lines []string) string {
	var stacks []string
	var commands []string
	for _, line := range lines {
		if strings.Contains(line, "[") {
			stacks = append(stacks, line)
		} else if strings.Contains(line, "move") {
			commands = append(commands, line)
		}
	}
	stackSlice := parseStacks(stacks)

	for _, command := range commands {
		qty, from, to := parseCommand(command)
		executeCommand(stackSlice, qty, from, to)
	}
	printSlice(stackSlice)
	return getTopOfStacks(stackSlice)
}

func part2(lines []string) string {
	var stacks []string
	var commands []string
	for _, line := range lines {
		if strings.Contains(line, "[") {
			stacks = append(stacks, line)
		} else if strings.Contains(line, "move") {
			commands = append(commands, line)
		}
	}
	stackSlice := parseStacks(stacks)
	for _, command := range commands {
		qty, from, to := parseCommand(command)
		executeCrateMover9001(stackSlice, qty, from, to)
	}
	printSlice(stackSlice)
	return getTopOfStacks(stackSlice)
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}