package day2

import (
	"bufio"
	"errors"
	"fmt"
)

type Player struct {
	Name   string
	Inputs []Option
	Points int
}

type Option struct {
	Name           OptionName
	Index          int
	RequiredResult int
}

type OptionName string

const (
	RockOption     OptionName = "rock"
	PaperOption    OptionName = "paper"
	ScissorsOption OptionName = "scissors"
	DefaultOption  OptionName = "default"
)

func getOptionFromLetter(letter string) Option {
	if letter == "A" || letter == "X" {
		return Option{
			Name:  RockOption,
			Index: 0,
		}
	}
	if letter == "B" || letter == "Y" {
		return Option{
			Name:  PaperOption,
			Index: 1,
		}
	}
	if letter == "C" || letter == "Z" {
		return Option{
			Name:  ScissorsOption,
			Index: 2,
		}
	}
	return Option{
		Name:  DefaultOption,
		Index: -10,
	}

}

func getRoundPoints(op1, op2 Option) (int, int) {
	if op1.Name == DefaultOption || op2.Name == DefaultOption {
		return -1, -1
	}
	var outcomes [][]int = [][]int{
		{3, 0, 6},
		{6, 3, 0},
		{0, 6, 3},
	}
	// Index of the selected base to compare to + attribute points (r-1, p-2, s-3)
	return outcomes[op1.Index][op2.Index] + (op1.Index + 1), outcomes[op2.Index][op1.Index] + (op2.Index + 1)
}

func LevelOne(scanner *bufio.Scanner) int {
	var p1 = Player{
		Name:   "Opponent",
		Inputs: []Option{},
		Points: 0,
	}
	var p2 = Player{
		Name:   "Me",
		Inputs: []Option{},
		Points: 0,
	}

	op1 := Option{
		Name:  DefaultOption,
		Index: -1,
	}
	op2 := Option{
		Name:  DefaultOption,
		Index: -1,
	}
	for scanner.Scan() {
		op1 = getOptionFromLetter(string(scanner.Text()[0]))
		op2 = getOptionFromLetter(string(scanner.Text()[2]))

		p1.Inputs = append(p1.Inputs, op1)
		p2.Inputs = append(p2.Inputs, op2)

		r1, r2 := getRoundPoints(op1, op2)
		fmt.Println(r1, r2)
		p1.Points += r1
		p2.Points += r2

		fmt.Printf("%s : %s %s : %s \n", string(scanner.Text()[0]), string(op1.Name), string(scanner.Text()[2]), string(op2.Name))
		fmt.Printf("Points: \n P1: %d P2: %d", p1.Points, p2.Points)
		fmt.Println("")
	}

	return p2.Points
}

// Counts only for level 2 of the task
func getOpponentOption(option string) Option {
	if option == "A" {
		return Option{
			Name:  RockOption,
			Index: 0,
		}
	}
	if option == "B" {
		return Option{
			Name:  PaperOption,
			Index: 1,
		}
	}
	if option == "C" {
		return Option{
			Name:  ScissorsOption,
			Index: 2,
		}
	}
	return Option{
		Name:  DefaultOption,
		Index: -10,
	}
}

// Counts only for level 2 of the task
func getMyOption(option string) Option {
	if option == "X" {
		return Option{
			Name:           RockOption,
			Index:          0,
			RequiredResult: 0,
		}
	}
	if option == "Y" {
		return Option{
			Name:           PaperOption,
			Index:          1,
			RequiredResult: 1,
		}
	}
	if option == "Z" {
		return Option{
			Name:           ScissorsOption,
			Index:          2,
			RequiredResult: 2,
		}
	}
	return Option{
		Name:  DefaultOption,
		Index: -10,
	}
}

func indexOf(arr []int, val int) int {
	for i := range arr {
		if arr[i] == val {
			return i
		}
	}

	return -1
}

func getMyRoundPoints(op Option, opponent Option) int {
	if op.Name == DefaultOption {
		return -1
	}
	var outcomes [][]int = [][]int{
		{3, 0, 6},
		{6, 3, 0},
		{0, 6, 3},
	}

	optionValue := 0
	switch op.RequiredResult {
	case 0:
		optionValue = indexOf(outcomes[opponent.Index], 6) + 1
		break

	case 1:
		optionValue = opponent.Index + 1
		op.RequiredResult = 3
		break

	case 2:
		optionValue = indexOf(outcomes[opponent.Index], 0) + 1
		op.RequiredResult = 6
		break

	default:
		panic(errors.New("player one's input is not correct"))
	}
	return optionValue + op.RequiredResult
}

func LevelTwo(scanner *bufio.Scanner) int {
	var p1 = Player{
		Name:   "Opponent",
		Inputs: []Option{},
		Points: 0,
	}
	var p2 = Player{
		Name:   "Me",
		Inputs: []Option{},
		Points: 0,
	}

	op1 := Option{
		Name:  DefaultOption,
		Index: -1,
	}
	op2 := Option{
		Name:  DefaultOption,
		Index: -1,
	}
	for scanner.Scan() {
		op1 = getOpponentOption(string(scanner.Text()[0]))
		op2 = getMyOption(string(scanner.Text()[2]))

		p1.Inputs = append(p1.Inputs, op1)
		p2.Inputs = append(p2.Inputs, op2)

		r1, _ := getRoundPoints(op1, op2)

		r2 := getMyRoundPoints(op2, op1)
		fmt.Println("Round points: ", r1, r2)
		p1.Points += r1
		p2.Points += r2

		fmt.Printf("%s(%s) - %s \n", string(scanner.Text()[0]), string(op1.Name), string(scanner.Text()[2]))
		fmt.Printf("Points: P1: %d P2: %d", p1.Points, p2.Points)
		fmt.Println("")
	}

	return p2.Points
}
