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

	for i := range s1 {
		for j := range s2 {
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

func PartOne(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		p1, p2 := scanner.Text()[:(len(scanner.Text())/2)], scanner.Text()[(len(scanner.Text())/2):]
		commonLetters := string(findCommonLetters([]byte(p1), []byte(p2)))
		sum += letterToPoints(commonLetters[0])
	}

	return sum
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findCommonLettersInTriplet(triplet []string) []byte {
	for i, chunk := range triplet {
		triplet[i] = sortString(chunk)
	}

	f, s, t := 0, 0, 0
	var fVal, sVal, tVal int32

	for f < len(triplet[0]) || s <= len(triplet[1]) || t < len(triplet[2]) {
		fVal = rune(triplet[0][f])
		sVal = rune(triplet[1][s])
		tVal = rune(triplet[2][t])

		if fVal == sVal && sVal == tVal {
			fmt.Println("Match: ", string(fVal))
			return []byte{triplet[0][f]}
		}

		if fVal < sVal {
			f++
		} else if sVal < tVal {
			s++
		} else {
			t++
		}
	}
	return []byte{}
}

func PartTwo(scanner *bufio.Scanner) int {
	sum, idx := 0, 0
	triplet := []string{}

	for scanner.Scan() {
		idx++
		triplet = append(triplet, scanner.Text())

		if idx == 3 {
			commonLetters := findCommonLettersInTriplet(triplet)
			sum += letterToPoints(commonLetters[0])
			triplet = []string{}
			idx = 0
		}
	}

	return sum
}
