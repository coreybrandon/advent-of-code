package main

import "testing"

func TestIncreasedDepth(t *testing.T) {
	sample := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	got := IncreasedDepth(sample)
	want := 7

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, sample)
	}
}

func TestPartTwo(t *testing.T) {
	sample := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	got := PartTwo(sample)
	want := 5

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, sample)
	}
}
