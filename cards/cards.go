package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var suites = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var values = []string{"Ace", "Two", "Three", "Four"}

func createDeck() deck {
	result := deck{}
	for _, suite := range suites {
		for _, val := range values {
			result = append(result, fmt.Sprintf("%s of %s", val, suite))
		}
	}
	return result
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	result := make(deck, len(d))
	copy(result, d)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return result
}

func (d deck) String() string {
	result := strings.Join(d, ", ")
	return result
}

func (d deck) serialize(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.String()), 0666)
}

func (d deck) deserialize(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Couldn't read file ", filename)
		os.Exit(1)
	}
	return strings.Split(string(content), ", ")
}
