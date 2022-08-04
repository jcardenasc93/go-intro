package helpers

import "testing"

var tests = []struct {
	name       string
	dividend   float32
	divisor    float32
	expected   float32
	raiseError bool
}{
	{"validInput", 100.0, 10.0, 10.0, false},
	{"validInput", 100.0, 0.0, 0.0, true},
}

//NOTE: Run go test -cover to check test coverage

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := DivideFloats(tt.dividend, tt.divisor)
		// First check the errors edge cases
		if tt.raiseError {
			if err == nil {
				t.Error("Error not raised")
			}
		} else {
			if err != nil {
				t.Error("Unexpected error raised")
			}
		}
		// Then look for expected values cases
		if got != tt.expected {
			t.Errorf("Division test result %f is not %f", got, tt.expected)
		}

	}
}
