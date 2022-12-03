package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	testPath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day1\\input.txt"
)

func main() {
	answerOne(testPath)
	answerTwo(testPath)
}

func answerOne(inputPath string) {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	elfToCalorieMap := make(map[int]int)
	elfCount := 1
	var calorieCount int
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			elfToCalorieMap[elfCount] = calorieCount
			elfCount++
			calorieCount = 0
			continue
		}
		if i, err := strconv.Atoi(currentLine); err == nil {
			calorieCount += i
		}
	}
	var elfToCalorieSlice [][]int
	for k, v := range elfToCalorieMap {
		elfToCalorieSlice = append(elfToCalorieSlice, []int{k, v})
	}
	sort.Slice(elfToCalorieSlice, func(i, j int) bool {
		return elfToCalorieSlice[i][1] > elfToCalorieSlice[j][1]
	})
	fmt.Println(elfToCalorieSlice[0][1])
}

func answerTwo(inputPath string) {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	elfToCalorieMap := make(map[int]int)
	elfCount := 1
	var calorieCount int
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			elfToCalorieMap[elfCount] = calorieCount
			elfCount++
			calorieCount = 0
			continue
		}
		if i, err := strconv.Atoi(currentLine); err == nil {
			calorieCount += i
		}
	}
	var elfToCalorieSlice [][]int
	for k, v := range elfToCalorieMap {
		elfToCalorieSlice = append(elfToCalorieSlice, []int{k, v})
	}
	sort.Slice(elfToCalorieSlice, func(i, j int) bool {
		return elfToCalorieSlice[i][1] > elfToCalorieSlice[j][1]
	})
	var sumTopThreeElfCalorieCount int
	for _, v := range elfToCalorieSlice[:3] {
		sumTopThreeElfCalorieCount += v[1]
	}
	fmt.Println(sumTopThreeElfCalorieCount)
}
