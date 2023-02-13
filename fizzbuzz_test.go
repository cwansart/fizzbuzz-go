package fizzbuzz_test

import (
	"testing"

	"github.com/cwansart/fizzbuzz-go"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want string
	}{
		{"fizz for 3", 3, "fizz"},
		{"fizz for multiples of 3", 6, "fizz"},
		{"fizz for multiples of 3", 9, "fizz"},
		{"fizz for multiples of 3", 12, "fizz"},
		{"buzz for 5", 5, "buzz"},
		{"buzz for multiples of 5", 10, "buzz"},
		{"buzz for multiples of 5", 20, "buzz"},
		{"buzz for multiples of 5", 25, "buzz"},
		{"fizzbuzz for 15", 15, "fizzbuzz"},
	}

	for _, tt := range tests {
		want := tt.want
		got := fizzbuzz.Test(tt.i)
		if want != got {
			t.Errorf("got '%s' but wanted '%s'", got, want)
		}
	}
}
