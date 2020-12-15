package main

import "fmt"

var simpleGreeting = "Hello from sandbox"
var anotherGreeting = "Alternative hello from sandbox"

func getGreeting() string {
	return simpleGreeting
}

func getGreetings() []string {
	return []string{simpleGreeting, anotherGreeting}
}

func main()  {
	singleVal := getGreeting()
	fmt.Println(singleVal)
	multiVal := getGreetings()
	fmt.Printf("%#v", multiVal)
}