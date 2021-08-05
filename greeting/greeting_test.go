package greeting

import "testing"

// func TestGreetingYourName(t *testing.T) {
// 	given := ""
// 	want := "Hello, my friend."

// 	get := Greet(given)

// 	if want != get {
// 		t.Errorf("given a name %s want greeting %q but got %q", given, want, get)
// 	}
// }

func TestGreetingCapital(t *testing.T) {
	given := "BOB"
	want := "Capital"

	get := Greet(given)

	if want != get {
		t.Errorf("given a name %s want greeting %q but got %q", given, want, get)
	}
}
