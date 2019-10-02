package mymod

import "testing"

func TestConfig(t *testing.T) {
	want := "mymod config"
	if got := Config(); got != want {
		t.Errorf("Config() = %q, want %q", got, want)
	}
}
