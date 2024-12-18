package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type MulOperation struct {
	Left  int
	Right int
}

func (m *MulOperation) Execute() int {
	return m.Left * m.Right
}

func NewMul(l string, r string) MulOperation {
	return MulOperation{
		Left:  int(Must(strconv.ParseInt(l, 0, 64))),
		Right: int(Must(strconv.ParseInt(r, 0, 64))),
	}
}

type Condition string

const (
	Do   Condition = "do()"
	Dont Condition = "don't()"
)

func isCondition(str string) bool {
	return str == string(Do) || str == string(Dont)
}

func main() {
	fmt.Println("==== DAY 3 ====")

	file := Must(os.Open("input/part2.txt"))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instruction := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	isDoing := true
	integer := regexp.MustCompile(`\d+`)

	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		matched := instruction.FindAll([]byte(text), -1)
		for _, m := range matched {
			op := string(m)
			if isCondition(op) {
				if op == string(Dont) {
					isDoing = false
				}
				if op == string(Do) {
					isDoing = true
				}
				continue
			}

			if !isDoing {
				continue
			}

			integers := integer.FindAll(m, -1)
			operation := NewMul(string(integers[0]), string(integers[1]))
			total += operation.Execute()
		}
	}
	fmt.Printf("Result : %v\n", total)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
