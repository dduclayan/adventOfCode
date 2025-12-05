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
	position                float64
	zeroCounter             float64
	reachedZeroLastRotation bool
}

func (d *dial) nextPosition(instruction string) error {
	direction := string(instruction[0])

	clicks, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return err
	}

	log.Printf("Rotating %v. Starting at %v", instruction, d.position)
	switch strings.ToUpper(direction) {

	case "L":
		for i := 0; i < clicks; i++ {
			// Solves the issue of ending AND starting at zero. If we ended on 0 on the last rotation reset the dial
			// so we're not counting back into negative numbers.
			if d.position == 0 && d.reachedZeroLastRotation == true {
				d.position = 100
				d.position--
				d.reachedZeroLastRotation = false
				continue
			}
			d.position--
			if d.position == 0 {
				d.zeroCounter++
				d.reachedZeroLastRotation = true
			}
			log.Printf("went from %v to %v\n", d.position+1, d.position)
		}

	case "R":
		for i := 0; i < clicks; i++ {
			// Solves the issue of ending AND starting at zero. If we ended on 100 on the last rotation reset the dial
			// to prevent the dial from going above 100.
			if d.position == 100 && d.reachedZeroLastRotation == true {
				d.position = 0
				d.position++
				d.reachedZeroLastRotation = false
				continue
			}
			d.position++
			if d.position == 100 {
				d.zeroCounter++
				d.reachedZeroLastRotation = true
			}
			log.Printf("went from %v to %v\n", d.position-1, d.position)
		}
	}

	return nil
}

func timesAtZero(path string) (float64, error) {
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
	}

	return d.zeroCounter, nil
}

func main() {
	n, err := timesAtZero("2025/day1/part2/input.txt")
	if err != nil {
		fmt.Printf("an unexpected error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("password: %v", n)
}
