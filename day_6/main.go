package main

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Direction string

const (
	Forward        Direction = "^"
	Backward       Direction = "v"
	Right          Direction = ">"
	Left           Direction = "<"
	Obstacle       string    = "#"
	CustomObstacle string    = "O"
	None           string    = "."
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

func (g *Grid) Copy() Grid {
	nBoard := make([][]string, 0, g.Height)
	for _, row := range g.Board {
		nRow := make([]string, 0, g.Width)
		for _, value := range row {
			nRow = append(nRow, value)
		}
		nBoard = append(nBoard, nRow)
	}
	return NewGrid(nBoard)
}

func (g *Grid) PlaceObstacle(x int, y int) {
	if !g.IsInBound(x, y) {
		panic("Obstacle not in bound")
	}
	g.Board[y][x] = CustomObstacle
}

func (g *Grid) ClearObstacle(x, y int) {
	if !g.IsInBound(x, y) {
		panic("Obstacle not in bound")
	}
	g.Board[y][x] = None
}

func (g *Grid) Display(guard Guard, nbIt int) {
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
	fmt.Printf("it n°%v\n", nbIt)
	time.Sleep(time.Duration(50 * time.Millisecond))
}

func (g *Grid) HasObstacle(x int, y int) bool {
	if !g.IsInBound(x, y) {
		return false
	}
	return g.Board[y][x] == Obstacle || g.Board[y][x] == CustomObstacle
}

func (g *Grid) HasCustomObstacle(x, y int) bool {
	if !g.IsInBound(x, y) {
		return false
	}
	return g.Board[y][x] == CustomObstacle
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

func (g *Guard) MoveForward(grid Grid, x int, y int) (bool, error) {
	nextX := g.X
	nextY := g.Y - 1
	if grid.HasObstacle(nextX, nextY) {
		return false, nil
	}
	if g.HasVisited(nextX, nextY) {
		return true, errors.New("Already pass")
	}
	g.Visit(nextX, nextY)
	return true, nil
}

func (g *Guard) MoveRight(grid Grid, x int, y int) (bool, error) {
	nextX := g.X + 1
	nextY := g.Y
	if grid.HasObstacle(nextX, nextY) {
		return false, nil
	}
	if g.HasVisited(nextX, nextY) {
		return true, errors.New("Already pass")
	}
	g.Visit(nextX, nextY)
	return true, nil
}

func (g *Guard) MoveBackward(grid Grid, x int, y int) (bool, error) {
	nextX := g.X
	nextY := g.Y + 1
	if grid.HasObstacle(nextX, nextY) {
		return false, nil
	}
	if g.HasVisited(nextX, nextY) {
		return true, errors.New("Already pass")
	}
	g.Visit(nextX, nextY)
	return true, nil
}

func (g *Guard) MoveLeft(grid Grid, x int, y int) (bool, error) {
	nextX := g.X - 1
	nextY := g.Y
	if grid.HasObstacle(nextX, nextY) {
		return false, nil
	}
	if g.HasVisited(nextX, nextY) {
		return true, errors.New("Already pass")
	}
	g.Visit(nextX, nextY)
	return true, nil
}

func (g *Guard) Move(grid Grid) (bool, error) {
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

func Hash(x int, y int, d Direction) []byte {
	key := fmt.Sprintf("%v|%v|%v", x, y, d)
	hash := md5.New()
	hash.Write([]byte(key))
	return hash.Sum(nil)
}

func NewGuard(x int, y int, direction Direction) Guard {
	h := Hash(x, y, direction)
	return Guard{
		X:       x,
		Y:       y,
		Char:    direction,
		Visited: map[string]bool{string(h): true},
	}
}

func (g *Guard) Copy() Guard {
	return Guard{
		X:       g.X,
		Y:       g.Y,
		Char:    g.Char,
		Visited: make(map[string]bool),
	}
}

func (g *Guard) Visit(x int, y int) {
	h := Hash(x, y, g.Char)
	g.Visited[string(h)] = true
	g.X = x
	g.Y = y
}

func (g *Guard) HasVisited(x int, y int) bool {
	h := Hash(x, y, g.Char)
	_, ok := g.Visited[string(h)]
	return ok
}

func (g *Grid) IsInBound(x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func main() {
	start := time.Now()
	file := Must(os.Open("input/part1.txt"))
	scanner := bufio.NewScanner(file)
	fmt.Println(" ==== Day 6 ====")

	scanner.Split(bufio.ScanLines)

	arr := make([][]string, 0)
	var originGuard Guard
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := make([]string, 0)
		for x, v := range strings.Split(text, "") {
			if v == string(Forward) || v == string(Backward) || v == string(Right) || v == string(Left) {
				originGuard = NewGuard(x, y, Direction(v))
				line = append(line, None)
				continue
			}
			line = append(line, v)
		}
		arr = append(arr, line)
		y++

	}
	res := 0
	origin := NewGrid(arr)
	yToClear := 0
	nbIt := 1
	for y := 0; y < origin.Height; y++ {
		for x := 0; x < origin.Width; x++ {
			//			fmt.Printf("it n°%v\n", nbIt)
			if origin.HasObstacle(x, y) || (x == originGuard.X && y == originGuard.Y) {
				continue
			}
			grid := origin.Copy()
			grid.PlaceObstacle(x, y)
			guard := originGuard.Copy()
			if RunPath(&guard, &grid, nbIt) {
				res++
			}
			nbIt++
		}
		yToClear++
	}
	end := time.Now()
	fmt.Printf("Result = %v\n", res)
	fmt.Printf("Duration : %v\n", end.Sub(start).Seconds())
}

func RunPath(guard *Guard, grid *Grid, nbIt int) bool {
	for grid.IsInBound(guard.X, guard.Y) {
		hasMove, err := guard.Move(*grid)
		if err != nil {
			return true
		}
		if !hasMove {
			guard.Char = guard.ChangeDirection()
		}

		// grid.Display(*guard, nbIt)
	}
	return false
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
