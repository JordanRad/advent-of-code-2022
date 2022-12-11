package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JordanRad/advent-of-code-2022/day6"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanFile(filePath string) *bufio.Scanner {
	file, err := os.Open(filePath)
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func main() {
	fmt.Println(day6.PartOne(ScanFile("day6/input.txt")))
}
