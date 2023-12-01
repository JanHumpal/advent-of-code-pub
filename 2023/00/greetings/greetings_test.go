package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Jeff"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Jeff")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Jeff") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
