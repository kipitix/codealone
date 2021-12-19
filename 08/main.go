package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

var result string
var currentNumber, numberCount int

func readInt() (number int, eof bool) {
	char, readErr := reader.ReadByte()

	if readErr == io.EOF {
		eof = true
	} else if readErr == nil {
		num, _ := strconv.Atoi(string(char))
		number = num
	}

	return
}

func processArgs() (success bool) {
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

func generateWord(chars []byte, count int) (result string) {
	dightCount := len(chars)

	mod := count % dightCount
	div := count / dightCount

	if mod > 0 {
		result += string(chars[mod-1])
	}
	if div > 0 {
		result += strings.Repeat(string(chars[dightCount-1]), div)
	}

	return
}

func generateString(number, count int) (result string) {
	switch number {
	case 2:
		result = generateWord([]byte{'A', 'B', 'C'}, count)
	case 3:
		result = generateWord([]byte{'D', 'E', 'F'}, count)
	case 4:
		result = generateWord([]byte{'G', 'H', 'I'}, count)
	case 5:
		result = generateWord([]byte{'J', 'K', 'L'}, count)
	case 6:
		result = generateWord([]byte{'M', 'N', 'O'}, count)
	case 7:
		result = generateWord([]byte{'P', 'Q', 'R', 'S'}, count)
	case 8:
		result = generateWord([]byte{'T', 'U', 'V'}, count)
	case 9:
		result = generateWord([]byte{'W', 'X', 'Y', 'Z'}, count)
	default:
		panic("Unknown number")
	}
	return
}

func processNumber(number int) {
	if number != currentNumber {
		if currentNumber > 0 {
			result += generateString(currentNumber, numberCount)
		}
		currentNumber = number
		numberCount = 1
	} else {
		numberCount++
	}
}

func main() {

	if !processArgs() {
		return
	}

	for num, eof := readInt(); !eof; num, eof = readInt() {
		processNumber(num)
	}
	// To handle last sequence
	processNumber(0)

	fmt.Println(result)
}
