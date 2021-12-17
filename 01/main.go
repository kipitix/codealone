package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argCount := len(os.Args)

	if argCount != 3 {
		fmt.Println("Usage <switch count 1> <switch count 2>")
		return
	}

	switchesCount1, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Println("Enter arg 1 as number")
	}

	switchesCount2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Enter arg 2 as number")
	}

	min := switchesCount1
	max := switchesCount2

	if min > max {
		min, max = max, min
	}

	result := min*2 + max

	fmt.Println(result)
}
