package test_helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SetupTestDataString(exercise string, dataset string) []string {
	file, err := os.Open(fmt.Sprintf("../test-data/%s/%s", exercise, dataset))
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	file.Seek(0, 0)

	fileScanner = bufio.NewScanner(file)
	fileLines := make([]string, lineCount)
	count := 0
	for fileScanner.Scan() {
		fileLines[count] = fileScanner.Text()
		count++
	}
	return fileLines
}

func SetupTestData(exercise string, dataset string) []int {
	file, err := os.Open(fmt.Sprintf("../test-data/%s/%s", exercise, dataset))
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	file.Seek(0, 0)

	fileScanner = bufio.NewScanner(file)
	fileLines := make([]int, lineCount)
	count := 0

	for fileScanner.Scan() {
		line, err := strconv.Atoi(fileScanner.Text())
		check(err)
		fileLines[count] = line
		count++
	}

	return fileLines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
