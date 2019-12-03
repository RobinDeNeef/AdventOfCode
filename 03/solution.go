package main

import (
//     "strconv"
		"fmt"
		"math"
		"strings"
		"strconv"
    "io/ioutil"
)

type Coordinates struct {
	x, y int
}

func readData(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil { return nil, err }
    
	lines := strings.Split(string(data), "\n")
    
	return lines, nil
}

func getCoordinates(directions []string) []Coordinates {
	coordinates := make([]Coordinates, 0)
	currentCoordinates := Coordinates{0,0}
	for _, instruction := range directions{
		direction := instruction[:1]
		distance, _ := strconv.Atoi(instruction[1:])
		switch direction{
		case "U":
			for i := 0; i < distance; i++ {
				currentCoordinates.y += 1
				coordinates = append(coordinates, currentCoordinates)
			}
		case "D":
			for i := 0; i < distance; i++ {
				currentCoordinates.y -= 1
				coordinates = append(coordinates, currentCoordinates)
			}
		case "R":
			for i := 0; i < distance; i++ {
				currentCoordinates.x += 1
				coordinates = append(coordinates, currentCoordinates)
			}
		case "L":
			for i := 0; i < distance; i++ {
				currentCoordinates.x -= 1
				coordinates = append(coordinates, currentCoordinates)
			}
		}
	}
	return coordinates
}

func findCollisions(wirePositions []([]Coordinates)) []Coordinates {
	var collisions []Coordinates
	for _, coorA := range wirePositions[0] {
		for _, coorB := range wirePositions[1] {
			if coorA == coorB {
				collisions = append(collisions, coorA)
			}
		}
	}
	return collisions
}

func manhattenDistance(coor Coordinates) int{
	distance := math.Abs(float64(coor.x)) + math.Abs(float64(coor.y))
	fmt.Println(coor, " -> ",distance)
	return int(distance)
}

func stepsToCollision(wirePositions []([]Coordinates), collisions []Coordinates) string{
	
	for _, collision := range collisions{
		stepsA := 0
		stepsB := 0
		for steps, coorA := range wirePositions[0] {
			if coorA == collision{
				stepsA = steps + 1
				break
			}
		}
		for steps, coorB := range wirePositions[1] {
			if coorB == collision{
				stepsB = steps + 1
				break
			}
		}
		fmt.Println(collision, "reached after", stepsA, "+", stepsB, "-->" , stepsA+stepsB)
	}
	return "nil"
}


func main() {
	input,_ := readData("input.txt")
	var wirePositions []([]Coordinates)
	for _, line := range input{
		directions := strings.Split(line, ",")
		wireCoordinates := getCoordinates(directions)
		wirePositions = append(wirePositions, wireCoordinates)
	}
	collisions := findCollisions(wirePositions)

	// PART 1
	// for _, collision := range collisions{
	// 	manhattenDistance(collision)
	// }

	// PART 2
	fmt.Println(stepsToCollision(wirePositions, collisions))
}

