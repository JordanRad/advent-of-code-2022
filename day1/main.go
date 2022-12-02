package day1

// TASK: https://adventofcode.com/2022/day/1

import (
	"bufio"
	"sort"
	"strconv"
)

func LevelOne(scanner *bufio.Scanner) int {
	max, curr := 0, 0

	for scanner.Scan() {
		entry := string(scanner.Text())
		if entry == "" {
			if curr > max {
				max = curr
			}
			curr = 0
			continue
		}
		num, err := strconv.Atoi(entry)
		if err != nil {
			panic(err)
		}
		curr += num
	}

	return max
}

func arrSum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func isGreaterOnIndex(arr []int, number int) (bool, int) {
	for i := range arr {
		if number > arr[i] {
			return true, i
		}
	}
	return false, 0
}

func LevelTwo(scanner *bufio.Scanner) int {
	maxArr, curr := [3]int{0, 0, 0}, 0

	for scanner.Scan() {
		entry := string(scanner.Text())

		if entry == "" {
			sort.Ints(maxArr[:])
			if isGrater, idx := isGreaterOnIndex(maxArr[:], curr); isGrater {
				maxArr[idx] = curr
			}

			curr = 0
			continue
		}

		num, err := strconv.Atoi(entry)
		if err != nil {
			panic(err)
		}
		curr += num
	}

	return arrSum(maxArr[:])
}
