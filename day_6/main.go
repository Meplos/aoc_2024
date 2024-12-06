package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Direction string

const (
	Forward  Direction = "^"
	Backward Direction = "v"
	Right    Direction = ">"
	Left     Direction = "<"
	Obstacle string    = "#"
)

type Grid struct {
	Height int
	Width  int
	Board  [][]string
}

func NewGrid(grid [][]string) Grid {
	return Grid{
		Height: len(grid),
		Width:  len(grid[0]),
		Board:  grid,
	}
}

func (g *Guard) HasVisited(x int, y int) bool {
	h := Hash(x, y)
	_, ok := g.Visited[string(h)]
	return ok
}

func (g *Grid) Display(guard Guard) {
	ClearScreen()

	fmt.Printf("Screen[%v:%v]\n", g.Width, g.Height)
	for y, row := range g.Board {
		for x, col := range row {
			if y == guard.Y && x == guard.X {
				fmt.Printf("%v ", guard.Char)
				continue
			} else {
				if guard.HasVisited(x, y) {
					fmt.Printf("%v ", "X")
					continue
				}
			}
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
	fmt.Printf("Guard[%v][%v:%v] | Next dir : %v\n", guard.Char, guard.X, guard.Y, guard.ChangeDirection())
	time.Sleep(time.Duration(100 * time.Millisecond))
}

func (g *Grid) HasObstacle(x int, y int) bool {
	if !g.IsInBound(x, y) {
		return false
	}
	return g.Board[y][x] == Obstacle
}

type (
	Guard struct {
		X       int
		Y       int
		Char    Direction
		Visited map[string]bool
	}
)

func (g *Guard) ChangeDirection() Direction {
	switch g.Char {
	case Forward:
		return Right
	case Backward:
		return Left
	case Right:
		return Backward
	case Left:
		return Forward
	}
	panic("Invalid direction")
}

func (g *Guard) MoveForward(grid Grid, x int, y int) bool {
	nextX := g.X
	nextY := g.Y - 1
	if grid.HasObstacle(nextX, nextY) {
		return false
	}
	g.Visit(nextX, nextY)
	return true
}

func (g *Guard) MoveRight(grid Grid, x int, y int) bool {
	nextX := g.X + 1
	nextY := g.Y
	if grid.HasObstacle(nextX, nextY) {
		return false
	}
	g.Visit(nextX, nextY)
	return true
}

func (g *Guard) MoveBackward(grid Grid, x int, y int) bool {
	nextX := g.X
	nextY := g.Y + 1
	if grid.HasObstacle(nextX, nextY) {
		return false
	}
	g.Visit(nextX, nextY)
	return true
}

func (g *Guard) MoveLeft(grid Grid, x int, y int) bool {
	nextX := g.X - 1
	nextY := g.Y
	if grid.HasObstacle(nextX, nextY) {
		return false
	}
	g.Visit(nextX, nextY)
	return true
}

func (g *Guard) Move(grid Grid) bool {
	switch g.Char {
	case Forward:
		return g.MoveForward(grid, g.X, g.Y)
	case Backward:
		return g.MoveBackward(grid, g.X, g.Y)
	case Right:
		return g.MoveRight(grid, g.X, g.Y)
	case Left:
		return g.MoveLeft(grid, g.X, g.Y)
	}
	panic("Invalid direction")
}

func Hash(x int, y int) []byte {
	key := fmt.Sprintf("%v|%v", x, y)
	hash := md5.New()
	hash.Write([]byte(key))
	return hash.Sum(nil)
}

func NewGuard(x int, y int, direction Direction) Guard {
	h := Hash(x, y)
	return Guard{
		X:       x,
		Y:       y,
		Char:    direction,
		Visited: map[string]bool{string(h): true},
	}
}

func (g *Guard) Visit(x int, y int) {
	h := Hash(x, y)
	g.Visited[string(h)] = true
	g.X = x
	g.Y = y
}

func (g *Grid) IsInBound(x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func main() {
	file := Must(os.Open("input/part1.txt"))
	scanner := bufio.NewScanner(file)
	fmt.Println(" ==== Day 6 ====")

	scanner.Split(bufio.ScanLines)

	arr := make([][]string, 0)
	var guard Guard
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := make([]string, 0)
		for x, v := range strings.Split(text, "") {
			if v == string(Forward) || v == string(Backward) || v == string(Right) || v == string(Left) {
				guard = NewGuard(x, y, Direction(v))
				line = append(line, "X")
				continue
			}
			line = append(line, v)
		}
		arr = append(arr, line)

		y++

	}
	grid := NewGrid(arr)
	for grid.IsInBound(guard.X, guard.Y) {
		for !guard.Move(grid) {
			guard.Char = guard.ChangeDirection()
		}

		// grid.Display(guard)
	}
	fmt.Printf("Result = %v\n", len(guard.Visited)-1)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func ClearScreen() {
	out, _ := exec.Command("/usr/bin/clear").Output()
	os.Stdout.Write(out)
}
