/*
	part one prompt:
		1. identify the first position where the four most recently received characters were all different
		2. report the number of characters from the beginning of the buffer to the end of the first such four-character marker

	example:
		input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
		answer: 7
		explanation: 'jpqm' are the first non-repeating chars in the string

	part two prompt:
		1. 14 distinct characters rather than 4

	example:
		input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
		answer: 19
*/
package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

var (
	testDataPathTemplate = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\%s\\input.txt"
	smallDataTemplate    = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\%s\\input2.txt"
	logger               *zap.SugaredLogger
)

func init() {
	var err error
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger = l.Sugar()
}

func main() {
	start := time.Now()
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("couldn't get working dir: %v", err)
	}
	splitDir := strings.Split(dir, "\\")
	day := splitDir[6]
	actualInputPath := fmt.Sprintf(testDataPathTemplate, day)
	//smallInputPath := fmt.Sprintf(smallDataTemplate, day)

	a, err := answerOne(actualInputPath)
	if err != nil {
		fmt.Printf("couldn't retrieve answer: %v\n", err)
	}
	b, err := answerTwo(actualInputPath)
	if err != nil {
		fmt.Printf("couldn't retrieve answer: %v\n", err)
	}
	fmt.Printf("part one answer: %v\n", a)
	fmt.Printf("part two answer: %v\n", b)
	fmt.Printf("finished executing in %v\n", time.Since(start))
}

func answerOne(filePath string) (ans int, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("couldn't open file %v: %v", filePath, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "//") { // skip comments
			continue
		}
		chars := strings.Split(scanner.Text(), "")
		for i := 0; i < len(chars)-3; i++ {
			charMap := make(map[string]int)
			charGroup := chars[i : i+4]
			for _, c := range charGroup {
				charMap[c]++
			}
			if len(charMap) == 4 {
				logger.Infow("found first group with non-repeating chars",
					"characterMap", charMap,
					"index", i,
				)
				ans = len(chars[:i+4])
				break
			}
		}
	}
	return ans, nil
}

func answerTwo(filePath string) (ans int, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("couldn't open file %v: %v", filePath, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "//") { // skip comments
			continue
		}
		chars := strings.Split(scanner.Text(), "")
		for i := 0; i < len(chars)-13; i++ {
			charMap := make(map[string]int)
			charGroup := chars[i : i+14]
			for _, c := range charGroup {
				charMap[c]++
			}
			if len(charMap) == 14 {
				logger.Infow("found first group with non-repeating chars",
					"characterMap", charMap,
					"index", i,
				)
				ans = len(chars[:i+14])
				break
			}
		}
	}
	return ans, nil
}
