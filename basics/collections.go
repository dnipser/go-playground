package main

import "fmt"

func printSliceElements() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}

func printMapElements() {
	holidays := map[string]int{
		"Jan": 2, "Mar": 1, "May": 3,
		"Jun": 2, "Aug": 1, "Oct": 1,
		"Dec": 30, "Apr": 22,
	}

	holidays["Dec"] = 1
	delete(holidays, "Apr")

	for month, days := range holidays {
		fmt.Printf("There are %d holidays in %s\n", days, month)
	}
}

func runCollectionsExample() {
	printSliceElements()
	printMapElements()
}
