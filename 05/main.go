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

func readInput() ([][]int, int) {
	// Read first line
	dightCount := readInt(' ')
	targetIndex := readInt('\n')

	dights := make([][]int, dightCount)

	// Read dight variants line by line
	for i := 0; i < dightCount; i++ {
		line := readInts('\n', ' ')
		dights[i] = line
	}

	return dights, targetIndex
}

func composeNumber(dights [][]int, targetIndex int) int {
	result := 0

	variantCount := 1

	for _, dightVariants := range dights {
		variantCount *= len(dightVariants)
	}

	targetIndex--

	for _, dightVariants := range dights {

		variantCount /= len(dightVariants)

		variantIndex := targetIndex / variantCount

		result = result*10 + dightVariants[variantIndex]

		targetIndex %= variantCount
	}

	return result
}

func main() {

	if !processArgs() {
		return
	}

	result := composeNumber(readInput())

	fmt.Println(result)
}
