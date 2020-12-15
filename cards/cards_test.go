package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()
	assert.NotNil(t, d)
	assert.Equal(t, len(d), 16)
}

func TestDealDeck(t *testing.T) {
	d := createDeck()
	left, right := d.deal(10)
	assert.NotNil(t, left)
	assert.NotNil(t, right)
	assert.Equal(t, 10, len(left))
	assert.Equal(t, 6, len(right))
}
