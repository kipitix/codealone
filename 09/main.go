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

func readInput() (dights []int) {
	// Read first line
	readInt('\n') // void

	dights = readInts('\n', ' ')

	return
}

func findLongestSequence(dights []int, offset int) (length int) {
	processed := make(map[int]bool)
	index := offset
	length = 0

	for !processed[index] && index != dights[index]-1 {
		processed[index] = true
		index = dights[index] - 1
		length++
	}

	return
}

func processVariants(dights []int) (length int) {
	for i, _ := range dights {
		seqLen := findLongestSequence(dights, i)
		if seqLen > length {
			length = seqLen
		}
	}
	return
}

func main() {

	if !processArgs() {
		return
	}

	result := processVariants(readInput())

	fmt.Println(result)
}
