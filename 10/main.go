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

func readInput() (heights []int) {
	// Read first line
	readInt('\n') // void

	heights = readInts('\n', ' ')

	return
}

func findMaxArea(heights []int) (area int) {

	allCount := len(heights)

	for startInd, startVal := range heights {
		count := 1

		for i := startInd + 1; i < allCount; i++ {
			if heights[i] >= startVal {
				count++
			} else {
				break
			}
		}

		for i := startInd - 1; i >= 0; i-- {
			if heights[i] >= startVal {
				count++
			} else {
				break
			}
		}

		currArea := startVal * count
		if currArea > area {
			area = currArea
		}
	}
	return
}

func main() {

	if !processArgs() {
		return
	}

	result := findMaxArea(readInput())

	fmt.Println(result)
}
