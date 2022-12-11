package day5

// TASK: https://adventofcode.com/2022/day/5

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Instruction Struct
type Instruction struct {
	Amount, From, To int
}

func (inst Instruction) execute(inventory *[]*Stack) error {
	if inventory == nil {
		return errors.New("inventory is empty")
	}

	// START: INCLUDE ONLY FOR PART ONE
	for i := 0; i < inst.Amount; i++ {
		target := (*inventory)[inst.From-1].Pop()
		if target == -1 {
			continue
		}
		(*inventory)[inst.To-1].Push(target)
	}

	// START: INCLUDE ONLY FOR PART TWO
	// target := (*inventory)[inst.From-1].Entries[:inst.Amount]

	// // Delete this segment from the
	// // requested stack
	// (*inventory)[inst.From-1].Entries = (*inventory)[inst.From-1].Entries[inst.Amount:]
	// (*inventory)[inst.To-1].PushSubStack(target)

	return nil
}

// createInstruction returns an instruction based on string input and
// eventually returns an error if the format is incorrect
func createInstruction(instruction string) (Instruction, error) {
	if !strings.Contains(instruction, "move") {
		return Instruction{}, errors.New(fmt.Sprintf("instruction shoul have the %q, %q, %q", "move", "from", "to"))
	}

	inst := strings.Split(instruction, " ")
	if len(inst) != 6 {
		return Instruction{}, errors.New("each list should be an array of exactly 6 elements")
	}

	var props []int = make([]int, 0, 3)
	for i := range inst {
		if i%2 != 0 {
			num, err := strconv.Atoi(inst[i])
			if err != nil {
				return Instruction{}, err
			}
			props = append(props, num)
		}
	}

	if len(props) != 3 {
		return Instruction{}, errors.New("each props list should be an array of exactly 3 elements")
	}
	return Instruction{
		Amount: props[0],
		From:   props[1],
		To:     props[2],
	}, nil
}

type Stack struct {
	ID       int
	Entries  []rune
	Position int
}

func (s *Stack) Push(r rune) {
	t := []rune{r}
	t = append(t, s.Entries...)
	s.Entries = t
}

// Used for PART TWO
func (s *Stack) PushSubStack(r []rune) {
	t := []rune{}
	t = append(t, r...)
	t = append(t, s.Entries...)
	s.Entries = t
}

func (s *Stack) Pop() rune {
	if len(s.Entries) == 0 {
		return -1
	}
	r := s.Entries[0]
	s.Entries = s.Entries[1:]
	return r
}

func isPartFromStack(row string) bool {
	if strings.Contains(row, "[") {
		return true
	}
	return false
}

func parseRow(row string) []string {
	lenght := (len(row) + 1) / 4
	result := make([]string, lenght)
	for idx := range result {
		position := 4*idx + 1
		//Check if it is space
		if row[position] == 32 {
			result[idx] = "!"
			continue
		}
		result[idx] = string(row[position])
	}
	return result
}

func fillInventory(inventory *[]*Stack, data [][]string) error {
	if inventory == nil {
		return errors.New("inventory is empty")
	}
	for _, row := range data {
		for idx, cell := range row {
			if cell == "!" {
				continue
			}

			r, _ := utf8.DecodeRuneInString(cell)
			(*inventory)[idx].Entries = append((*inventory)[idx].Entries, r)
		}

	}
	return nil
}

func PartOne(scanner *bufio.Scanner) string {
	initialStructure := [][]string{}
	queue := []Instruction{}
	currentRow := ""

	for scanner.Scan() {
		currentRow = scanner.Text()

		if isPartFromStack(currentRow) {
			initialStructure = append(initialStructure, parseRow((currentRow)))
		} else {
			instruction, err := createInstruction(currentRow)
			if err != nil {
				continue
			}
			queue = append(queue, instruction)
		}
	}

	if len(initialStructure) == 0 {
		return ""
	}

	inventory := []*Stack{}
	for idx := range initialStructure[0] {
		inventory = append(inventory, &Stack{
			ID:       idx + 1,
			Entries:  []rune{},
			Position: idx,
		})
	}

	fillInventory(&inventory, initialStructure)
	for _, instruction := range queue {
		instruction.execute(&inventory)
	}

	s := ""
	for v := range inventory {
		s += string(*&inventory[v].Entries[0])
	}

	return s
}
