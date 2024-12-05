package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(" ==== Day 5 ====")
	file := Must(os.Open("input/sample.txt"))
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
}

func ParseUpdateList(str string) {
	list := strings.Split(str, ",")
	fmt.Println(list)
}

func ParseRules(str string) {
	list := strings.Split(str, "|")
	fmt.Println(list)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
