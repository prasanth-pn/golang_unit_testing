package main

import "testing"

func Test_newDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected lenth is 16 but got %d", len(d))
	}
	if d[0]!="AceofSpades"{
		t.Errorf("Espected first card isAceofSpades but got %v",d[0])
	}
	if d[len(d)-1]!="FourofClubs"{
		t.Errorf("Expected last card is Four of Clubs but got %v",d[len(d)-1])
	}
}