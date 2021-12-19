package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func readInput() (indexSum int) {
	// Read first line
	count := readInt('\n')

	// Read dight variants line by line
	for i := 0; i < count; i++ {
		line := readInts('\n', ' ')
		if line[0] == 0 {
			indexSum += i + 1
		}
	}

	return
}

func main() {

	if !processArgs() {
		return
	}

	result := readInput()

	fmt.Println(result)
}
