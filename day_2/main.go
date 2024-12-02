package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := Must(os.Open("./input/sample.txt"))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
