package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	actualInput = "9191906840-9191941337,7671-13230,2669677096-2669816099,2-12,229599-392092,48403409-48523311,96763-229430,1919163519-1919240770,74928-96389,638049-668065,34781-73835,736781-819688,831765539-831907263,5615884-5749554,14101091-14196519,7134383-7169141,413340-625418,849755289-849920418,7745350-7815119,16717-26267,4396832-4549887,87161544-87241541,4747436629-4747494891,335-549,867623-929630,53-77,1414-3089,940604-1043283,3444659-3500714,3629-7368,79-129,5488908-5597446,97922755-98097602,182-281,8336644992-8336729448,24-47,613-1077"
)

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
		if i <= 10 {
			continue
		}

		log.Printf("i = %v", i)
		s := strconv.Itoa(i)
		if err != nil {
			return 0, err
		}

		// case 1: first half == second half
		if s[:len(s)/2] == s[len(s)/2:] {
			log.Printf("found match where first half equal second half")
			sum += i
			continue
		}

		// case 2: num is multiple of 3, at least two groups of three consecutively match
		// E.g., 824824824
		if len(s)%3 == 0 {
			var substr []string
			for i := 0; i < len(s); i += 3 {
				substr = append(substr, s[i:i+3])
			}
			log.Printf("%v", substr)

			matchedCount := 0
			initialChunk := substr[0]
			for idx := range substr {
				if initialChunk == substr[idx] {
					matchedCount++
				}
			}
			if matchedCount == len(substr) && len(substr) >= 2 {
				log.Printf("found match for case 2")
				sum += i
				continue
			}
		}

		// case 3: num is multiple of 2
		if len(s)%2 == 0 {

			// handle nums like 11 and 99
			if i < 100 {
				num := strings.Split(s, "")
				if num[0] == num[1] {
					sum += i
					log.Printf("found match with %v. num[0] = %v, num[1] = %v", i, num[0], num[1])
					continue
				}
				continue
			}

			var substr []string
			for i := 0; i < len(s); i += 2 {
				substr = append(substr, s[i:i+2])
			}
			log.Printf("%v", substr)

			matchedCount := 0
			initialChunk := substr[0]
			for idx := range substr {
				if initialChunk == substr[idx] {
					matchedCount++
				}
			}
			if matchedCount == len(substr) {
				log.Printf("found match for case 3")
				sum += i
				continue
			}
		}

		// case 4: odd value place nums like 11111
		if len(s)%2 == 1 {
			substr := strings.Split(s, "")
			matchedCount := 0
			initialChunk := substr[0]
			for idx := range substr {
				if initialChunk == substr[idx] {
					matchedCount++
				}
			}
			if matchedCount == len(substr) && len(substr) > 1 {
				log.Printf("found match for case 4")
				sum += i
				continue
			}
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
