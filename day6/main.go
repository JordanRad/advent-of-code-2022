package day6

import (
	"bufio"
)

// TASK: https://adventofcode.com/2022/day/6

func isValidMarker(chunk string) bool {
	hashmap := map[int32]int{}
	for _, item := range chunk {
		if hashmap[item] == 0 {
			hashmap[item] = 1
			continue
		}
		return false
	}
	return true
}

func PartOne(scanner *bufio.Scanner) int {
	message := ""
	for scanner.Scan() {
		message = scanner.Text()
	}

	t := ""

	idx := 0
	for i := 0; i < len(message); i++ {
		t += string(message[i])
		// Change the condition t len(t) == 14 for part 2
		if len(t) == 4 {
			if !isValidMarker(t) {
				t = t[1:]
			}
			idx = i + 1
		}
	}
	return idx
}
