package main

import (
    "bufio"
    "fmt"
    "log"
    "strings"
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

// Split backpack in 2 equal parts.
func splitCompartments(backpack string) (string, string) {
    comp1 := backpack[:len(backpack)/2]
    comp2 := backpack[len(backpack)/2:]
    return comp1, comp2
}

func checkDuplicatesCompartments(comp1 string, comp2 string) rune {
    for _, char := range comp1 {
        if strings.Contains(comp2, string(char)) {
            return char
        }
    }
    return -1
}

func checkDuplicatesBackpacks(bp1 string, bp2 string, bp3 string) rune {
    for _, char := range bp1 {
        if strings.Contains(bp2, string(char)) {
            if strings.Contains(bp3, string(char)){
                return char
            }
        }
    }
    return -1
}

func getPriority(char rune) int {
    ascii_value := int(char)
    priority := 0
    if ascii_value > 96 {
        priority = ascii_value - 96
    } else {
        priority = ascii_value - 38
    }
    return priority
}

func part1(backpacks []string) int {
    totalPriority := 0
    for _, backpack := range backpacks {
        comp1, comp2 := splitCompartments(backpack)
        duplicate := checkDuplicatesCompartments(comp1, comp2)
        priority := getPriority(duplicate)
        totalPriority = totalPriority + priority
    }
    return totalPriority
}

func part2(backpacks []string) int {
    totalPriority := 0
    for i := 0; i < len(backpacks); i = i + 3 {
        duplicate := checkDuplicatesBackpacks(backpacks[i], backpacks[i+1], backpacks[i+2])
        priority := getPriority(duplicate)
        totalPriority = totalPriority + priority 
    }
    return totalPriority
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}