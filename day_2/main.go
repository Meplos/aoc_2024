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

func stringToInts(values []string) []int {
	intVals := make([]int, 0)
	for idx := range values {
		intVal := Must(strconv.ParseInt(values[idx], 0, 64))
		intVals = append(intVals, int(intVal))
	}
	return intVals
}

func NewLevelP2(values []string) Level {
	ordinals := stringToInts(values)

	isSafe := isLevelSave(ordinals[0], ordinals[1:], None, false)

	return Level{
		Ordinals: ordinals,
		IsSafe:   isSafe,
		Ope:      None,
	}
}

func isLevelSave(origin int, values []int, lastOpe Operation, sbd bool) bool {
	if len(values) <= 0 {
		return true
	}
	next := values[0]

	diff := origin - next
	operation := getOperation(origin, next)
	if lastOpe == None {
		lastOpe = operation
	}

	if lastOpe != operation || operation == Equity || math.Abs(float64(diff)) > 3 {
		if !sbd {
			return isLevelSave(values[0], values[1:], lastOpe, true)
		}
		return false
	}
	return isLevelSave(next, values[1:], lastOpe, sbd)
}

func NewLevel(values []string) Level {
	intVals := stringToInts(values)

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
	file := Must(os.Open("./input/part2.txt"))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	levels := make([]Level, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		values := strings.Split(text, " ")
		levels = append(levels, NewLevelP2(values))
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
