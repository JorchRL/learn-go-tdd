package main

import "testing"

func TestHello (t *testing.T) {
	got := Hello("Jorch")
	want := "Hello, Jorch!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}