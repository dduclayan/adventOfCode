package main

import (
	"go.uber.org/zap"
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
	// start := time.Now()
	// dir, err := os.Getwd()
	// if err != nil {
	//	logger.Fatalf("couldn't get working dir: %v", err)
	// }
	// splitDir := strings.Split(dir, "\\")
	// day := splitDir[6]
	// actualInputPath := fmt.Sprintf(testDataPathTemplate, day)
	// smallInputPath := fmt.Sprintf(smallDataTemplate, day)
	// answerOne(smallInputPath)
	// answerTwo(smallInputPath)
	// fmt.Printf("part one answer: %v\n", a)
	// fmt.Printf("part two answer: %v\n", b)
	// fmt.Printf("finished executing in %v\n", time.Since(start))
}

// func answerOne(filePath string) {
// }

// func answerTwo(filePath string) {
// }
