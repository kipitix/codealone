package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Result struct {
	coord Coord
	dist  int
}

var reader = bufio.NewReader(os.Stdin)

func readInt(sep byte) int {
	str, _ := reader.ReadString(sep)
	str = strings.TrimSuffix(str, string(sep))
	result, _ := strconv.Atoi(str)
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

func readInput() (coords []Coord) {

	coordCount := readInt('\n')

	coords = make([]Coord, coordCount)

	for i := 0; i < coordCount; i++ {
		coords[i].x = readInt(' ')
		coords[i].y = readInt('\n')
	}

	return
}

func calcDist(one, two Coord) (dist int) {

	x1, x2 := one.x, two.x
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	y1, y2 := one.y, two.y
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	dist = (x2 - x1) + (y2 - y1)

	return
}

func findSolution(input []Coord) (result Coord) {

	minX := math.MaxInt64
	minY := math.MaxInt64
	maxX := 0
	maxY := 0
	stepX := 1
	stepY := 1

	for _, coord := range input {
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	distSumMin := math.MaxInt64

	// minY = 49500
	// maxY = 50000
	minX = 45000
	minY = 46000

	fmt.Println("x min max", minX, maxX)
	fmt.Println("y min max", minY, maxY)

	chanRes := make(chan Result)
	chanLimit := make(chan bool, 16)
	//statTicker := time.NewTicker(10 * time.Second)

	// Send requests
	go func() {
		for x := minX; x <= maxX; x += stepX {
			for y := minY; y <= maxY; y += stepY {

				chanLimit <- true

				coord := Coord{x, y}

				go func(checkCoord Coord, limit int) {

					distSum := 0

					for _, anotherCoord := range input {
						distSum += calcDist(checkCoord, anotherCoord)
						if distSum > limit {
							break
						}
					}

					chanRes <- Result{checkCoord, distSum}
				}(coord, distSumMin)

			}
		}
	}()

	// Handle results
	count := (maxX - minX) * (maxY - minY) / stepX / stepY
	for i := 0; i < count; i++ {
		select {
		case calcRes := <-chanRes:

			// Check with previous
			if calcRes.dist < distSumMin {
				distSumMin = calcRes.dist
				result = calcRes.coord

				fmt.Println(calcRes, result, distSumMin)

			} else if calcRes.dist == distSumMin {
				if calcRes.coord.x+calcRes.coord.y < result.x+result.y {
					result = calcRes.coord

					fmt.Println(calcRes, result, distSumMin)
				}
			}

			_ = <-chanLimit

		}
	}

	return
}

func main() {

	if !processArgs() {
		return
	}

	input := readInput()

	fmt.Println(findSolution(input))
}
