package greet_test

import (
	"github.com/quii/go-mod-example/greet"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, World!"
	if got := greet.TheWorld(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
