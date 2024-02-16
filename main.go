package main

import (
	"github.com/MarkTBSS/go-comparingInterface/greeting"
	"github.com/MarkTBSS/go-comparingInterface/hello"
)

func main() {
	// Initialize the Hello service
	helloService := hello.NewHelloPointer()
	helloService2 := hello.NewHello()

	// Invoke the HelloFriend method
	helloService.HelloFriendString("helloService1")
	helloService2.HelloFriendString("helloServices2")

	greetingService := greeting.NewGreetingInterfaceInstance()
	greetingService.Greeting("greetingService")
}
