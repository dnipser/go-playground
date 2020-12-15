package main

import "fmt"

func main() {
	d := createDeck()
	leftDeck, rightDeck := d.deal(8)
	fmt.Println(leftDeck)
	fmt.Println(rightDeck)

	leftDeck.serialize("ld.txt")
	restored := leftDeck.deserialize("ld.txt")
	fmt.Println(restored)

	shuffled := rightDeck.shuffle()
	fmt.Println(shuffled)
}
