package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	rules    = make(map[int][]int)
	goodList = make([][]int, 0)
)

func main() {
	fmt.Println(" ==== Day 5 ====")
	file := Must(os.Open("input/part1.txt"))
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	isUpdateList := false

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			isUpdateList = true
			continue
		}
		if isUpdateList {
			ParseUpdateList(text)
		} else {
			ParseRules(text)
		}
	}
	fmt.Println(goodList)
	counter := 0
	for _, currentList := range goodList {
		size := len(currentList)
		mid := size / 2
		counter += currentList[mid]
	}
	fmt.Printf("Result : %v\n", counter)
}

func ParseUpdateList(str string) {
	list := ConvertToIntList(strings.Split(str, ","))
	for curr, val := range list {
		for i := 0; i < curr; i++ {
			if slices.Contains(rules[val], list[i]) {
				return
			}
		}
	}

	goodList = append(goodList, list)
}

func ConvertToIntList(arr []string) []int {
	list := make([]int, 0)
	for _, val := range arr {
		list = append(list, int(Must(strconv.ParseInt(val, 0, 64))))
	}

	return list
}

func ParseRules(str string) {
	list := strings.Split(str, "|")
	x := int(Must(strconv.ParseInt(list[0], 0, 64)))
	y := int(Must(strconv.ParseInt(list[1], 0, 64)))
	if rules[x] == nil {
		rules[x] = make([]int, 0)
	}
	rules[x] = append(rules[x], y)
	fmt.Println(list)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
