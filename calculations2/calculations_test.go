package calculations

import (
	"testing"
)

func TestSI(t *testing.T) {
	got := calculateSI(42000.00, 2.5, 5.00)
	want := "5250.00"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "5250.00")
	}

	got = calculateSI(36000.00, 3.75, 5)
	want = "6750.00"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "6750.00")
	}
}

func TestCI(t *testing.T) {

	got := calculateCI(42000.00, 2.5, 5.00)
	want := "5519.14"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "5519.14")
	}

	got = calculateCI(36000.00, 3.75, 5)
	want = "7275.59"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "7275.59")
	}
}
