package main

import "testing"

func TestMinimumBribes1(t *testing.T) {
	q := []int32{2, 1, 5, 3, 4}

	minimumBribes(q)
}

// output should be "3"

func TestMinimumBribes2(t *testing.T) {
	q := []int32{2, 5, 1, 3, 4}

	minimumBribes(q)
}

// output should be "Too chaotic"

func TestMinimumBribes3(t *testing.T) {
	q := []int32{5, 1, 2, 3, 7, 8, 6, 4}

	minimumBribes(q)
}

// output should be "Too chaotic"

func TestMinimumBribes4(t *testing.T) {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	minimumBribes(q)
}

// output should be "7"
