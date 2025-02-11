package main

import "fmt"

// array list
type ArrayList struct {
	capacity int
	count    int
	element  []byte
}

type Stack struct {
	capacity int
	count int
	items []int
}

type Queue struct {
	startIndex int
	endInded int
	capacity int
	items []int
}

func main() {
	arrayList := &ArrayList{}
	fmt.Println(arrayList)

	stack := &Stack{
		capacity: 10,
		count: 5,
		items: []int{1,2,3,4,5},
	}
	fmt.Println(stack)
	// get lenght of the stack
	lengthStack := stack.count
	fmt.Printf("Length of the stack %d\n", lengthStack)
}

func (s *Stack) Add(elem int) {
	s.count += 1
	
}
