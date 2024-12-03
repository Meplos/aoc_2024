package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("==== day 4 ====")
	file := Must(os.Open("input/sample.txt"))
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