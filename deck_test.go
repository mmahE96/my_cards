package main

//file for all test code

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected ... but got %d", len(d))
	}

}
