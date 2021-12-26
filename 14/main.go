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

func readChar(sep byte) (result string) {
	str, _ := reader.ReadString(sep)
	result = strings.TrimSuffix(str, string(sep))
	return
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

func readInput() (symbols []string) {
	// Read first line
	symbolCount := readInt(' ')
	restrictionsCount := readInt('\n')

	symbols = make([]string, symbolCount)
	for i := range symbols {
		symbols[i] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	for i := 0; i < restrictionsCount; i++ {
		symbolIndex := readInt(' ')
		symbolToRemove := readChar('\n')

		symbolIndex--

		symbols[symbolIndex] = strings.Replace(symbols[symbolIndex], symbolToRemove, "", 1)
	}

	return
}

func composeResult(input []string) (result string) {

	for _, v := range input {
		strLen := len(v)
		if strLen%2 == 0 {
			strLen--
		}
		medIndex := strLen / 2
		result += string(v[medIndex])
	}

	return
}

func main() {

	if !processArgs() {
		return
	}

	input := readInput()
	fmt.Println(composeResult(input))
}
