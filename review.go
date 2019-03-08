package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data   int
	Height int
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func balance_factor(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func (self *Node) insert(data int) *Node {
	if self == nil {
		return &Node{Data: data, Height: 1}
	}
	switch {

	case data < self.Data:
		self.Left = self.Left.insert(data)

	case data > self.Data:
		self.Right = self.Right.insert(data)
	}
	// Update the height.
	self.Height = 1 + max(height(self.Left), height(self.Right))

	// Rotations
	if balance_factor(self) > 1 && data < self.Left.Data { // Left Left
		self = self.rotate_right()
	} else if balance_factor(self) < -1 && data > self.Right.Data { // Right Right
		self = self.rotate_left()
	} else if balance_factor(self) > 1 && data > self.Left.Data { // Left Right
		self.Left = self.Left.rotate_left()
		self = self.rotate_right()
	} else if balance_factor(self) < -1 && data < self.Right.Data { // Right Left
		self.Right = self.Right.rotate_right()
		self = self.rotate_left()
	}
	return self
}

func (self *Node) rotate_left() *Node {
	if self == nil {
		return nil
	}
	new_parent := self.Right
	self.Right = new_parent.Left
	new_parent.Left = self

	// Update heights.
	self.Height = max(height(self.Left), height(self.Right)) + 1
	new_parent.Height = max(height(new_parent.Left), height(new_parent.Right)) + 1
	return new_parent
}

func (self *Node) rotate_right() *Node {
	if self == nil {
		return nil
	}
	new_parent := self.Left
	self.Left = new_parent.Right
	new_parent.Right = self

	// Update heights.
	self.Height = max(height(self.Left), height(self.Right)) + 1
	new_parent.Height = max(height(new_parent.Left), height(new_parent.Right)) + 1
	return new_parent
}

func (self *Node) preorder_traverse() {
	if self == nil {
		return
	}
	fmt.Printf("(%d, %d) ", self.Data, balance_factor(self))
	self.Left.preorder_traverse()
	self.Right.preorder_traverse()
}

func (self *Tree) preorder_traverse() error {
	if self == nil {
		return errors.New("EMPTY TREE: Cannot traverse uninitilized tree.")
	}
	fmt.Printf("Preorder (node, balance_factor): ")
	self.Root.preorder_traverse()
	fmt.Printf("\n")
	return nil
}

func (self *Tree) insert(data int) {
	self.Root = self.Root.insert(data)
}

func main() {
	tree := Tree{}
	elem := [13]int{21, 26, 30, 9, 4, 14, 28, 18, 15, 10, 2, 3, 7}
	for _, num := range elem {
		tree.insert(num)
	}
	if err := tree.preorder_traverse(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
