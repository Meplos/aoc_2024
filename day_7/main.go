package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value          int
	Addition       *Node // Addition
	Multiplication *Node // Multiplication
}

func NewNode(value int) Node {
	return Node{
		Value:          value,
		Addition:       nil,
		Multiplication: nil,
	}
}

func (n *Node) AppendAddition(add *Node) {
	n.Addition = add
}

func (n *Node) AppendMultiplication(mult *Node) {
	n.Multiplication = mult
}

func main() {
	fmt.Println("==== Day 7 ====")
	file := Must(os.Open("input/part1.txt"))
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		res, operandes := ParseEquation(text)
		var last []*Node = make([]*Node, 0)
		for i, op := range operandes {
			nodeToCreate := int(math.Pow(2, float64(i)))
			nodeCreated := make([]*Node, 0)
			for x := 0; x < nodeToCreate; x++ {
				node := NewNode(op)
				nodeCreated = append(nodeCreated, &node)
			}
			nodeToAdd := 0
			for x := 0; x < len(last); x++ {

				nodeCreated[nodeToAdd].Value += last[x].Value
				last[x].AppendAddition(nodeCreated[nodeToAdd])
				nodeToAdd++
				nodeCreated[nodeToAdd].Value *= last[x].Value
				last[x].AppendMultiplication(nodeCreated[nodeToAdd])
				nodeToAdd++
			}
			last = make([]*Node, len(nodeCreated))
			copy(last, nodeCreated)
		}

		isValid := false
		for _, node := range last {
			if node.Value == res {
				isValid = true
				break
			}
		}

		if isValid {
			fmt.Println(text)
			count += res
		}

	}
	fmt.Printf("Result : %v\n", count)
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func ParseEquation(equation string) (int, []int) {
	eqSplit := strings.Split(equation, ": ")
	result := int(Must(strconv.ParseInt(eqSplit[0], 0, 64)))
	operandes := make([]int, 0)

	for _, op := range strings.Split(eqSplit[1], " ") {
		operandes = append(operandes, int(Must(strconv.ParseInt(op, 0, 64))))
	}

	return result, operandes
}
