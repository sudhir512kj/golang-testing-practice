package tests

import (
	"testing"
)

func TestFacto(t *testing.T) {

	got := facto(5)
	want := 120

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, 120)
	}

	got = facto(0)
	want = 1

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, 1)
	}
}
