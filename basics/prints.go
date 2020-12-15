package main

import (
	"fmt"
	"os"
	"strings"
)

// variable declaration without assignment
var packageLevelOne int64

// variable declaration with assignment
var packageLevelTwo int64 = 42

func printQuotes() {
	fmt.Println("Let's play with Go!!!")
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
	printQuotes()
	printEnv()
	printVars()
}
