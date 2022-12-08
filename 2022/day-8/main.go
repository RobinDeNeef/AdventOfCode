package main


import(
	"os"
	"log"
	"fmt"
	"bufio"
	// "strings"
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

func buildGrid(data []string) [][]int {
	grid := make([][]int, len(data))
	for x, line := range data {
		for _, char := range line {
			number, _ := strconv.Atoi(string(char))
			grid[x] = append(grid[x], number)
		}
	}
	return grid
}

func getColumn(grid [][]int, columnIdx int) []int {
	column := make([]int, 0)
    for _, row := range grid {
        column = append(column, row[columnIdx])
    }
    return column
}

func checkHigherTree(treeHeight int, intSlice []int) bool{
	for _, height := range intSlice {
		if height >= treeHeight {
			return true
		}
	} 
	return false
}

func measureHigherTreeDistance(treeHeight int, intSlice []int) int{
	distance := 1
	for idx, height := range intSlice {
		if height >= treeHeight {
			return distance
		} else {
			if idx + 1 < len(intSlice){ // Otherwise counts "out of bounds" as extra tree
				distance += 1
			}
		}
	} 
	return distance
}

func reverseSlice(slice []int) []int {
	revSlice := make([]int, 0)
	for idx := range slice {
		revSlice = append(revSlice, slice[len(slice)-1-idx])
	}
	return revSlice
}

func checkVisability(x int, y int, row []int, column []int) bool {
	left := checkHigherTree(row[y], row[:y])
	right := checkHigherTree(row[y], row[y+1:])
	top := checkHigherTree(column[x], column[:x])
	bottom := checkHigherTree(column[x], column[x+1:])

	return (left && right && top && bottom)
}

func calculateScenicScore(x int, y int, row []int, column []int) int {
	left := measureHigherTreeDistance(row[y], reverseSlice(row[:y]))
	right := measureHigherTreeDistance(row[y], row[y+1:])
	top := measureHigherTreeDistance(column[x], reverseSlice(column[:x]))
	bottom := measureHigherTreeDistance(column[x], column[x+1:])

	return (left * right * top * bottom)
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func part1(grid [][]int) int {
	visibilityCounter := 0
	for x, row := range grid {
		for y, _ := range row {
			// Detect Edge
			if x == 0 || x == len(grid) -1 || y == 0 || y == len(row) -1{  
				visibilityCounter += 1
			} else {
				column := getColumn(grid, y)
				if !checkVisability(x, y, row, column) {
					visibilityCounter += 1	
				}
			}
			
		}
	}
	return visibilityCounter
}


func part2(grid [][]int) int {
	maxScenicScore := 0
	for x, row := range grid {
		for y, _ := range row {
			column := getColumn(grid, y)
			if !(x == 0 || x == len(grid) -1 || y == 0 || y == len(row) -1) { // Skip edges
				newScore := calculateScenicScore(x, y, row, column)
				if newScore > maxScenicScore {
					maxScenicScore = newScore
				}
			}		
		}
	}
	return maxScenicScore
}


func main() {
    data, _ := readData("input.txt")
	grid := buildGrid(data)
	// printGrid(grid)
	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}