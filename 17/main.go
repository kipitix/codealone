package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func readChar() (eof bool, char byte) {
	// Read line
	char, error := reader.ReadByte()

	if error == nil {
		eof = false
	} else if error == io.EOF {
		eof = true
		return
	} else {
		panic(error)
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

type Stack []byte

func (s *Stack) Count() int {
	return len(*s)
}

func (s *Stack) Top() byte {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(data byte) {
	*s = append(*s, data)
}

func (s *Stack) Pop() byte {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func process() (funcCount int) {
	var stack Stack
	var squareBracketCount int
	var roundBracketCount int
	var curveBracketCount int

	init := func() {
		stack = make(Stack, 0)
		squareBracketCount = 0
		roundBracketCount = 0
		curveBracketCount = 0
	}

	init()

	for eof, char := readChar(); !eof; eof, char = readChar() {
		switch char {
		case '[':
			squareBracketCount++
			stack.Push(char)
		case '(':
			roundBracketCount++
			stack.Push(char)
		case '{':
			curveBracketCount++
			stack.Push(char)
		case ']':
			if stack.Count() > 0 && stack.Top() == '[' {
				stack.Pop()
				squareBracketCount--
			} else {
				//panic("Break sequence")
				init()
			}
		case ')':
			if stack.Count() > 0 && stack.Top() == '(' {
				stack.Pop()
				roundBracketCount--
			} else {
				//panic("Break sequence")
				init()
			}
		case '}':
			if stack.Count() > 0 && stack.Top() == '{' {
				stack.Pop()
				curveBracketCount--

				if curveBracketCount == 0 {
					funcCount++
				}
			} else {
				//panic("Break sequence")
				init()
			}
		default:
			//panic("Unknown symbol")
			//fmt.Println(char)
			init()
		}
	}

	return
}

func main() {

	if !readArgs() {
		return
	}

	result := process()

	fmt.Println(result)
}
