package greetings

import (
	"regexp"
	"testing"
)

func TestHellName(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan")= %q, %v, quiere coincidecia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")= %q, %v, quiere coincidecia "", nil`, msg, err)
	}
}
