package mymod

import "testing"

func TestConfig(t *testing.T) {
	want := "mymod config"
	if got := Config(); got != want {
		t.Errorf("Config() = %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
