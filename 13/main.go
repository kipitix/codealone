package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type PhoneNumber struct {
	str string
	val int
}

func readInt(sep byte) int {
	str, _ := reader.ReadString(sep)
	str = strings.TrimSuffix(str, string(sep))
	result, _ := strconv.Atoi(str)
	return result
}

func readLine() (phoneNumber PhoneNumber) {
	// Read line
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, string('\n'))

	// Set straight string
	phoneNumber.str = str

	// Transform into number
	phoneNumber.val = convertToValue(str)

	return
}

func ignoreError(val bool, err error) bool {
	return val
}

func convertToValue(str string) (val int) {

	// if ignoreError(regexp.MatchString(`^\+7\d*\(`, str)) {
	// 	str = str[strings.Index(str, "("):]
	// } else if ignoreError(regexp.MatchString(`^\+7\d*\-`, str)) {
	// 	str = str[strings.Index(str, "-"):]
	// } else if ignoreError(regexp.MatchString(`^8\d*\(`, str)) {
	// 	str = str[strings.Index(str, "("):]
	// } else if ignoreError(regexp.MatchString(`^8\d*\-`, str)) {
	// 	str = str[strings.Index(str, "-"):]
	// } else

	if ignoreError(regexp.MatchString(`^\+7`, str)) {
		str = str[2:]
	} else if ignoreError(regexp.MatchString(`^8`, str)) {
		str = str[1:]
	} else {
		panic("No pattern")
	}

	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")

	number, err := strconv.Atoi(str)
	val = number

	if err != nil {
		panic(err.Error())
	}

	return
}

func readInput() (index int, phoneNumbers []PhoneNumber) {
	count := readInt(' ')
	index = readInt('\n')

	phoneNumbers = make([]PhoneNumber, count)

	for i := 0; i < count; i++ {
		phoneNumbers[i] = readLine()
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

func findPhone(index int, phoneNumbers []PhoneNumber) (result string) {

	sort.Slice(phoneNumbers, func(i, j int) bool {
		return phoneNumbers[i].val < phoneNumbers[j].val
	})

	result = phoneNumbers[index-1].str

	return
}

func main() {

	if !readArgs() {
		return
	}

	index, dict := readInput()

	fmt.Println(findPhone(index, dict))

	// fmt.Println(index)
	// fmt.Println(dict)
}
