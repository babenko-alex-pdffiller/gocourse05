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

func (h *HashMap) Insert(key string, value string) {
	index := getIndex(key)

	if h.Data[index] == nil {
		// index is empty, go ahead and insert
		h.Data[index] = &Node{key: key, value: value}
	} else {
		// there is a collision, get into linked-list mode
		starting_node := h.Data[index]
		for ; starting_node.next != nil; starting_node = starting_node.next {
			if starting_node.key == key {
				// the key exists, its a modifying operation
				starting_node.value = value
				return
			}
		}
		starting_node.next = &Node{key: key, value: value}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in linked list
		starting_node := h.Data[index]
		for ; ; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				return starting_node.value, true
			}

			if starting_node.next == nil {
				break
			}
		}
	}

	// key does not exists
	return "", false
}

func hash(key string) (hash uint8) {
	// a jenkins one-at-a-time-hash
	// refer https://en.wikipedia.org/wiki/Jenkins_hash_function

	hash = 0
	for _, ch := range key {
		hash += uint8(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

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
