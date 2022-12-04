package day4

// TASK: https://adventofcode.com/2022/day/4

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Slot struct {
	Min, Max int
}

func overlap(s1, s2 Slot) bool {
	if s1.Min <= s2.Min && s1.Max >= s2.Max {
		return true
	}
	if s2.Min <= s1.Min && s2.Max >= s1.Max {
		return true
	}

	// Uncomment the code below for part two
	// if s1.Min >= s2.Min && s1.Min <= s2.Max {
	// 	return true
	// }

	// if s2.Min >= s1.Min && s2.Min <= s1.Max {
	// 	return true
	// }

	return false
}

func buildSlot(r []string) Slot {
	min, err := strconv.Atoi(r[0])
	if err != nil {
		panic(fmt.Sprintf(r[0], "is not a proper number"))
	}

	max, err := strconv.Atoi(r[1])
	if err != nil {
		panic(fmt.Sprintf(r[0], "is not a proper number"))
	}

	return Slot{
		Min: min,
		Max: max,
	}
}

func PartOne(scanner *bufio.Scanner) int {
	sum := 0
	var s1, s2 Slot
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		if len(pair) != 2 {
			panic(errors.New("Invalid input: each row should contain data for EXACLTY 2 elfs"))
		}

		r1 := strings.Split(pair[0], "-")
		if len(r1) != 2 {
			panic(errors.New("(2nd Split r1): Invalid input: each row should contain data for EXACLTY 2 elfs"))
		}

		r2 := strings.Split(pair[1], "-")
		if len(r1) != 2 {
			panic(errors.New("(2nd Split r2): Invalid input: each row should contain data for EXACLTY 2 elfs"))
		}

		s1 = buildSlot(r1)
		s2 = buildSlot(r2)

		fmt.Printf("\nOverlap between %v-%v -> %v\n", s1, s2, overlap(s1, s2))
		if overlap(s1, s2) {
			sum++
		}
	}

	return sum
}
