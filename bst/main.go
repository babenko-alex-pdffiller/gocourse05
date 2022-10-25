package main

import "fmt"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

//func (b *bst) insert(value int) {

//func (b *bst) insertRec(node *bstnode, value int) *bstnode {

//func (b *bst) find(value int) error {

//func (b *bst) findRec(node *bstnode, value int) *bstnode {

func (b *bst) inorder() {
	b.inorderRec(b.root)
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
		b.inorderRec(node.right)
	}
}

func main() {
	bst := &bst{}
	eg := []int{2, 5, 7, -1, -1, 5, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("Printing Inorder:\n")
	bst.inorder()
	bst.reset()
	eg = []int{4, 5, 7, 6, -1, 99, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("\nPrinting Inorder:\n")
	bst.inorder()
	fmt.Printf("\nFinding Values:\n")
	err := bst.find(2)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 2)
	} else {
		fmt.Printf("Value %d Found\n", 2)
	}
	err = bst.find(6)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 6)
	} else {
		fmt.Printf("Value %d Found\n", 6)
	}
	err = bst.find(5)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 5)
	} else {
		fmt.Printf("Value %d Found\n", 5)
	}
	err = bst.find(1)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 1)
	} else {
		fmt.Printf("Value %d Found\n", 1)
	}
}
