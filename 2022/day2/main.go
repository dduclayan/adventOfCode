// A = Rock
// B = Paper
// C = Scissors
//
// X = Rock
// Y = Paper
// Z = Scissors
//
// Rock = 1 pts
// Paper = 2 pts
// Scissors = 3 pts
//
// Lose = 0 pts
// Draw = 3 pts
// Win = 6 pts
//
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
//
// Part Two
// X = need to lose
// Y = need to draw
// Z = need to win

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	testDataPath = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day2\\input.txt"
)

func main() {
	score, err := answerOne(testDataPath)
	if err != nil {
		fmt.Printf("failed to retrieve score: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(score)

	score2, err := answerTwo(testDataPath)
	if err != nil {
		fmt.Printf("failed to retrieve score: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(score2)
}

func answerOne(inputPath string) (int, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var score int
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		opp := splitLine[0]
		me := splitLine[1]

		switch me {
		case "X":
			if opp == "A" { // rock v rock = draw
				score += 1 // point for choosing rock
				score += 3 // 3 pts for draw
			}
			if opp == "B" { // paper v rock = opp wins. no points granted in a loss
				score += 1 // point for choosing rock
			}
			if opp == "C" { // scissors v rock = I win.
				score += 1 // point for choosing rock
				score += 6 // 6 pts granted for a win
			}
		case "Y":
			if opp == "A" { // rock v paper = win
				score += 2 // point for choosing paper
				score += 6 // 6 pts for win
			}
			if opp == "B" { // paper v paper = draw. 3 pts.
				score += 2 // point for choosing paper
				score += 3 // 3 pts for draw
			}
			if opp == "C" { // scissors v paper = lose
				score += 2 // point for choosing paper
			}
		case "Z":
			if opp == "A" { // rock v scissors = lose
				score += 3 // point for choosing scissors
			}
			if opp == "B" { // paper v scissors = win. 6 pts.
				score += 3 // point for choosing scissors
				score += 6 // 6 pts for win
			}
			if opp == "C" { // scissors v scissors = draw
				score += 3 // point for choosing scissors
				score += 3 // 3 pts for draw
			}
		default:
			fmt.Printf("unknown case: %q. Please verify input is correct.\n", me)
		}
	}
	return score, nil
}

func answerTwo(inputPath string) (int, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var score int
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		opp := splitLine[0]
		me := splitLine[1]

		switch me {
		case "X": // need to lose
			if opp == "A" { // rock v scissors = lose
				score += 3 // point for choosing scissors
			}
			if opp == "B" { // paper v rock = lose
				score += 1 // point for choosing rock
			}
			if opp == "C" { // scissors v paper = lose
				score += 2 // point for choosing paper
			}
		case "Y": // need to draw
			if opp == "A" { // rock v rock = win
				score += 1 // point for choosing rock
				score += 3 // 3 pts for draw
			}
			if opp == "B" { // paper v paper = draw. 3 pts.
				score += 2 // point for choosing paper
				score += 3 // 3 pts for draw
			}
			if opp == "C" { // scissors v scissors = lose
				score += 3 // point for choosing scissors
				score += 3 // 3 pts for draw
			}
		case "Z": // need to win
			if opp == "A" { // rock v paper = win
				score += 2 // point for choosing paper
				score += 6 // 6 pts for win
			}
			if opp == "B" { // paper v scissors = win. 6 pts.
				score += 3 // point for choosing scissors
				score += 6 // 6 pts for win
			}
			if opp == "C" { // scissors v rock = draw
				score += 1 // point for choosing rock
				score += 6 // 6 pts for win
			}
		default:
			fmt.Printf("unknown case: %q. Please verify input is correct.\n", me)
		}
	}
	return score, nil
}
