/*
part one plan:
split line into two halves
loop over characters in first half
use strings.Contains to check if current iterator is found in right half
if found, add to sum
*/

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
				sum += priorityMap[string(charA)]
				break
			}
		}
	}
	return sum, nil
}
