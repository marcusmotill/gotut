package heap

import (
	"fmt"
	"strconv"
)

type tree struct {
	Left      *tree
	Value     int
	Right     *tree
	Root      []int
	Postition int
}

func (t tree) String() string {
	var str = ""
	for i := range t.Root {
		var node *tree
		if i == 0 {
			node = t.getNode("")
		} else {
			node = t.getNode(getPath(i + 1))
		}

		if node.Left == nil && node.Right == nil {
			continue
		}

		if node.Left != nil {
			str = str + fmt.Sprintf("%v <- ", node.Left.Value)
		}

		str = str + fmt.Sprintf("%v", node.Value)

		if node.Right != nil {
			str = str + fmt.Sprintf(" -> %v", node.Right.Value)
		}

		str = str + "\n"
	}
	return str
}

// Sort run heap sort on array of int
func Sort(arr []int) []int {

	var root = tree{Root: arr}

	root.heapify(root.Root)

	for i := range root.Root {
		if i == len(root.Root)-1 {
			break
		}
		root.maxHeapify(len(root.Root))
		//fmt.Printf("%v \n", root.Root)

		// swap root and most leaf
		rootNode := root.getNode(getPath(-1))
		leafNode := root.getNode(getPath(len(root.Root) - i))

		root.swap(leafNode, rootNode)

		//delete new most leaf from tree
		leafPath := getPath(len(root.Root) - i)
		parentPath := leafPath[:len(leafPath)-1]
		parentNode := root.getNode(parentPath)

		if leafPath[len(leafPath)-1] == '0' {
			parentNode.Left = nil
		} else if leafPath[len(leafPath)-1] == '1' {
			parentNode.Right = nil
		}
	}

	return root.Root
}

func (t *tree) validateMaxHeap() {
	for i := range t.Root {
		var node *tree
		if i == 0 {
			node = t.getNode("")
		} else {
			node = t.getNode(getPath(i + 1))
		}

		if t.Root[node.Postition] != node.Value {
			fmt.Printf("swap error! %v expected at %d got %d\n", t.Root[node.Postition], node.Postition, node.Value)
		}

		if node.Left == nil && node.Right == nil {
			continue
		}

		if node.Left != nil {
			if node.Left.Value > node.Value {
				fmt.Printf("left error! %v %d\n", node.Value, node.Left.Value)
			}
		}

		if node.Right != nil {
			if node.Right.Value > node.Value {
				fmt.Printf("right error! %v %d\n", node.Value, node.Right.Value)
			}
		}
	}
}

func (t *tree) heapify(arr []int) {

	for i, value := range arr {
		if i == 0 {
			t.insert("", value, i)
		} else {
			t.insert(getPath(i+1), value, i)
		}
	}
}

func (t *tree) maxHeapify(startInt int) {

	if startInt == 0 {
		return
	}

	nodePath := getPath(startInt)
	node := t.getNode(nodePath)

	//find parent node if not currently on the root
	if len(nodePath) > 0 {

		parentPath := nodePath[:len(nodePath)-1]
		parentNode := t.getNode(parentPath)

		if node.Value > parentNode.Value {
			t.swap(node, parentNode)
			t.maxHeapify(len(t.Root))
			return
		}
	}

	t.maxHeapify(startInt - 1)
	return

}

func (t *tree) insert(path string, val int, postition int) {

	node := t.getNode(path)
	node.Value = val
	node.Postition = postition
}

func (t *tree) swap(node1, node2 *tree) {

	//swap Values
	node1Val := node1.Value
	node2Val := node2.Value

	node1.Value = node2Val
	node2.Value = node1Val

	node1Position := node1.Postition
	node2Position := node2.Postition

	//swap positions in array
	val1 := t.Root[node1Position]
	val2 := t.Root[node2Position]

	t.Root[node1Position] = val2
	t.Root[node2Position] = val1

}

func (t *tree) getNode(path string) *tree {
	currentTree := t

	for _, value := range path {
		if currentTree.Left == nil {
			currentTree.Left = &tree{}
		}

		if currentTree.Right == nil {
			currentTree.Right = &tree{}
		}

		if value == '0' {
			currentTree = currentTree.Left
		} else {
			currentTree = currentTree.Right
		}
	}

	return currentTree
}

func getPath(x int) string {
	if x < 0 {
		return ""
	}

	// int to binary, drop leading 1.  0 means left 1 means right on tree traversal
	var binary = strconv.FormatInt(int64(x), 2)
	if binary[0] == '1' {
		binary = binary[1:]
	}

	return binary
}
