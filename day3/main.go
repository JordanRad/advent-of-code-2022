package day3

// TASK: https://adventofcode.com/2022/day/3

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func findCommonLetters(s1, s2 []byte) []byte {
	sort.Slice(s1, func(i, j int) bool { return s1[i] > s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] > s2[j] })

	t1, t2 := 0, 0

	fmt.Println(s1, s2)
	for i := range s1 {
		fmt.Println("t", t1, t2)
		for j := range s2 {
			fmt.Println("iteration: ", string(s1[i]), string(s2[j]))
			if s1[i] == s2[j] {
				return []byte{s1[i]}
			}
		}

	}
	return []byte{}
}

func letterToPoints(letter byte) int {

	isLowerCase := func(letter byte) bool {
		return letter == byte(unicode.ToLower(rune(letter)))
	}

	for i := range alphabet {
		lowerCaseLetter := []byte(strings.ToLower(string(letter)))[0]
		if lowerCaseLetter == alphabet[i] {
			points := i + 1
			if isLowerCase(letter) {
				return points
			}
			return points + 26
		}
	}
	return -1
}

func LevelOne(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		p1, p2 := scanner.Text()[:(len(scanner.Text())/2)], scanner.Text()[(len(scanner.Text())/2):]
		commonLetters := string(findCommonLetters([]byte(p1), []byte(p2)))
		sum += letterToPoints(commonLetters[0])
	}

	return sum
}
