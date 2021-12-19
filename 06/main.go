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
	FINISH
	VISIT
)

const FIELD_SIZE = 160

var reader = bufio.NewReader(os.Stdin)

var field [FIELD_SIZE][FIELD_SIZE]Direction

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

func fill() {
	// Start from center of field
	x := FIELD_SIZE / 2
	y := x

	for move, eof := readEnum(); !eof; move, eof = readEnum() {
		field[x][y] = move
		x, y = moveForward(x, y, move)
	}
	field[x][y] = FINISH
}

func travel() (up, down, left, right int) {
	// Start from center of field
	x := FIELD_SIZE / 2
	y := x

	// Counters
	move := field[x][y]

	for move != FINISH {

		switch move {
		case UP:
			up++
		case DOWN:
			down++
		case LEFT:
			left++
		case RIGHT:
			right++
		}

		field[x][y] = VISIT
		x, y = moveForward(x, y, move)
		move = field[x][y]

		fmt.Println(strings.Repeat("-", FIELD_SIZE))
		printField()
	}

	return
}

func composeResult(up, down, left, right int) (result string) {
	result += strings.Repeat("D", down)
	result += strings.Repeat("L", left)
	result += strings.Repeat("R", right)
	result += strings.Repeat("U", up)
	return
}

func printField() {
	for y := FIELD_SIZE - 1; y >= 0; y-- {
		for x := 0; x < FIELD_SIZE; x++ {
			switch field[x][y] {
			case UP:
				fmt.Print("U")
			case DOWN:
				fmt.Print("D")
			case LEFT:
				fmt.Print("L")
			case RIGHT:
				fmt.Print("R")
			case FINISH:
				fmt.Print("*")
			case VISIT:
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {

	if !processArgs() {
		return
	}

	fill()

	printField()

	fmt.Println(composeResult(travel()))
}
