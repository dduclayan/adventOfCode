package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	testDataPathTemplate = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\%s\\input.txt"
	smallDataTemplate    = "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\%s\\input2.txt"
	logger               *zap.SugaredLogger
	maxSize              = 100000
	totalSize            = 0
)

type node struct {
	Isdir    bool
	Name     string
	Size     int
	Parent   *node
	Children []*node
}

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
	isdir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("couldn't get working Isdir: %v", err)
	}
	splitDir := strings.Split(isdir, "\\")
	day := splitDir[6]
	actualInputPath := fmt.Sprintf(testDataPathTemplate, day)
	// smallInputPath := fmt.Sprintf(actualInputPath, day)
	fs := createFs(actualInputPath)
	ppChildren(fs)
	findBigDirs(fs)
	fmt.Println(totalSize)
	// answerTwo(smallInputPath)
	// fmt.Printf("part one answer: %v\n", a)
	// fmt.Printf("part two answer: %v\n", b)
	fmt.Printf("finished executing in %v\n", time.Since(start))
}

func createFs(filePath string) *node {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	root := node{
		Isdir:    true,
		Name:     "/",
		Size:     0,
		Parent:   nil,
		Children: []*node{},
	}
	var currentDir *node
	currentDir = &root
	lineCounter := 0
	for scanner.Scan() {
		lineCounter++
		line := strings.Split(scanner.Text(), " ")
		if scanner.Text() == "$ ls" {
			continue
		}
		if string(line[0]) == "dir" {
			dirName := line[1]
			n := mkDir(dirName, currentDir)
			addChild(currentDir, n)
			continue
		}
		if strings.Join(line[:2], " ") == "$ cd" {
			if scanner.Text() == "$ cd /" { // skip first line
				continue
			}
			if scanner.Text() == "$ cd .." {
				currentDir = currentDir.Parent
				continue
			}
			currentDir = currentDir.seek(line[2])
			if currentDir == nil {
				logger.Panicf("line=%v. exiting. couldn't find target=%v", lineCounter, line[2])
			}
			continue
		}
		fileSz, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatalf("strconv.Atoi(%v) failed: %v\nv", line[0], err)
		}
		n := mkFile(line[1], fileSz, currentDir)
		addChild(currentDir, n)
		currentDir.Size += n.Size
		if currentDir.Name == "/" {
			continue
		}
		currentDir.Parent.Size += n.Size
		if currentDir.Parent.Name == "/" {
			continue
		}
		currentDir.Parent.Parent.Size += n.Size
	}
	return &root
}

// func answerTwo(filePath string) {
// }

func addChild(parent *node, child *node) {
	parent.Children = append(parent.Children, child)
}

func mkDir(name string, parentNode *node) *node {
	return &node{Name: name, Isdir: true, Parent: parentNode}
}

func mkFile(name string, size int, parentNode *node) *node {
	return &node{Name: name, Size: size, Isdir: false, Parent: parentNode}
}

func (n *node) seek(target string) *node {
	for _, child := range n.Children {
		if child.Name == target {
			// logger.Debugf("found target node: %v", target)
			return child
		}
	}
	return nil
}

func findBigDirs(n *node) *node {
	if n == nil {
		return nil
	}
	for _, child := range n.Children {
		if child.Size <= maxSize && child.Isdir == true {
			totalSize += child.Size
		}
		findBigDirs(child)
	}
	return nil
}

// TODO(dduclayan): Do this recursively instead.
func ppChildren(n *node) *node {
	if n == nil {
		return nil
	}
	if n.Name == "/" {
		fmt.Println("/")
		for _, v := range n.Children {
			fmt.Printf("|-- %v %v %v\n", v.Name, v.Isdir, v.Size)
			for _, i := range v.Children {
				fmt.Printf("|\t|-- %v %v %v\n", i.Name, i.Isdir, i.Size)
				for _, j := range i.Children {
					fmt.Printf("|\t|\t|-- %v %v %v\n", j.Name, j.Isdir, j.Size)
				}
			}
		}
	}
	return nil
}
