package main

import (
	"fmt"
    "strings"
	"strconv"
)

func validate(number string) bool {
	prev := rune(number[0])
	foundDoubles := false
	for _, r := range number[1:] {
		if prev > r {
			return false
		} else if prev == r {
			foundDoubles = true
		}
		prev = r
	}
	return foundDoubles
}

func foundMoreThanDoubles(number string) bool {
	for _, r := range number {
		if strings.Count(number, string(r)) == 2 {
			return false
		} 
	}
	return true
}

func part1(start int, stop int) []int  {
	possibilities := []int{}
	for i:= start; i < stop; i++{
		if validate(strconv.Itoa(i)){
			possibilities = append(possibilities, i)
		} 
	}
	return possibilities
}

func part2(possibilities []int) []int  {
	narrowed_possibilities := []int{}
	for _, i := range possibilities{
		if !foundMoreThanDoubles(strconv.Itoa(i)){
			narrowed_possibilities = append(narrowed_possibilities, i)
		} 
	}
	return narrowed_possibilities
}

func main() {
	start := 382345
	stop := 843167
	possibilities := part1(start, stop)
	fmt.Println("Part 1: ", len(possibilities))
	fmt.Println("Part 1: ", len(part2(possibilities)))
}

