package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	total := len(suites) * len(values)

	deck := createDeck()
	assert.NotNil(t, deck)
	assert.Equal(t, len(deck), total)

	for _, suite := range suites {
		assert.Truef(t, elementContains(deck, suite), "Suite %s is not present in deck", suite)
	}
	for _, val := range values {
		assert.Truef(t, elementContains(deck, val), "Value %s is not present in deck", val)
	}
}

func TestDealDeck(t *testing.T) {
	handSize := 10
	total := len(suites) * len(values)

	deck := createDeck()
	left, right := deck.deal(handSize)
	assert.NotNil(t, left)
	assert.NotNil(t, right)
	assert.Equal(t, handSize, len(left))
	assert.Equal(t, total-handSize, len(right))
}

func TestShuffleDeck(t *testing.T) {
	original := createDeck()
	shuffled := original.shuffle()

	for _, deckVal := range original {
		assert.Containsf(t, shuffled, deckVal, "Shuffled deck doesn't contain '%s'", deckVal)
	}
}

func TestSerialization(t *testing.T) {
	fileName := "ld.txt"
	// remove file if it exists
	os.Remove(fileName)
	// cleanup file on test completion
	defer os.Remove(fileName)

	deck := createDeck()
	deck.serialize(fileName)
	restored := deck.deserialize(fileName)

	assert.Equal(t, deck, restored, "Deserialized deck differs from original")
}

func elementContains(input []string, elementSubstr string) bool {
	for _, val := range input {
		if strings.Contains(val, elementSubstr) {
			return true
		}
	}
	return false
}
