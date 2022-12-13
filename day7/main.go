package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func isCommand(row string) bool {
	if strings.Contains(row, "$") {
		return true
	}
	return false
}

type FileSystem map[string]int

func (fs FileSystem) containsDir(path string) bool {
	for dirPath, _ := range fs {
		if dirPath == path {
			return true
		}
	}
	return false
}

func (fs *FileSystem) addDir(path string) error {
	if fs.containsDir(path) {
		return nil
	}
	(*fs)[path] = 0
	return nil
}

func (fs *FileSystem) increseRelatedDirsForFile(path string, fileSize int) {
	for k := range *fs {
		if k == "" || k == path {
			continue
		}

		if strings.Contains(path, k) {
			fmt.Println("))) a ", path, k, (*fs)[k], fileSize)
		}

	}
	(*fs)[path] += fileSize
}

func GoBack(path *string) {
	arr := strings.Split(*path, "/")
	newPath := strings.Join(arr[:len(arr)-1], "/")
	*path = newPath
}

func PartOne(scanner *bufio.Scanner) int {
	consoleOut := []string{}
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		if line == "$ cd /" {
			continue
		}
		consoleOut = append(consoleOut, line)
	}

	fileSystem := FileSystem{}
	currentDirPath := ""

	fileSystem[currentDirPath] = 0
	for _, row := range consoleOut {
		fmt.Println(row)
		if isCommand(row) && strings.Contains(row, "ls") {
			fmt.Println("Listing: ", currentDirPath)
			continue
		}
		if isCommand(row) && strings.Contains(row, "cd ..") {
			fmt.Println("Go back...", row)
			GoBack(&currentDirPath)
			fmt.Println("PWD: ", currentDirPath)
			continue
		}
		if isCommand(row) && strings.Contains(row, "cd") {
			fmt.Println("Change dir: ", row)
			arr := strings.Split(row, "cd ")
			if len(arr) != 2 {
				fmt.Println("error with cd command: ", row, len(arr))
				panic("cd error")
			}
			currentDirPath += fmt.Sprintf("/%s", arr[1])
			fileSystem[currentDirPath] = 0
			continue
		}
		if strings.Contains(row, "dir") {
			fmt.Println("List dir: ", row)
			arr := strings.Split(row, "dir ")
			if len(arr) != 2 {
				fmt.Println("error with list dir command: ", row, len(arr))
				panic("list dir error")
			}
			dirPath := fmt.Sprintf("%s/%s", currentDirPath, arr[1])
			if fileSystem.containsDir(dirPath) {
				fileSystem.addDir(dirPath)
			}
		} else {
			fmt.Println("List file: ", row)
			arr := strings.Split(row, " ")
			if len(arr) != 2 {
				fmt.Println("error with list file command: ", row, arr, len(arr))
				panic("list file error")
			}

			_, found := fileSystem[currentDirPath]
			if !found {
				panic("non existing dir on list files")
			}

			fileSize, err := strconv.Atoi(arr[0])
			if err != nil {
				panic(err)
			}
			fileSystem[currentDirPath] += fileSize
		}
	}

	for k, _ := range fileSystem {
		for ki, _ := range fileSystem {
			if k == ki || ki == "" || k == "" {
				continue
			}
			if strings.Contains(k, ki) && len(k) > len(ki) {
				fileSystem[ki] += fileSystem[k]
			}

		}
	}
	sumUnder100k := 0
	for _, v := range fileSystem {
		if v <= 100_000 {
			sumUnder100k += v
		}
	}
	return sumUnder100k
}
