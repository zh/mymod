package serverlib

import "testing"

func TestHello(t *testing.T) {
	want := "serverlib hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
