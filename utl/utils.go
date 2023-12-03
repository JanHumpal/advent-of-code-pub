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
