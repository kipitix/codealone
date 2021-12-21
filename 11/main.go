package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	number   int
	children []*Node
	parents  []*Node
	visited  bool
}

var reader = bufio.NewReader(os.Stdin)

func readInt(sep byte) int {
	str, _ := reader.ReadString(sep)
	str = strings.TrimSuffix(str, string(sep))
	result, _ := strconv.Atoi(str)
	return result
}

func readInts(lineSep byte, dightSep byte) []int {
	lineStr, _ := reader.ReadString(lineSep)
	lineStr = strings.TrimSuffix(lineStr, string(lineSep))
	lineStr = strings.TrimSuffix(lineStr, string(dightSep))
	dightsStr := strings.Split(lineStr, string(dightSep))

	result := make([]int, len(dightsStr))

	for i, dightStr := range dightsStr {
		dight, _ := strconv.Atoi(dightStr)
		result[i] = dight
	}

	return result
}

func processArgs() bool {
	argCount := len(os.Args)

	if argCount != 2 {
		fmt.Println("Usage <data file>")
		return false
	}

	fileName := os.Args[1]

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Println("Can not open file")
		fmt.Println(fileErr.Error())
		return false
	}

	// Set global var reader
	reader = bufio.NewReader(file)

	return true
}

func readInput() (screamers [][]int) {
	// Read first line
	count := readInt('\n')

	screamers = make([][]int, count)

	// Read dight variants line by line
	for i := 0; i < count; i++ {
		screamers[i] = readInts('\n', ' ')[1:]
	}

	return
}

func makeGraph(input [][]int) (nodes []*Node) {

	nodes = make([]*Node, len(input))

	for i, _ := range input {
		node := &Node{
			number:   i + 1,
			parents:  make([]*Node, 0),
			children: make([]*Node, 0),
			visited:  false,
		}
		nodes[i] = node
	}

	for index, binds := range input {
		for _, bind := range binds {
			bindIndex := bind - 1
			nodes[index].children = append(nodes[index].children, nodes[bindIndex])
			nodes[bindIndex].parents = append(nodes[bindIndex].parents, nodes[index])
		}
	}

	return
}

func handleGraph(nodes []*Node) (sequence []int) {

	work := true

	for work {

		allVisited := true

		for _, node := range nodes {
			if !node.visited {

				allVisited = false

				if checkChildrenVisited(node) {
					node.visited = true
					sequence = append(sequence, node.number)
					break
				}
			}
		}

		work = !allVisited
	}

	return
}

func checkChildrenVisited(node *Node) (allVisited bool) {
	allVisited = true
	// True if no children too
	for _, v := range node.children {
		if !v.visited {
			allVisited = false
			break
		}
	}
	return allVisited
}

func composeResult(sequence []int) (result int) {
	result = 0
	for i, v := range sequence {
		result += (i + 1) * v
	}
	return
}

func main() {

	if !processArgs() {
		return
	}

	input := readInput()

	graph := makeGraph(input)

	sequence := handleGraph(graph)

	fmt.Println(composeResult(sequence))
}
