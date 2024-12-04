package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var grid = make([][]string, 0)

var (
	width  int
	height int
)

func main() {
	fmt.Println("==== day 4 ====")
	file := Must(os.Open("input/part2.txt"))
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, strings.Split(text, ""))
	}

	width = len(grid[0])
	height = len(grid[1])
	fmt.Printf("Size: %v:%v\n", width, height)

	counter := 0
	cross := 0
	for y, line := range grid {
		for x, letter := range line {
			if letter == "X" {
				if SearchRight(x, y) {
					counter++
				}
				if SearchBottom(x, y) {
					counter++
				}
				if SearchLeft(x, y) {
					counter++
				}
				if SearchTop(x, y) {
					counter++
				}

				if SearchDiagTopLeft(x, y) {
					counter++
				}
				if SearchDiagTopRight(x, y) {
					counter++
				}
				if SearchDiagBotRight(x, y) {
					counter++
				}
				if SearchDiagBotLeft(x, y) {
					counter++
				}
			}

			if letter == "A" {
				if !IsInBound(x, y) {
					continue
				}
				diag1 := grid[y-1][x-1] + letter + grid[y+1][x+1]
				diag2 := grid[y-1][x+1] + letter + grid[y+1][x-1]

				if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
					cross++
				}
			}
		}
	}

	fmt.Printf("XMAS=%v\n", counter)
	fmt.Printf("X-MAS=%v\n", cross)
}

func SearchRight(x int, y int) bool {
	str := ""

	for curr := x; curr < width && curr < x+4; curr++ {
		str += grid[y][curr]
	}
	return str == "XMAS"
}

func SearchBottom(x int, y int) bool {
	str := ""

	for curr := y; curr < height && curr < y+4; curr++ {
		str += grid[curr][x]
	}
	return str == "XMAS"
}

func SearchLeft(x int, y int) bool {
	str := ""

	for curr := x; curr >= 0 && curr > x-4; curr-- {
		str += grid[y][curr]
	}
	return str == "XMAS"
}

func SearchTop(x int, y int) bool {
	str := ""

	for curr := y; curr >= 0 && curr > y-4; curr-- {
		str += grid[curr][x]
	}
	return str == "XMAS"
}

func SearchDiagTopLeft(x int, y int) bool {
	str := ""
	for delta := 0; delta < 4; delta++ {
		if x-delta < 0 || y-delta < 0 {
			break
		}
		str += grid[y-delta][x-delta]
	}

	return str == "XMAS"
}

func IsInBound(x int, y int) bool {
	return x-1 >= 0 && y-1 >= 0 && x+1 < width && y+1 < height
}

func SearchDiagTopRight(x int, y int) bool {
	str := ""
	for delta := 0; delta < 4; delta++ {
		if x+delta > width-1 || y-delta < 0 {
			break
		}
		str += grid[y-delta][x+delta]
	}

	return str == "XMAS"
}

func SearchDiagBotRight(x int, y int) bool {
	str := ""
	for delta := 0; delta < 4; delta++ {
		if x+delta > width-1 || y+delta > height-1 {
			break
		}
		str += grid[y+delta][x+delta]
	}

	return str == "XMAS"
}

func SearchDiagBotLeft(x int, y int) bool {
	str := ""
	for delta := 0; delta < 4; delta++ {
		if x-delta < 0 || y+delta > height-1 {
			break
		}
		str += grid[y+delta][x-delta]
	}

	return str == "XMAS"
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
