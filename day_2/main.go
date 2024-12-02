package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Level struct {
	Ordinals []int
	IsSafe   bool
	Ope      Operation
}

type Operation string

const (
	Increase Operation = "inc"
	Decrease Operation = "dec"
	Equity   Operation = "eq"
	None     Operation = ""
)

func getOperation(a int, b int) Operation {
	if a > b {
		return Decrease
	}
	if a == b {
		return Equity
	}
	return Increase
}

func NewLevel(values []string) Level {
	intVals := make([]int, 0)
	for idx := range values {
		intVal := Must(strconv.ParseInt(values[idx], 0, 64))
		intVals = append(intVals, int(intVal))
	}

	isSafe := true
	var lastOpe Operation
	for idx := range values {
		next := idx + 1
		if next >= len(intVals) {
			break
		}
		diff := intVals[idx] - intVals[next]
		operation := getOperation(intVals[idx], intVals[next])
		if lastOpe == None {
			lastOpe = operation
		}

		if lastOpe != operation {
			isSafe = false
			break
		}

		if operation == Equity {
			isSafe = false
			break
		}

		if math.Abs(float64(diff)) > 3 {
			isSafe = false
			break
		}

	}
	return Level{
		Ordinals: intVals,
		IsSafe:   isSafe,
		Ope:      lastOpe,
	}
}

func main() {
	file := Must(os.Open("./input/part1.txt"))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	levels := make([]Level, 0)

	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, " ")
		levels = append(levels, NewLevel(values))
	}

	count := 0
	for idx := range levels {
		if levels[idx].IsSafe {
			count++
		}
	}
	fmt.Printf("Level safe: %v\n", count)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
