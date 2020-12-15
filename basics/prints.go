package main

import (
	"fmt"
	"os"
	"strings"

	"rsc.io/quote"
)

// variable declaration without assignment
var packageLevelOne int64

// variable declaration with assignment
var packageLevelTwo int64 = 42

func printQuotes() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
}

func printEnv() {
	for _, val := range os.Environ() {
		pair := strings.Split(val, "=")
		fmt.Println(pair[0], pair[1])
	}
}

func printVars() {
	fmt.Printf("type: %T, val: %#v", packageLevelOne, packageLevelOne)
	fmt.Printf("type: %T, val: %#v", packageLevelTwo, packageLevelTwo)
}

func runPrintsExample() {
	fmt.Println("Hello, World!")

	printQuotes()
	printEnv()
	printVars()
}
