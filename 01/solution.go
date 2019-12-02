package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)

func readData(filename string) ([]int, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil { return nil, err }

    lines := strings.Split(string(data), "\n")
    sliceOfInt := make([]int, 0, len(lines))

    for _, l := range lines {
        if len(l) == 0 { continue }
        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }
        sliceOfInt = append(sliceOfInt, n)
    }

    return sliceOfInt, nil
}

func calculateFuel(mass int) int {
    return mass/3 -2
}

func part1(modules []int) int {
    output := 0
    for _, module := range modules{
        output += calculateFuel(module)
    }
    return output
}

func part2(modules []int) int {
    output := 0
    for _, mass := range modules{
        for {
            mass = calculateFuel(mass)
            if mass <= 0 {
                break
            }
            output += mass
        }
    }
    return output
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}

