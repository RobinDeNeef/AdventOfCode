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

// Bad Solution because range 11-12 will match range 1-2
// func fillRange(min string, max string) string {
//     rangeString := ""
//     minInt, _ := strconv.Atoi(min)
//     maxInt, _ := strconv.Atoi(max)
//     for i := minInt; i <= maxInt; i++ {
//         rangeString = rangeString + strconv.Itoa(i)
//     }
//     return rangeString
// }

func containsRange(min1 int, max1 int, min2 int, max2 int ) bool{
    if min1 <= min2 && max1 >= max2 {
        return true
    }
    return false
}

func notContainsRange(min1 int, max1 int, min2 int, max2 int ) bool{
    if max1 < min2 || max2 < min1 {
        print()
        return true
    }
    return false
}


func part1(lines []string) int {
    totalCount := 0
    for _, line := range lines{
        pairs := strings.Split(line, ",")
        minMax1 := strings.Split(pairs[0], "-")
        minMax2 := strings.Split(pairs[1], "-")
        min1, _ := strconv.Atoi(minMax1[0])
        max1, _ := strconv.Atoi(minMax1[1])
        min2, _ := strconv.Atoi(minMax2[0])
        max2, _ := strconv.Atoi(minMax2[1])
        if containsRange(min1, max1, min2, max2) || containsRange(min2, max2, min1, max1){
            totalCount++
        }
    } 
    return totalCount
}

func part2(lines []string) int {
    totalCount := 0
    for _, line := range lines{
        pairs := strings.Split(line, ",")
        minMax1 := strings.Split(pairs[0], "-")
        minMax2 := strings.Split(pairs[1], "-")
        min1, _ := strconv.Atoi(minMax1[0])
        max1, _ := strconv.Atoi(minMax1[1])
        min2, _ := strconv.Atoi(minMax2[0])
        max2, _ := strconv.Atoi(minMax2[1])
        if  !notContainsRange(min1, max1, min2, max2){
            fmt.Println(minMax1, minMax2)
            totalCount++
        }
    } 
    return totalCount
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}