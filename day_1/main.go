package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Input struct {
	Left  []int
	Right []int
}

func NewInput() Input {
	return Input{
		Left:  make([]int, 0),
		Right: make([]int, 0),
	}
}

func (i *Input) AppendLeft(value int) {
	i.Left = append(i.Left, value)
}

func (i *Input) AppendRight(value int) {
	i.Right = append(i.Right, value)
}

func (i *Input) Sort() {
	slices.Sort(i.Left)
	slices.Sort(i.Right)
	fmt.Printf("Left: %v\n", i.Left)
	fmt.Printf("Right: %v\n", i.Right)
}

func (i *Input) Distance() int {
	d := 0
	for idx := range i.Left {
		diff := i.Left[idx] - i.Right[idx]
		d += int(math.Abs(float64(diff)))
	}
	return d
}

func main() {
	file := Must(os.Open("input/part1.txt"))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := NewInput()

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, "   ")
		left := Must(strconv.ParseInt(values[0], 0, 64))
		right := Must(strconv.ParseInt(values[1], 0, 64))

		input.AppendLeft(int(left))
		input.AppendRight(int(right))
	}

	input.Sort()
	fmt.Println(input.Distance())
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
