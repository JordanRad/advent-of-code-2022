package day6

import (
	"bufio"
	"fmt"
)

// TASK: https://adventofcode.com/2022/day/6

func isValidMarker(chunk string) bool {
	hashmap := map[int32]int{}

	for _, item := range chunk {
		fmt.Println(hashmap[item] == 0)
		if hashmap[item] == 0 {
			hashmap[item] = 1
			continue
		}
		return false
	}

	return true
}

type InitialChunk struct {
	Value           string
	IndexOfLastChar int
}

func PartOne(scanner *bufio.Scanner) int {
	message := ""
	for scanner.Scan() {
		message = scanner.Text()
	}

	t := ""
	validMarker := InitialChunk{
		Value:           "",
		IndexOfLastChar: -1,
	}
	for i := 0; i < len(message); i++ {
		fmt.Println(string(message[i]))
		t += string(message[i])
		fmt.Println("val: ", t)
		// PART ONE
		// if len(t) == 4 {
		// 	if !isValidMarker(t) {
		// 		t = t[1:]
		// 	}
		// 	validMarker.Value = t
		// 	validMarker.IndexOfLastChar = i + 1
		// }

		// PART TWO
		if len(t) == 14 {
			if !isValidMarker(t) {
				t = t[1:]
			}
			validMarker.Value = t
			validMarker.IndexOfLastChar = i + 1
		}

	}
	fmt.Println(validMarker)

	return -1
}
