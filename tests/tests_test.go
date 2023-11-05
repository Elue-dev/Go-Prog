package main

import (
	"testing"
)

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expectation float32
	shouldHaveErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision (t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.shouldHaveErr {
			if err == nil {
				t.Error("Expected an error but did not get one ❌")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but get one ❌", err.Error())
			}
		}

		if got != tt.expectation {
			t.Errorf("Epected %f but got %f", tt.expectation, got)
		}
	}


}

// func TestDivideFunc(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should't have ✅")
// 	}
// }

// func TestBadDivideFunc(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have ❌")
// 	}
// }