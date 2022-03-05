package main

import (
	"fmt"
	"testing"
)

func TestScraper(t *testing.T) {
	site := ScrapeTitle("https://stockx.com/sneakers/most-popular")
	if site != "Buy & Sell Deadstock Shoes - Most Popular" {
		t.Errorf("Expected 'Buy & Sell Deadstock Shoes - Most Popular', got %s", site)
	}
	fmt.Println(site)
}

// create benchmark test
func BenchmarkScraper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScrapeTitle("https://stockx.com/sneakers/most-popular")
	}
}
