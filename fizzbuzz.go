package fizzbuzz

import "strconv"

func Test(i int) string {
	r := strconv.Itoa(i)
	mod3 := i % 3
	mod5 := i % 5
	switch {
	case mod3 == 0 && mod5 == 0:
		r = "fizzbuzz"
	case mod3 == 0:
		r = "fizz"
	case mod5 == 0:
		r = "buzz"
	}
	return r
}
