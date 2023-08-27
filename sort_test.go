package main

import (
	"log"
	"testing"
)

var testData []RandomUser

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	ru, err := getData()
	if err != nil {
		log.Fatal(err)
	}
	testData = ru
}

func BenchmarkOneSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortUsingSort(testData)
	}
}

func BenchmarkTwoSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortUsingSlices(testData)
	}
}
