/*
plan:
expand each pair
make each pair a string
use string.Contains to see if either string fully contains the other

Edit:
Above plan sucks. Much simpler to mathematically compare.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	testDataPath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day4\\input.txt"
	smallTest    = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day4\\input2.txt"
)

func main() {
	ans, err := answerOneAlt(testDataPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}

func answerOneAlt(filePath string) (int, error) {
	var assignmentPairCount int
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		leftHalf := strings.Split(scanner.Text(), ",")[0]
		rightHalf := strings.Split(scanner.Text(), ",")[1]
		leftHalfFirstNum, err := strconv.Atoi(strings.Split(leftHalf, "-")[0])
		if err != nil {
			panic(err)
		}
		leftHalfSecNum, err := strconv.Atoi(strings.Split(leftHalf, "-")[1])
		if err != nil {
			panic(err)
		}
		rightHalfFirstNum, err := strconv.Atoi(strings.Split(rightHalf, "-")[0])
		if err != nil {
			panic(err)
		}
		rightHalfSecNum, err := strconv.Atoi(strings.Split(rightHalf, "-")[1])
		if err != nil {
			panic(err)
		}
		// check if first pair contains second pair
		if leftHalfFirstNum >= rightHalfFirstNum {
			if leftHalfSecNum <= rightHalfSecNum {
				assignmentPairCount++
				continue
			}
		}
		// check if second pair contains first pair
		if rightHalfFirstNum >= leftHalfFirstNum {
			if rightHalfSecNum <= leftHalfSecNum {
				assignmentPairCount++
			}
		}
	}
	return assignmentPairCount, nil
}
