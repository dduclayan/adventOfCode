package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	officialInputFilePath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2021\\day2\\input.txt"
	testFilePath          = "2021/day2/test.txt"
)

func main() {
	fmt.Println(answer1(officialInputFilePath))
	fmt.Println(answer2(officialInputFilePath))
	fmt.Println(answer2(testFilePath))
}

func answer1(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var dPos, hPos = 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		amt, err2 := strconv.Atoi(splitLine[1])
		if err2 != nil {
			log.Fatalf("unable to convert string %v to int\n", splitLine[1])
		}

		switch direction {
		case "forward":
			hPos += amt
		case "down":
			dPos += amt
		case "up":
			dPos -= amt
		default:
			fmt.Printf("unknown command: %v. Skipping...\n", direction)
			continue
		}
	}
	return dPos * hPos
}

func answer2(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var dPos, hPos, aim = 0, 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		amt, err2 := strconv.Atoi(splitLine[1])
		if err2 != nil {
			log.Fatalf("unable to convert string %v to int\n", splitLine[1])
		}

		switch direction {
		case "forward":
			hPos += amt
			if aim == 0 {
				continue
			}
			dPos += aim * amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		default:
			fmt.Printf("unknown command: %v. Skipping...\n", direction)
			continue
		}
	}
	fmt.Printf("dpos:%v\thpos:%v\taim:%v\n", dPos, hPos, aim)
	return dPos * hPos
}
