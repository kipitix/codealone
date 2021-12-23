package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

const ROUTINES_COUNT = 16

var primeNumbers [][]int

func checkForPrime(number int) (prime bool) {
	var wg sync.WaitGroup
	results := make(chan bool, ROUTINES_COUNT)
	wg.Add(ROUTINES_COUNT)
	halfNumber := number / 2

	for _, nums := range primeNumbers {
		go func(searchNums []int) {
			defer wg.Done()
			prime := true
			for _, v := range searchNums {
				if number%v == 0 {
					prime = false
					break
				} else if v > halfNumber {
					break
				}
			}
			results <- prime
		}(nums)
	}

	wg.Wait()

	prime = true
	for i := 0; i < ROUTINES_COUNT; i++ {
		prime = prime && <-results
	}

	return
}

func prepareCheckTable(maxPrime int) {
	primeNumbers = make([][]int, ROUTINES_COUNT)
	for i := range primeNumbers {
		primeNumbers[i] = make([]int, 0)
	}

	primeNumbers[0] = append(primeNumbers[0], 2)
	number := 3
	indexToAdd := 1

	mlnCount := 0

	for number <= maxPrime {
		if checkForPrime(number) {
			primeNumbers[indexToAdd] = append(primeNumbers[indexToAdd], number)
			indexToAdd++
			if indexToAdd >= ROUTINES_COUNT {
				indexToAdd = 0
			}
		}

		number += 2

		newMlnCount := number / 100_000
		if newMlnCount != mlnCount {
			mlnCount = newMlnCount
			fmt.Println(mlnCount)
		}
	}
}

func saveToFile(filename string) (result bool) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		result = false
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	totalCount := 0
	for _, v := range primeNumbers {
		totalCount += len(v)
	}

	writer.WriteString(strconv.Itoa(totalCount) + "\n")

	for i := 0; i < totalCount; i++ {
		num := primeNumbers[i%ROUTINES_COUNT][i/ROUTINES_COUNT]
		writer.WriteString(strconv.Itoa(num) + " ")
	}
	writer.Flush()

	result = true
	return
}

func loadFromFile(filename string) (result bool) {
	primeNumbers = make([][]int, ROUTINES_COUNT)
	for i := range primeNumbers {
		primeNumbers[i] = make([]int, 0)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		result = false
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	countStr, _ := reader.ReadString('\n')
	totalCount, _ := strconv.Atoi(strings.TrimSuffix(countStr, "\n"))

	for i := 0; i < totalCount; i++ {
		numStr, _ := reader.ReadString(' ')
		num, _ := strconv.Atoi(strings.TrimSuffix(numStr, " "))

		primeNumbers[i%ROUTINES_COUNT] = append(primeNumbers[i%ROUTINES_COUNT], num)
	}

	result = true
	return
}

func findSolution() {
	number := 10_010_001

	bigPrimes := make([]int, 0)
	biggest := 0

	for {
		if checkForPrime(number) {
			//fmt.Println(number)
			for _, prime := range bigPrimes {
				if math.Abs(float64(prime-number)) <= 100 {
					check := prime * number
					if check > biggest && check < 100_000_000_000_000 {
						biggest = check
						fmt.Println(biggest)
						return
					}
				}
			}

			bigPrimes = append(bigPrimes, number)
		}

		number -= 2
	}
}

func main() {
	// prepareCheckTable(10_000_000)
	// saveToFile("primes.txt")

	loadFromFile("primes.txt")
	findSolution()
}
