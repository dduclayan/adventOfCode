package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	input       = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	actualInput = "9191906840-9191941337,7671-13230,2669677096-2669816099,2-12,229599-392092,48403409-48523311,96763-229430,1919163519-1919240770,74928-96389,638049-668065,34781-73835,736781-819688,831765539-831907263,5615884-5749554,14101091-14196519,7134383-7169141,413340-625418,849755289-849920418,7745350-7815119,16717-26267,4396832-4549887,87161544-87241541,4747436629-4747494891,335-549,867623-929630,53-77,1414-3089,940604-1043283,3444659-3500714,3629-7368,79-129,5488908-5597446,97922755-98097602,182-281,8336644992-8336729448,24-47,613-1077"
)

/*
1. Check if slice of nums % 2 == 0. that is if slice len is odd we can totally ignore it.
2. Split slice in half.
3. convert to int
4. compare both ints
*/

func findInvalidIDs(tok string) (int, error) {
	fmt.Printf("finding invalid IDs in range: %s\n", tok)
	var start, end, sum int
	var err error
	start, err = strconv.Atoi(strings.Split(tok, "-")[0])
	if err != nil {
		return 0, err
	}
	end, err = strconv.Atoi(strings.Split(tok, "-")[1])
	if err != nil {
		return 0, err
	}

	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if err != nil {
			return 0, err
		}

		// case 1: first half == second half
		if s[:len(s)/2] == s[len(s)/2:] {
			sum += i
		}
	}

	return sum, nil
}

func sumInvalidIDs(in string) (sum int64, err error) {
	tokens := strings.Split(in, ",")
	for _, tok := range tokens {
		ids, err := findInvalidIDs(tok)
		if err != nil {
			return 0, err
		}
		sum += int64(ids)
	}
	return sum, nil
}

func main() {
	answer, err := sumInvalidIDs(actualInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
