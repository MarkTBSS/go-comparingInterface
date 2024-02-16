package hello

import "fmt"

type hello struct {
}

func (h *hello) Hello(name string) {
	fmt.Printf("Hello, %v!\n", name)
}

func NewHelloPointer() *hello {
	return &hello{}
}

func NewHello() hello {
	return hello{}
}

func (h *hello) HelloFriendString(name string) {
	fmt.Printf("Hello, %v!\n", name)
	// Change Hello to Helloooo
	//fmt.Printf("Helloooo, %v!\n", name)
}
