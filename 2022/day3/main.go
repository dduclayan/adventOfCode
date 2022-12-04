// create map[string]bool
// split line into two compartments
// loop over first compartment
// do a nested loop of the second compartment
// check if value of the current iterator of second compartment matches first compartment
// if it DOES, set value to true in the map
// if it DOESN't, set value to false

package main

import (
	"advent_of_code/2022/points"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	testDataPath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day3\\input.txt"
)

func main() {
	answer, err := answerOne(testDataPath)
	check(err)
	fmt.Println(answer)
}

func check(err error) {
	if err != nil {
		log.Fatalf("some error occured: %v\n", err)
	}
}

func answerOne(filePath string) (int, error) {
	var sum int
	priorityMap := points.NewLettersToPointsMap()
	fmt.Println(priorityMap)
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(f)
	lineCounter := 0
	for scanner.Scan() {
		lineCounter++
		leftHalf := scanner.Text()[:len(scanner.Text())/2]
		rightHalf := scanner.Text()[len(scanner.Text())/2:]
		for _, charA := range leftHalf {
			if strings.Contains(rightHalf, string(charA)) {
				fmt.Printf("match found! line: %v, left half: %v, right half: %v, matching char: %q, priority value: %v\n", lineCounter, leftHalf, rightHalf, string(charA), priorityMap[string(charA)])
				sum += priorityMap[string(charA)]
				break
			}
		}
	}
	return sum, nil
}
