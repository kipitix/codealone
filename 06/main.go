package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Direction int

const (
	UNDEFINED Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

var reader = bufio.NewReader(os.Stdin)

func readEnum() (dir Direction, eof bool) {
	char, readErr := reader.ReadByte()

	if readErr == io.EOF {
		return UNDEFINED, true
	} else if readErr != nil {
		panic(readErr)
	}

	switch char {
	case 'L':
		return LEFT, false
	case 'R':
		return RIGHT, false
	case 'D':
		return DOWN, false
	case 'U':
		return UP, false
	default:
		panic("Unknown symbol")
	}
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

func moveForward(x, y int, dir Direction) (newX, newY int) {
	switch dir {
	case UP:
		newX, newY = x, y+1
	case DOWN:
		newX, newY = x, y-1
	case LEFT:
		newX, newY = x-1, y
	case RIGHT:
		newX, newY = x+1, y
	default:
		newX, newY = x, y
	}
	return
}

func travel() (x, y int) {
	x = 0
	y = 0

	for move, eof := readEnum(); !eof; move, eof = readEnum() {
		x, y = moveForward(x, y, move)
	}

	return
}

func composeResult(x, y int) (result string) {
	if y < 0 {
		result += strings.Repeat("D", -y)
	}
	if x < 0 {
		result += strings.Repeat("L", -x)
	}
	if x > 0 {
		result += strings.Repeat("R", x)
	}
	if y > 0 {
		result += strings.Repeat("U", y)
	}
	return
}

func main() {

	if !processArgs() {
		return
	}

	fmt.Println(composeResult(travel()))
}
