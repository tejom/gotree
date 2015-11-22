package main

import "fmt"

func main() {

	var node = Tree{Value:4}
	node.addNode(1)
	node.addNode(2)
	node.addNode(6)
	node.addNode(3)
	node.addNode(10)
	node.smallest()
	node.largest()
	node.addNode(9)
	node.addNode(11)
	node.largest()
	fmt.Println("pringting tree")
	node.printTree()

	

}

func addOne(v *int) int {
	*v = *v+1
	return *v
}

type Tree struct {
	Value 		int
	Left,Right 	*Tree //pointer to a tree
	Parent 		*Tree //pointer to a tree
}

func (node *Tree) addNode(v int) {
	var newNode = Tree{Value:v}
	if v > node.Value { //go right
		if node.Right ==nil{ 
			node.Right = &newNode //assign in its a leaf
		} else {
			node.Right.addNode(v) //recusrse though right
		}
	} else {
		if node.Left ==nil {
			node.Left = &newNode
		} else {
			node.Left.addNode(v)
		}
	}
}

func (node *Tree) smallest() {
	if node.Left == nil {
		fmt.Println(node.Value)
	} else {
		node.Left.smallest()
	}
}

func (node *Tree) largest() {
	if node.Right == nil {
		fmt.Println(node.Value)
	} else {
		node.Right.largest()
	}
}

func (node *Tree) printTree() {
	
	if node.Left != nil {
		node.Left.printTree()
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		node.Right.printTree()
	}
	
}