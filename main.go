package main

import "fmt"

// type NodeList struct {
// 	length     int
// 	nextNode   string
// 	backNode   string
// 	getCurrent string
// }

type Node struct {
	data string
	back *Node
	next *Node
}

func (n Node) nextNode() Node {
	return *n.next
}

func (n Node) backNode() Node {
	return *n.back
}

//	func (n Node) newNode(b *Node, ne *Node, d string) Node {
//		n.back = b
//		n.next = ne
//		n.data = d
//		return n
//	}
func main() {
	node1 := &Node{data: "a"}
	node2 := &Node{data: "b", back: node1}
	node3 := &Node{data: "c", back: node2}
	node1.next = node2
	node2.next = node3
	fmt.Println("hade node ", node1)
	fmt.Println("nex from node 1", node1.nextNode())
	fmt.Println("node 2", node2)
	fmt.Println("last node ", node3)
}
