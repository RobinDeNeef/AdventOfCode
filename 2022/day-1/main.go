package main

import (
    "bufio"
    "fmt"
    "log"
    "sort"
    "os"
    "strconv"
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

func convertLinesToInt(sliceOfString []string) []int {
    var sliceOfInt []int
    for _, textLine := range sliceOfString {
        number, _ := strconv.Atoi(textLine)
        sliceOfInt = append(sliceOfInt, number)
    }
    return sliceOfInt
}

func countIntInSlice(sliceOfInt []int) int {
    total := 0
    for _, number := range sliceOfInt {
        total = total + number
    }
    return total
}

func splitOnEmptyValue(sliceOfStrings []string) [][]string  {
    var sliceOfStringSlices [][]string
    lastChunck := 0
    for i, value := range sliceOfStrings{
        if value == "" {
            sliceOfStringSlices = append(sliceOfStringSlices, sliceOfStrings[lastChunck:i])
            // fmt.Println("|"+ strings.Join(sliceOfStrings[lastChunck:i], ",")+"|")
            lastChunck = i+1
        }
    }
    return sliceOfStringSlices
}

func part1(snacks []string) int {
    var snacksByElf [][]string = splitOnEmptyValue(snacks)
    highestCalorieCount := 0

    for _, elf := range snacksByElf{
        totalCalories := countIntInSlice(convertLinesToInt(elf))
        if totalCalories > highestCalorieCount {
            highestCalorieCount = totalCalories
        }
    }
    return highestCalorieCount
}

func part2(snacks []string) int {
    var snacksByElf [][]string = splitOnEmptyValue(snacks)
    var totalCaloriesByElf []int 

    for _, elf := range snacksByElf{
        totalCalories := countIntInSlice(convertLinesToInt(elf))
        totalCaloriesByElf = append(totalCaloriesByElf, totalCalories)
    }
    sort.Ints(totalCaloriesByElf)

    totalTopThree := countIntInSlice(totalCaloriesByElf[len(totalCaloriesByElf)-3:])

    return totalTopThree
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}