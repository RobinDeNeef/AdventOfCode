package main

import(
	"os"
	"log"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// Global Variable to pick the map sizes out of the recursive function
var FOLDER_SIZES = map[string]int{}

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

type file struct {
	name string
	size int
}

type folder struct {
	absolutePath string
	parentFolder *folder
	folders []*folder
	files []file
}

func (f folder) calculateFolderSize() int {
	totalSize := 0
	for _, nestedFolder := range f.folders {
		folderSize := nestedFolder.calculateFolderSize()
		totalSize += folderSize
	}
	for _, file := range f.files {
		totalSize += file.size
	}
	// fmt.Println(f.absolutePath, totalSize)
	FOLDER_SIZES[f.absolutePath] = totalSize
	return totalSize
}

func buildFileSystem(data []string) *folder {
	root := folder{absolutePath: "/"}
	var currentPath string = "/"
	var workingDirectory *folder
	cursor := 0
	for cursor < len(data) {
		line := strings.Split(data[cursor], " ") 
		if string(line[0]) == "$" {
			// fmt.Println(line)
			if line[1] == "cd" {
				if line[2] == "/" {
					workingDirectory = &root
					currentPath = "/"
				} else if line[2] == ".." {
					currentPath = currentPath[:strings.LastIndex(currentPath, "/") + 1]
					workingDirectory = workingDirectory.parentFolder
				} else {
					newFolder := folder{absolutePath: currentPath + "/" + line[2]}
					workingDirectory.folders = append(workingDirectory.folders, &newFolder)
					newFolder.parentFolder = workingDirectory
					workingDirectory = &newFolder
					currentPath = currentPath + line[2] + "/"
				}
			}
			// fmt.Println(root)
			// fmt.Println(currentPath)
			// fmt.Println(workingDirectory)
		} else if line[0] != "dir" {
			size, _ := strconv.Atoi(line[0])
			newFile := file{name: line[1], size: size}
			workingDirectory.files = append(workingDirectory.files, newFile)
		}
		// fmt.Println("---")
		cursor++
	}
	// newFile := file{name: "test.txt", size: 123}
	// root.files = append(root.files, newFile)
	// fmt.Println(root.calculateFolderSize())
	return &root
}

func printFolder(f *folder, depth int ) {
	depthString := ""
	for i := 0; i < depth; i++  {
		depthString += "  "
	}
	fmt.Println(depthString + f.absolutePath, "[DIR]")
	for _, nestedFolder := range f.folders {
		printFolder(nestedFolder, depth + 1)
	}
	for _, file := range f.files {
		fmt.Println(depthString + file.name, "(", file.size, ")" )
	}
}

func part1(folderSizes map[string]int) int {
	solution := 0
	for _, size := range folderSizes {
		if size <= 100000 {
			solution += size
		}
	}
	return solution
}

func part2(folderSizes map[string]int) int {
	freeSpace := 70000000 - folderSizes["/"]
	smallestMap := "/"
	leftOverStorage := freeSpace
	for folderName, size := range folderSizes {
		// Iterate over every folder and check which one will 
		// have the least leftover space once the update is installed
		if freeSpace + size > 30000000 {
			if freeSpace + size - 30000000 < leftOverStorage {
				leftOverStorage = freeSpace + size - 30000000
				smallestMap = folderName
			}
		}
	}
	return folderSizes[smallestMap]
}

func main() {
    data, _ := readData("input.txt")
	fileSystem := buildFileSystem(data)
	totalSize := fileSystem.calculateFolderSize()
	fmt.Println(totalSize)
    fmt.Println(part1(FOLDER_SIZES))
    fmt.Println(part2(FOLDER_SIZES))
}