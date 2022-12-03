import (
    "bufio"
    "fmt"
    "log"
    "os"
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

func calculateWin(match []string) int {
	if match[0] == "A" && match[1] == "Y" {
		return 6
	} else if match[0] == "B" && match[1] == "Z" {
		return 6
	} else if match[0] == "C" && match[1] == "X" {
		return 6
	} else if match[0] == "A" && match[1] == "X" {
		return 3
	} else if match[0] == "B" && match[1] == "Y" {
		return 3
	} else if match[0] == "C" && match[1] == "Z" {
		return 3
	} else {
		return 0
	}
}

func handBonus(hand string) int {
	if hand == "X" {
		return 1
	} else if hand == "Y" {
		return 2
	} else {
		return 3
	}
}

func getScore(finish string) int {
	if finish == "X" {
		return 0
	} else if finish == "Y" {
		return 3
	} else {
		return 6
	}
}

func calculateHand(match []string) string {
	if match[0] == "A" && match[1] == "X" {
		return "Z"
	} else if match[0] == "A" && match[1] == "Y" {
		return "X"
	} else if match[0] == "A" && match[1] == "Z" {
		return "Y"
	} else if match[0] == "B" && match[1] == "X" {
		return "X"
	} else if match[0] == "B" && match[1] == "Y" {
		return "Y"
	} else if match[0] == "B" && match[1] == "Z" {
		return "Z"
	} else if match[0] == "C" && match[1] == "X" {
		return "Y"
	} else if match[0] == "C" && match[1] == "Y" {
		return "Z"
	} else if match[0] == "C" && match[1] == "Z" {
		return "X"
	} else {
		return "ERROR"
	}
}

func part1(games []string) int {
	totalScore := 0
    for _, game := range games{
		match := strings.Split(game, " ")
        score := calculateWin(match)
		score = score + handBonus(match[1])
		totalScore = totalScore + score
    }
    return totalScore
}

func part2(games []string) int {
	totalScore := 0
    for _, game := range games{
		match := strings.Split(game, " ")
        score := getScore(match[1])
		score = score + handBonus(calculateHand(match))
		totalScore = totalScore + score
    }
    return totalScore
}

func main() {
    data, _ := readData("input.txt")
    fmt.Println(part1(data))
    fmt.Println(part2(data))
}