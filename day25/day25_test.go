package main

import (
	"testing"
)

func TestHandshake(t *testing.T) {
	cardPubKey := 5764801
	doorPubKey := 17807724

	result := findEncryptionKey(cardPubKey, doorPubKey)

	if result != 14897079 {
		t.Errorf("Wrong encryption key, expected 14897079, got %v",
			result)
	}
}
