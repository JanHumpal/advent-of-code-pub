package utl

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput(fileName string) []string {
	inputFile, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	var textTs []string
	for scanner.Scan() {
		textTs = append(textTs, scanner.Text())
	}
	err = inputFile.Close()
	Check(err)
	return textTs
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func IntOf(s string) int {
	res, err := strconv.Atoi(s)
	Check(err)
	return res
}

func DigitCount(number int) int {
	return len(strconv.Itoa(number))
}

func GetSurrounding[T any](Ts []T, i int) []T {
	length := len(Ts)
	if i < 0 || i >= length {
		log.Fatalf("Cannot get surrounding, %v is out of bounds of passed array.", i)
	}
	if length == 1 {
		return []T{Ts[0]}
	}
	if i == 0 {
		return []T{Ts[i], Ts[i+1]}
	}
	if i == length-1 {
		return []T{Ts[i-1], Ts[i]}
	}
	return []T{Ts[i-1], Ts[i], Ts[i+1]}
}

func GetSurroundingOnly[T any](Ts []T, i int) []T {
	length := len(Ts)
	if i < 0 || i >= length {
		log.Fatalf("Cannot get surrounding, %v is out of bounds of passed array.", i)
	}
	if length == 1 {
		return []T{}
	}
	if i == 0 {
		return []T{Ts[i+1]}
	}
	if i == length-1 {
		return []T{Ts[i-1]}
	}
	return []T{Ts[i-1], Ts[i+1]}
}

// GCD
// Greatest Common Divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Last[T any](array []T) T {
	return array[LastI(array)]
}

func LastI[T any](array []T) int {
	return len(array) - 1
}

// RemoveAt Remove the element at index i from array; unstable
func RemoveAt[T any](i int, a *[]T) {
	(*a)[i] = Last(*a) // Copy last element to index i.
	var zero T
	(*a)[LastI(*a)] = zero // Erase last element (write zero value).
	*a = (*a)[:LastI(*a)]  // Truncate slice.
}

func RemoveAtStable[T any](i int, a *[]T) {
	copy((*a)[i:], (*a)[i+1:]) // Shift a[i+1:] left one index.
	var zero T
	(*a)[LastI(*a)] = zero // Erase last element (write zero value).
	*a = (*a)[:LastI(*a)]  // Truncate slice.
}
