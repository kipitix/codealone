package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argCount := len(os.Args)

	if argCount != 3 {
		fmt.Println("Usage <window length> <meet coordinate>")
		return
	}

	wndLen, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Println("Enter arg 1 as number")
	}

	meetCoord, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Enter arg 2 as number")
	}

	car1PosTarg := 0
	car2PosTarg := wndLen

	car1Pos := car1PosTarg
	car2Pos := car2PosTarg

	car1Move := 1
	car2Move := -1

	iterCount := 0

	for car1Pos != car1PosTarg || car2Pos != car2PosTarg || iterCount <= 0 {

		car1Pos += car1Move
		car2Pos += car2Move

		if car1Pos == meetCoord || car1Pos == car1PosTarg {
			car1Move *= -1
		}

		if car2Pos == meetCoord || car2Pos == car2PosTarg {
			car2Move *= -1
		}

		iterCount++
	}

	fmt.Println(iterCount)
}
