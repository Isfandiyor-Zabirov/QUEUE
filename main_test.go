package main

import "testing"

func TestList_Add(t *testing.T) {
	q := List{}

	if q.length != 0 {
		want := 0
		got := q.length
		t.Errorf("method len() in Method (ADD) is not correct for empty queue got %d want %d", got, want)
	}
}
