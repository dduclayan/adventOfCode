package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

func init() {
	totalSize = 0
}

func TestCreateFs(t *testing.T) {
	bTxt := &node{
		Isdir:    false,
		Name:     "b.txt",
		Size:     14848514,
		Parent:   nil,
		Children: nil,
	}
	cDat := &node{
		Isdir:    false,
		Name:     "c.dat",
		Size:     8504156,
		Parent:   nil,
		Children: nil,
	}
	eDir := &node{
		Isdir:    true,
		Name:     "e",
		Size:     584,
		Parent:   nil,
		Children: nil,
	}
	iFile := &node{
		Isdir:    false,
		Name:     "i",
		Size:     584,
		Parent:   nil,
		Children: nil,
	}
	fFile := &node{
		Isdir:    false,
		Name:     "f",
		Size:     29116,
		Parent:   nil,
		Children: nil,
	}
	gFile := &node{
		Isdir:    false,
		Name:     "g",
		Size:     2557,
		Parent:   nil,
		Children: nil,
	}
	hLst := &node{
		Isdir:    false,
		Name:     "h.lst",
		Size:     62596,
		Parent:   nil,
		Children: nil,
	}
	jFile := &node{
		Isdir:    false,
		Name:     "j",
		Size:     4060174,
		Parent:   nil,
		Children: nil,
	}
	dLog := &node{
		Isdir:    false,
		Name:     "d.log",
		Size:     8033020,
		Parent:   nil,
		Children: nil,
	}
	dExt := &node{
		Isdir:    false,
		Name:     "d.ext",
		Size:     5626152,
		Parent:   nil,
		Children: nil,
	}
	kFile := &node{
		Isdir:    false,
		Name:     "k",
		Size:     7214296,
		Parent:   nil,
		Children: nil,
	}
	dDir := &node{
		Isdir:    true,
		Name:     "d",
		Size:     24933642,
		Parent:   nil,
		Children: nil,
	}
	aDir := &node{
		Isdir:    true,
		Name:     "a",
		Size:     94853,
		Parent:   nil,
		Children: nil,
	}
	root := &node{
		Isdir:    true,
		Name:     "/",
		Size:     48381165,
		Parent:   nil,
		Children: nil,
	}

	tests := []struct {
		testName  string
		inputPath string
		want      *node
	}{
		{
			testName:  "filesystem compare",
			inputPath: "D:\\Documents\\go_workspace\\src\\advent_of_code\\2022\\day7\\input2.txt",
			want:      generateRootNode(t, root, aDir, eDir, iFile, fFile, gFile, hLst, bTxt, cDat, dDir, jFile, dLog, dExt, kFile),
		},
	}

	for _, tc := range tests {
		got := createFs(tc.inputPath)
		sort.Slice(got.Children, func(i, j int) bool {
			return got.Children[i].Name < got.Children[j].Name
		})
		sort.Slice(tc.want.Children, func(i, j int) bool {
			return tc.want.Children[i].Name < tc.want.Children[j].Name
		})

		fmt.Println("======================")
		fmt.Println("printing got:")
		fmt.Println("======================")
		fmt.Println(ppChildren(got))
		fmt.Println("======================")
		fmt.Println("printing want:")
		fmt.Println("======================")
		fmt.Println(ppChildren(tc.want))

		if diff := cmp.Diff(tc.want, got); diff != "" {
			fmt.Printf("test %q failed\n", tc.testName)
			t.Errorf("createFs() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestFindBigDirs(t *testing.T) {
	bTxt := &node{
		Isdir:    false,
		Name:     "b.txt",
		Size:     14848514,
		Parent:   nil,
		Children: nil,
	}
	cDat := &node{
		Isdir:    false,
		Name:     "c.dat",
		Size:     8504156,
		Parent:   nil,
		Children: nil,
	}
	eDir := &node{
		Isdir:    true,
		Name:     "e",
		Size:     584,
		Parent:   nil,
		Children: nil,
	}
	iFile := &node{
		Isdir:    false,
		Name:     "i",
		Size:     584,
		Parent:   nil,
		Children: nil,
	}
	fFile := &node{
		Isdir:    false,
		Name:     "f",
		Size:     29116,
		Parent:   nil,
		Children: nil,
	}
	gFile := &node{
		Isdir:    false,
		Name:     "g",
		Size:     2557,
		Parent:   nil,
		Children: nil,
	}
	hLst := &node{
		Isdir:    false,
		Name:     "h.lst",
		Size:     62596,
		Parent:   nil,
		Children: nil,
	}
	jFile := &node{
		Isdir:    false,
		Name:     "j",
		Size:     4060174,
		Parent:   nil,
		Children: nil,
	}
	dLog := &node{
		Isdir:    false,
		Name:     "d.log",
		Size:     8033020,
		Parent:   nil,
		Children: nil,
	}
	dExt := &node{
		Isdir:    false,
		Name:     "d.ext",
		Size:     5626152,
		Parent:   nil,
		Children: nil,
	}
	kFile := &node{
		Isdir:    false,
		Name:     "k",
		Size:     7214296,
		Parent:   nil,
		Children: nil,
	}
	dDir := &node{
		Isdir:    true,
		Name:     "d",
		Size:     24933642,
		Parent:   nil,
		Children: nil,
	}
	aDir := &node{
		Isdir:    true,
		Name:     "a",
		Size:     94853,
		Parent:   nil,
		Children: nil,
	}
	root := &node{
		Isdir:    true,
		Name:     "/",
		Size:     48381165,
		Parent:   nil,
		Children: nil,
	}

	tests := []struct {
		testName  string
		inputNode *node
		want      int
	}{
		{
			testName:  "small sample",
			inputNode: generateRootNode(t, root, aDir, eDir, iFile, fFile, gFile, hLst, bTxt, cDat, dDir, jFile, dLog, dExt, kFile),
			want:      95437,
		},
	}

	for _, tc := range tests {
		_ = findBigDirs(tc.inputNode)
		if totalSize != tc.want {
			t.Errorf("findBigDirs() failed. %v != %v", totalSize, tc.want)
		}
	}
}

func TestAddSizeUpToRoot(t *testing.T) {
	bTxt := &node{
		Isdir:    false,
		Name:     "b.txt",
		Size:     14848514,
		Parent:   nil,
		Children: nil,
	}
	cDat := &node{
		Isdir:    false,
		Name:     "c.dat",
		Size:     8504156,
		Parent:   nil,
		Children: nil,
	}
	eDir := &node{
		Isdir:    true,
		Name:     "e",
		Size:     0,
		Parent:   nil,
		Children: nil,
	}
	// iFile := &node{
	// 	Isdir:    true,
	// 	Name:     "i",
	// 	Size:     0,
	// 	Parent:   nil,
	// 	Children: nil,
	// }
	// fFile := &node{
	// 	Isdir:    true,
	// 	Name:     "f",
	// 	Size:     0,
	// 	Parent:   nil,
	// 	Children: nil,
	// }
	// gFile := &node{
	// 	Isdir:    true,
	// 	Name:     "g",
	// 	Size:     0,
	// 	Parent:   nil,
	// 	Children: nil,
	// }
	// hLst := &node{
	// 	Isdir:    true,
	// 	Name:     "h",
	// 	Size:     0,
	// 	Parent:   nil,
	// 	Children: nil,
	// }
	aDir := &node{
		Isdir:    true,
		Name:     "a",
		Size:     0,
		Parent:   nil,
		Children: nil,
	}
	root := &node{
		Isdir:    true,
		Name:     "/",
		Size:     0,
		Parent:   nil,
		Children: nil,
	}

	tests := []struct {
		testName  string
		inputNode *node
		want      int
	}{
		{
			testName:  "small sample",
			inputNode: generateSmallDeepRootNode(t, root, aDir, eDir, bTxt, cDat),
			want:      23352670,
		},
	}

	for _, tc := range tests {
		ppChildren(tc.inputNode)
		addSizeUpToRoot(bTxt.Size, deepSeek(tc.inputNode, "b.txt"))
		addSizeUpToRoot(cDat.Size, deepSeek(tc.inputNode, "c.dat"))
		ppChildren(tc.inputNode)
		if tc.inputNode.Size != tc.want {
			t.Errorf("addSizeUpToRoot() failed. %v != %v", tc.inputNode.Size, tc.want)
		}
	}
}

func generateRootNode(t *testing.T, rootNode, aNode, eNode, iNode, fNode, gNode, hNode, bNode, cNode, dNode, jNode, dlogNode, dExtNode, kNode *node) *node {
	t.Helper()
	rootNode.Children = append(rootNode.Children, aNode, dNode, bNode, cNode)
	aNode.Children = append(aNode.Children, eNode, fNode, gNode, hNode)
	dNode.Children = append(dNode.Children, jNode, dlogNode, dExtNode, kNode)
	eNode.Children = append(eNode.Children, iNode)

	aNode.Parent = rootNode
	bNode.Parent = rootNode
	cNode.Parent = rootNode
	dNode.Parent = rootNode
	eNode.Parent = aNode
	iNode.Parent = eNode
	fNode.Parent = aNode
	gNode.Parent = aNode
	hNode.Parent = aNode
	jNode.Parent = dNode
	dlogNode.Parent = dNode
	dExtNode.Parent = dNode
	kNode.Parent = dNode

	return rootNode
}

func generateDeepRootNode(t *testing.T, rootNode, aNode, eNode, iNode, fNode, gNode, hNode, bNode, cNode *node) *node {
	t.Helper()
	rootNode.Children = append(rootNode.Children, aNode)
	aNode.Children = append(aNode.Children, eNode)
	eNode.Children = append(eNode.Children, iNode)
	iNode.Children = append(eNode.Children, fNode)
	fNode.Children = append(eNode.Children, gNode)
	gNode.Children = append(eNode.Children, hNode)
	hNode.Children = append(eNode.Children, bNode, cNode)

	aNode.Parent = rootNode
	eNode.Parent = aNode
	iNode.Parent = eNode
	fNode.Parent = iNode
	gNode.Parent = fNode
	hNode.Parent = gNode
	bNode.Parent = hNode
	cNode.Parent = bNode

	return rootNode
}

func generateSmallDeepRootNode(t *testing.T, rootNode, aNode, eNode, bNode, cNode *node) *node {
	t.Helper()
	rootNode.Children = append(rootNode.Children, aNode)
	aNode.Children = append(aNode.Children, eNode)
	eNode.Children = append(eNode.Children, bNode, cNode)

	aNode.Parent = rootNode
	eNode.Parent = aNode
	bNode.Parent = eNode
	cNode.Parent = eNode

	return rootNode
}
