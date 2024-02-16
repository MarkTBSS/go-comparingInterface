package greeting

import "fmt"

type IGreeting interface {
	Greeting(string)
}

type greetingInteface struct {
}

func NewGreetingInterfaceInstance() IGreeting {
	return &greetingInteface{}
}

func (h *greetingInteface) Greeting(name string) {
	fmt.Printf("Greeting, %v!\n", name)
}
