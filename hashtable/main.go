package main

import (
	"bytes"
	"fmt"
)

const MAP_SIZE = 10

type Node struct {
	key   string
	value string
	next  *Node
}

type HashMap struct {
	Data []*Node
}

func NewDict() *HashMap {
	return &HashMap{Data: make([]*Node, MAP_SIZE)}
}

func (n *Node) String() string {
	return fmt.Sprintf("<Key: %s, Value: %s>\n", n.key, n.value)
}

func (h *HashMap) String() string {
	var output bytes.Buffer
	fmt.Fprintln(&output, "{")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "\t%s: %s\n", n.key, n.value)
			for node := n.next; node != nil; node = node.next {
				fmt.Fprintf(&output, "\t%s: %s\n", node.key, node.value)
			}
		}
	}

	fmt.Fprintln(&output, "}")

	return output.String()
}

//func (h *HashMap) Insert(key string, value string) {

//func (h *HashMap) Get(key string) (string, bool) {

//func hash(key string) (hash uint8) {

func getIndex(key string) (index int) {
	return int(hash(key)) % MAP_SIZE
}

func main() {
	a := NewDict()
	a.Insert("name", "ishan")
	a.Insert("gender", "male")
	a.Insert("city", "mumbai")
	a.Insert("lastname", "khare")
	if value, ok := a.Get("name"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value did not match!")
	}

	fmt.Println(a)
}
