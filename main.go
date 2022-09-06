package main

import (
	"fmt"
)

type LinkedList struct {
	list   *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func createNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func (l *LinkedList) add(data int) {
	l.list = l.addToLinkedList(l.list, data)
	l.length++
	return
}
func (l *LinkedList) addToLinkedList(node *Node, data int) *Node {
	if node == nil {
		node = createNode(data)
	} else {
		node.next = l.addToLinkedList(node.next, data)
	}
	return node

}
func (l *LinkedList) remove(data int) {
	l.list = l.removeNode(l.list, data)
	return
}
func (l *LinkedList) removeNode(node *Node, data int) *Node {
	if node == nil {
		return nil
	} else if node.next.data == data {
		node.next = node.next.next
		l.length--
		return node
	}
	node.next = l.removeNode(node.next, data)
	return node

}
func (l *LinkedList) update(data int, newData int) {
	l.list = l.updateNode(l.list, data, newData)
	return
}
func (l *LinkedList) updateNode(node *Node, data int, newData int) *Node {
	if node == nil {
		return nil
	} else if node.data == data {
		node.data = newData
		return node
	}
	node.next = l.updateNode(node.next, data, newData)
	return node

}
func (l *LinkedList) DisplayList(node *Node) {

	if node != nil {
		fmt.Printf("%+v ->", node.data)
		l.DisplayList(node.next)
	}
	return
}

func (l *LinkedList) ShowBackwards(list *Node) {
	if list != nil {
		l.ShowBackwards(list.next)
		fmt.Printf("%v ->", list.data)
	}
	return
}

func main() {
	link := LinkedList{}
	link.add(1)
	link.add(2)
	link.add(3)
	link.add(4)
	link.ShowBackwards(link.list)
	link.removeNode(link.list, 3)
	link.update(1, 5)
	fmt.Println()
	link.DisplayList(link.list)
	fmt.Println()
	fmt.Println("list length", link.length)
}
