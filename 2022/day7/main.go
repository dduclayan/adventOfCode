package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"sort"
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
	fsSize               = 70000000
	updateSize           = 30000000
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
	findBigDirs(fs)
	fmt.Println(totalSize)
	answerTwo(fs)
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
		addSizeUpToRoot(n.Size, n)
	}
	return &root
}

func answerTwo(n *node) {
	// ppChildren(n)
	logger.Debugf("size of fs: %v", n.Size)
	currBytesFree := fsSize - n.Size
	logger.Debugf("current bytes free=%v\n", currBytesFree)
	if updateSize > currBytesFree {
		fmt.Printf("don't have enough space. Need %v bytes\n", updateSize-currBytesFree)
	}
	sizes := n.dirsizes()
	sort.Ints(sizes)
	i := sort.Search(len(sizes), func(i int) bool { return sizes[i] >= updateSize-currBytesFree })
	fmt.Println(sizes[i])
}

func (n *node) seek(target string) *node {
	for _, child := range n.Children {
		if child.Name == target {
			return child
		}
	}
	return nil
}

func (n *node) dirsizes() []int {
	res := []int{}
	if n.Isdir {
		res = append(res, n.Size)
		for _, child := range n.Children {
			res = append(res, child.dirsizes()...)
		}
	}
	return res
}

func addChild(parent *node, child *node) {
	parent.Children = append(parent.Children, child)
}

func addSizeUpToRoot(s int, n *node) *node {
	if n == nil {
		return nil
	}
	if n.Parent.Name == "/" && n.Isdir == true {
		n.Parent.Size += s
		return nil
	}
	if n.Parent.Name == "/" && n.Isdir == false {
		n.Parent.Size += s
		return nil
	}
	if n.Parent.Isdir == true {
		n.Parent.Size += s
		return addSizeUpToRoot(s, n.Parent)
	}
	return nil
}

func mkDir(name string, parentNode *node) *node {
	return &node{Name: name, Isdir: true, Parent: parentNode}
}

func mkFile(name string, size int, parentNode *node) *node {
	return &node{Name: name, Size: size, Isdir: false, Parent: parentNode}
}

func deepSeek(startingNode *node, target string) *node {
	if startingNode == nil {
		return nil
	}
	if startingNode.Name == target {
		return startingNode
	}
	for _, child := range startingNode.Children {
		if len(child.Children) > 1 {
			for _, c := range child.Children {
				if c.Name == target {
					return c
				}
			}
		}
		if child.Name == target {
			return child
		}
		return deepSeek(child, target)
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

// TODO(dduclayan): Find a better way to do this LOL.
func ppChildren(n *node) *node {
	if n == nil {
		return nil
	}
	if n.Name == "/" {
		fmt.Printf("%v %v %v\n", n.Name, n.Isdir, n.Size)
		for _, v := range n.Children {
			fmt.Printf("|-- %v %v %v\n", v.Name, v.Isdir, v.Size)
			for _, i := range v.Children {
				fmt.Printf("|\t|-- %v %v %v\n", i.Name, i.Isdir, i.Size)
				for _, j := range i.Children {
					fmt.Printf("|\t|\t|-- %v %v %v\n", j.Name, j.Isdir, j.Size)
					for _, k := range j.Children {
						fmt.Printf("|\t|\t|\t|-- %v %v %v\n", k.Name, k.Isdir, k.Size)
						for _, l := range k.Children {
							fmt.Printf("|\t|\t|\t|\t|-- %v %v %v\n", l.Name, l.Isdir, l.Size)
							for _, m := range l.Children {
								fmt.Printf("|\t|\t|\t|\t|\t|-- %v %v %v\n", m.Name, m.Isdir, m.Size)
								for _, p := range m.Children {
									fmt.Printf("|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", p.Name, p.Isdir, p.Size)
									for _, r := range p.Children {
										fmt.Printf("|\t|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", r.Name, r.Isdir, r.Size)
										for _, s := range r.Children {
											fmt.Printf("|\t|\t|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", s.Name, s.Isdir, s.Size)
											for _, t := range s.Children {
												fmt.Printf("|\t|\t|\t|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", t.Name, t.Isdir, t.Size)
												for _, u := range t.Children {
													fmt.Printf("|\t|\t|\t|\t|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", u.Name, u.Isdir, u.Size)
													for _, w := range u.Children {
														fmt.Println("****************") // if you see this it means there's more nested children
														fmt.Printf("|\t|\t|\t|\t|\t|\t|\t|\t|\t|\t|\t|-- %v %v %v\n", w.Name, w.Isdir, w.Size)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}
