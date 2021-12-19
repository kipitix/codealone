package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {

	argCount := len(os.Args)

	if argCount != 2 {
		fmt.Println("Usage <data file>")
		return
	}

	fileName := os.Args[1]

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Println("Can not open file")
		fmt.Println(fileErr.Error())
	}

	// Set global var reader
	reader = bufio.NewReader(file)

	// Read first line
	bladeCount := readInt('\n')

	lenArr := make([]int, bladeCount)

	// Read second line
	for i := 0; i < bladeCount; i++ {
		var val int
		if i == bladeCount-1 {
			val = readInt('\n')
		} else {
			val = readInt(' ')
		}
		lenArr[i] = val
	}

	sort.Slice(lenArr, func(i, j int) bool {
		return lenArr[i] > lenArr[j]
	})

	fmt.Println(bladeCount)
	fmt.Println(lenArr)

	successness, deepness := checkSum(lenArr, targLen)

	fmt.Println(successness)
	fmt.Println(deepness)
}

// Return success and deep count
func checkSum(lensArr []int, sum int) (bool, int) {

	for i := 0; i < len(lensArr); i++ {
		diff := sum - lensArr[i]

		if diff > 0 { // Go deep
			successness, deepness := checkSum(lensArr[1:], diff)
			if successness {
				fmt.Println(lensArr[i])
				return true, deepness + 1
			}
		} else { // >= 0 ! Found
			fmt.Println(lensArr[i])
			return true, 0
		}
	}

	return false, 0
}
