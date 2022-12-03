//  https://www.reddit.com/r/adventofcode/comments/xk2ex5/2021_day_1_part_1/ - why I kept getting 1265
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	path = `D:\Documents\go_workspace\src\advent_of_code\2021\day1\input.txt`
)

func main() {
	data := readFile(path)
	convertedData := convertDataToInt(data)
	fmt.Println(answerOne(convertedData))
	fmt.Println(answerTwo(convertedData))
}

func answerOne(input []int) int {
	increasedCount := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			increasedCount++
		}
	}
	return increasedCount
}

func answerTwo(input []int) int {
	increasedCount := 0

	for i := 0; i < len(input)-1; i++ {
		if i == len(input)-4 {
			currentWindow := input[i] + input[i+1] + input[i+2]
			nextWindow := input[i+1] + input[i+2] + input[i+3]
			if currentWindow < nextWindow {
				increasedCount++
			}
			return increasedCount
		}
		currentWindow := input[i] + input[i+1] + input[i+2]
		nextWindow := input[i+1] + input[i+2] + input[i+3]
		if currentWindow < nextWindow {
			increasedCount++
		}
	}
	return increasedCount
}

func convertDataToInt(input []string) []int {
	var convertedData []int
	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		convertedData = append(convertedData, num)
	}
	return convertedData
}

func readFile(filePath string) []string {
	var contents []string
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents
}
