package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"alphabet", "tebahpla"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			got := reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", tc.word, tc.want)
			}
		})
	}
}
