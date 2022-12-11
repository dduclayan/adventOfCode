package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	testDataPath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day5\\input.txt"
	smallData    = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day5\\input2.txt"
)

func main() {
	a, err := answerOne(testDataPath)
	if err != nil {
		fmt.Println("couldn't retrieve answer")
		fmt.Println(err.Error())
	}
	fmt.Println(a)
}

func reverse(s []string) (reversedSlice []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return reversedSlice
}

func answerOne(filePath string) (string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	split := strings.Split(string(input), "\r\n\r\n")
	drawing := strings.Split(split[0], "\r\n")
	instructions := strings.Split(split[1], "\r\n")
	stacks := make([][]string, 9)

	for _, line := range drawing {
		columnCount := 0
		for j := 1; j < len(line); j += 4 { // no index stepping like in python, stuck with this
			columnCount++
			if line[j] == 32 { // check if string == ""
				continue
			}
			crate := string(line[j])
			stacks[columnCount-1] = append(stacks[columnCount-1], crate)
		}
	}

	for _, stack := range stacks {
		stack = reverse(stack)
	}

	for _, ins := range instructions {
		line := strings.Split(ins, " ")
		numOfBlocksToMove, err := strconv.Atoi(line[1])
		if err != nil {
			return "", err
		}
		src, err := strconv.Atoi(line[3])
		if err != nil {
			panic(err)
		}
		src--
		dest, err := strconv.Atoi(line[5])
		if err != nil {
			return "", err
		}
		dest--

		stacksToRemove := stacks[src][len(stacks[src])-numOfBlocksToMove:]
		fmt.Printf("ins: %v, moving stacks: %v from src: %v, to dest: %v\n", ins, stacksToRemove, stacks[src], stacks[dest])
		for i := len(stacksToRemove) - 1; i >= 0; i-- {
			stacks[dest] = append(stacks[dest], stacksToRemove[i])
		}

		stacks[src] = stacks[src][:len(stacks[src])-numOfBlocksToMove]
		fmt.Printf("src: %v, dest: %v\n", stacks[src], stacks[dest])
	}
	var ans []string
	for _, v := range stacks {
		ans = append(ans, v[len(v)-1])
	}
	return strings.Join(ans, ""), nil
}
