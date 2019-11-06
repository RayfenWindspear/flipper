package flipper

import (
	"errors"
)

const plusByteValue = 43
const minusByteValue = 45

var errFlipTooMany = errors.New("specified too many pancakes to flip")
var errInvalidString = errors.New("invalid input string. Must only consist of + or - characters.")

// Stack represents a stack of pancakes and has some ease of use methods.
type Stack struct {
	cakes []bool
	flips int
}

// NewStack takes a string of + and - and returns a Stack struct pointer.
func NewStack(in string) (*Stack, error) {
	// restructure as []bool so we can work in place
	s := &Stack{}
	s.cakes = make([]bool, len(in))
	for i, v := range in {
		if v != plusByteValue && v != minusByteValue {
			return nil, errInvalidString
		}
		// byte values for + and - are 43 and 45 respectively when the string is ranged over, but stack's zero vals are already false, so just detect +
		s.cakes[i] = v == 43
	}
	return s, nil
}

// Flip flips the first n pancakes in-place in the slice of bools and increments the internal flip counter.
// n <= 0 is a no-op and does not increment the flip count.
func (s *Stack) Flip(n int) error {
	if n == 0 {
		return nil
	}
	if n > len(s.cakes) {
		return errFlipTooMany
	}
	// flip in place
	for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
		s.cakes[left], s.cakes[right] = !s.cakes[right], !s.cakes[left]
	}
	// the above is a regular reverse, which normally doesn't care about an odd middle, but we need to invert it
	if n%2 == 1 {
		s.cakes[n/2] = !s.cakes[n/2]
	}
	s.flips++
	return nil
}

// IsHappy checks if all the pancakes are happy side up and we are done.
func (s *Stack) IsHappy() bool {
	for _, v := range s.cakes {
		if !v {
			return false
		}
	}
	return true
}

// LowestFlip finds the first from the bottom that needs to be flipped.
// Returns how many to flip to fix it for directly passing into Flip
func (s *Stack) LowestFlip() int {
	for i := len(s.cakes) - 1; i >= 0; i-- {
		if !s.cakes[i] {
			return i + 1 // flip method wants how many to flip, not index. so +1
		}
	}
	return 0
}

// PrepTop returns the number to pre-flip on top to make sure at least the top pancake (or more) is -, so that a deep flip actually makes n'th happy.
// In other words, it returns the number of consecutive +'s from the top to flip.
// Directly compatible with Flip
func (s *Stack) PrepTop() int {
	num := 0
	for _, v := range s.cakes {
		if !v {
			break
		}
		num++
	}
	return num
}

// Equals is just a helper method for checking the state of those yummy pancakes to see if their current state is equal to the passed in slice.
func (s *Stack) Equals(test []bool) bool {
	if len(s.cakes) != len(test) {
		return false
	}
	for i := 0; i < len(s.cakes); i++ {
		if s.cakes[i] != test[i] {
			return false
		}
	}
	return true
}

// EqualsString accepts a string and passes to Equals.
func (s *Stack) EqualsString(test string) (bool, error) {
	ts, err := NewStack(test)
	if err != nil {
		return false, err
	}
	return s.Equals(ts.cakes), nil
}

// Count returns the current number of flips undertaken.
func (s *Stack) Count() int {
	return s.flips
}
