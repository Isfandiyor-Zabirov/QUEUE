package main

import "fmt"

type List struct {
	root   *ListNode // first element
	end    *ListNode // last element
	length int
}

func (receiver *List) PrintList() {
	fmt.Println("Start Print...")
	current := receiver.root
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
	fmt.Println("Finish Print...")

}

func (receiver *List) len() int {
	return receiver.length
}

func (receiver *List) Add(node ListNode) {
	if receiver.len() == 0 {
		receiver.root = &node
		receiver.end = &node
		receiver.length++
		return
	}
	LastElement := receiver.end
	receiver.end.Next = &node
	receiver.end = &node
	receiver.end.Prev = LastElement
	receiver.length++
}

func (receiver *List) Remove(node ListNode)  {
	receiver.length--

}

type ListNode struct {
	Prev      *ListNode
	Next      *ListNode
	Name      string
	Purchases int
}

func main() {
	person := ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Isfandak",
		Purchases: 97,
	}
	queue := List{}
	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Zacken",
		Purchases: 197,
	}

	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Umed",
		Purchases: 568,
	}

	queue.Add(person)
	fmt.Println(queue)

	queue.PrintList()

	fmt.Println(queue.root)
	fmt.Println(queue.end)
}
