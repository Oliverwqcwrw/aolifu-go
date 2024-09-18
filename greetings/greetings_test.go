package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// HellosMultipleNames tests the Hellos function with multiple names.
// It verifies that the function returns a message for each name without errors.
func TestHellosMultipleNames(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf("Hellos(%v) returned error: %v", names, err)
	}
	for _, name := range names {
		if msg, ok := messages[name]; !ok || msg == "" {
			t.Fatalf("Hellos(%v) did not return a message for %v", names, name)
		}
	}
}

func TestHellosEmptySlice(t *testing.T) {
	names := []string{}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf("Hellos(%v) returned error: %v", names, err)
	}
	if len(messages) != 0 {
		t.Fatalf("Hellos(%v) returned non-empty map: %v", names, messages)
	}
}

func TestHellosWithEmptyName(t *testing.T) {
	names := []string{"Alice", ""}
	_, err := Hellos(names)
	if err == nil {
		t.Fatalf("Hellos(%v) did not return an error", names)
	}
}
