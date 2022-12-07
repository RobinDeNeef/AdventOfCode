package main

import (
    "bufio"
    "fmt"
    "log"
    // "sort"
    "os"
    // "strconv"
	"strings"
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

func hasDuplicateCharacters(text string) bool {
	for i := 0; i < len(text); i++ {
		count := strings.Count(text, string(text[i]))
		if count > 1 {
			return true
		}
	}
	return false
}

func part1(lines []string) int {
	for _, line := range lines {
		// previousChars := line[:2]
		for i := 3; i < len(line); i++ {
			fmt.Println(line[i-3: i + 1],"-" ,string(line[i]))
			if !hasDuplicateCharacters(line[i-3: i + 1]){
				return i + 1
			}
		}
	} 
	
	return 1 
}

func part2(lines []string) int {
	for _, line := range lines {
		// previousChars := line[:2]
		for i := 13; i < len(line); i++ {
			fmt.Println(line[i-13: i + 1],"-" ,string(line[i]))
			if !hasDuplicateCharacters(line[i-13: i + 1]){
				return i + 1
			}
		}
	} 
	
	return 1 
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}