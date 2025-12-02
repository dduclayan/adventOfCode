package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dial struct {
	position    float64
	zeroCounter int
}

func (d *dial) nextPosition(instruction string) error {
	direction := string(instruction[0])

	numRotationStr := instruction[1:]
	numRotation, err := strconv.Atoi(numRotationStr)
	if err != nil {
		return err
	}

	var newPos float64
	switch strings.ToUpper(direction) {
	// rotating the dial left, so we're subtracting
	case "L":
		newPos = d.position - float64(numRotation)
		if newPos < 0 {
			newPos = foldUp(newPos)
		}
	// rotating the dial right, so we're adding
	case "R":
		newPos = d.position + float64(numRotation)
		if newPos > 100 {
			newPos = foldDown(newPos)
		}
	}

	// dial goes from 0 -> 99. So if we land on 100 it's the same as getting to 0
	if newPos == 100 || newPos == 0 {
		d.position = 0
		d.zeroCounter++
		return nil
	}

	d.position = newPos

	return nil
}

func timesAtZero(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	d := &dial{
		position:    50,
		zeroCounter: 0,
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		err := d.nextPosition(line)
		if err != nil {
			return 0, err
		}
		log.Printf("position: %v", d.position)
	}

	return d.zeroCounter, nil
}

func main() {
	n, err := timesAtZero("day1/part1/input.txt")
	if err != nil {
		fmt.Printf("an unexpected error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("dial pointed at zero %d times.\n", n)
}

func foldDown(num float64) float64 {
	if num < 100 {
		return num
	}
	return foldDown(num - 100)
}

func foldUp(num float64) float64 {
	if num > 0 {
		return num
	}
	return foldUp(num + 100)
}
