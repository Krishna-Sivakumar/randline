package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed wordslist.txt
var lineList []byte

func TestRandomness(t *testing.T) {
	// read random lines from a sample file of 100 lines, 1000 times.
	// the probability of each line must be uniform, with an error of 5%.
	benchmarkProbabilityRange := []float64{0.0095, 0.0105}

	lineCount := make(map[string]int)
	for i := 0; i < 1000; i++ {
		chosenLine, err := readFromByteBuffer(lineList)
		if err != nil {
			t.Fatal(err)
		}
		lineCount[chosenLine] += 1
	}

	var avgProbability float64 = 0
	for _, count := range lineCount {
		probability := float64(count) / float64(100)
		avgProbability += probability
	}
	avgProbability /= float64(1000)
	fmt.Println("Average Probability", avgProbability)

	if avgProbability > benchmarkProbabilityRange[1] || avgProbability < benchmarkProbabilityRange[0] {
		t.Fatalf("Failed the uniform probability test. Has an average probability of %f. The whole map: %v", avgProbability, lineCount)
	}
}
