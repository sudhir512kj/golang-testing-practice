package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Sudhir"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := Hello(name)
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Sudhir") = %q, want match for %#q, nil`, msg, want)
	}
}

// func TestH()  {

// }
