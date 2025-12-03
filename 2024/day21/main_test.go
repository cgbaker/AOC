package main

import (
	"testing"
)

func TestBestInput(t *testing.T) {
	for _, s := range []string{"029A", "980A", "179A", "456A", "379A"} {
		ExploreOptimalMoves(1, s)
	}
}
