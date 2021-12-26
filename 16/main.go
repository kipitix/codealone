package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type MapUnit int

const (
	WALL MapUnit = iota
	FREE
	START
	FINISH
)

type Direction int

const (
	UNDEF Direction = iota
	RIGHT
	UP
	LEFT
	DOWN
)

type DirUnit struct {
	moveHere Direction
	moveOut  Direction
}

var reader = bufio.NewReader(os.Stdin)

func readLine(delimVal, delimLine byte) (eof bool, result []MapUnit) {
	// Read line
	str, error := reader.ReadString(delimLine)

	if error == nil {
		eof = false
	} else if error == io.EOF {
		eof = true
		return
	} else {
		panic(error)
	}

	str = strings.TrimSuffix(str, string(delimLine))
	str = strings.TrimSuffix(str, string(delimVal))

	values := strings.Split(str, string(delimVal))

	result = make([]MapUnit, len(values))

	for i, val := range values {
		switch val {
		case "A":
			result[i] = START
		case "B":
			result[i] = FINISH
		case "#":
			result[i] = WALL
		case ".":
			result[i] = FREE
		default:
			panic("Unhandled symbol")
		}
	}

	return
}

func readArgs() bool {
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

func readInput() (result [][]MapUnit) {

	result = make([][]MapUnit, 0)

	for eof, line := readLine(' ', '\n'); !eof; eof, line = readLine(' ', '\n') {
		result = append(result, line)
	}

	return
}

func preprocess(input [][]MapUnit) (good bool, width, height, startX, startY int, readyMap [][]MapUnit) {

	width = len(input[0])

	for y, line := range input {
		if width != len(line) {
			panic("Wrong line width")
		}

		for x, val := range line {
			if val == START {
				startX = x
				startY = y
				break
			}
		}
	}

	height = len(input)

	good = true

	readyMap = addSurroundWall(input)

	startX++
	startY++
	width += 2
	height += 2

	return
}

func addSurroundWall(input [][]MapUnit) (result [][]MapUnit) {
	width := len(input[0])

	upDownLine := make([]MapUnit, width+2)
	for i := 0; i < width; i++ {
		upDownLine[i] = WALL
	}

	result = make([][]MapUnit, 0)

	result = append(result, upDownLine)

	for _, line := range input {
		newLine := append([]MapUnit{WALL}, line...)
		newLine = append(newLine, WALL)
		result = append(result, newLine)
	}

	result = append(result, upDownLine)

	return
}

func findSolution(input [][]MapUnit, width, height, startX, startY int) (minStepCount int) {

	var dirMap [][]DirUnit
	dirMap = make([][]DirUnit, height)
	for i := range dirMap {
		dirMap[i] = make([]DirUnit, width)
	}

	minStepCount = math.MaxInt64
	stepCount := 0

	x, y := startX, startY
	search := true

	for search {

		foundFree, foundFinish, freeDirection := findFreeDir(input, dirMap, x, y)

		if foundFree {
			dirMap[y][x].moveOut = freeDirection

			x, y = goForward(x, y, freeDirection)
			stepCount++

			dirMap[y][x].moveHere = freeDirection
		} else {
			if foundFinish && stepCount+1 < minStepCount {
				minStepCount = stepCount + 1
				fmt.Println(minStepCount)
			}

			moveHere := dirMap[y][x].moveHere

			dirMap[y][x].moveOut = UNDEF
			dirMap[y][x].moveHere = UNDEF

			x, y = goBackward(x, y, moveHere)
			stepCount--
		}

		search = input[y][x] != START
	}

	return
}

func findFreeDir(input [][]MapUnit, dirMap [][]DirUnit, x, y int) (free, finish bool, dir Direction) {
	free = false
	finish = false
	dir = dirMap[y][x].moveOut

	search := true

	checkFreeFunc := func(x, y int) {
		if dirMap[y][x].moveOut == UNDEF {
			if input[y][x] == FREE {
				free = true
				search = false
			} else if input[y][x] == FINISH {
				finish = true
				search = false
			}
		}
	}

	for search {
		switch dir {
		case UNDEF:
			dir = RIGHT
			newX, newY := goForward(x, y, RIGHT)
			checkFreeFunc(newX, newY)
		case RIGHT:
			newX, newY := goForward(x, y, UP)
			dir = UP
			checkFreeFunc(newX, newY)
		case UP:
			newX, newY := goForward(x, y, LEFT)
			dir = LEFT
			checkFreeFunc(newX, newY)
		case LEFT:
			newX, newY := goForward(x, y, DOWN)
			dir = DOWN
			checkFreeFunc(newX, newY)
		case DOWN:
			dir = UNDEF
			search = false
		default:
			panic("Undefined condition")
		}
	}
	return
}

func goForward(x, y int, dir Direction) (newX, newY int) {
	switch dir {
	case UNDEF:
		newX, newY = x, y
	case UP:
		newX, newY = x, y-1
	case DOWN:
		newX, newY = x, y+1
	case LEFT:
		newX, newY = x-1, y
	case RIGHT:
		newX, newY = x+1, y
	default:
		panic("Bad direction")
	}
	return
}

func goBackward(x, y int, dir Direction) (newX, newY int) {
	switch dir {
	case UNDEF:
		newX, newY = x, y
	case UP:
		newX, newY = x, y+1
	case DOWN:
		newX, newY = x, y-1
	case LEFT:
		newX, newY = x+1, y
	case RIGHT:
		newX, newY = x-1, y
	default:
		panic("Bad direction")
	}
	return
}

func main() {

	if !readArgs() {
		return
	}

	tvMap := readInput()
	good, width, height, startX, startY, tvMap := preprocess(tvMap)

	if !good {
		return
	}

	fmt.Println(width, height, startX, startY)

	length := findSolution(tvMap, width, height, startX, startY)

	fmt.Println(length)

	// fmt.Println(index)
	// fmt.Println(dict)
}
