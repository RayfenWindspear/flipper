package flipper

import (
	"errors"
)

const plusByteValue = 43
const minusByteValue = 45

var errFlipTooMany = errors.New("specified too many pancakes to flip")
var errInvalidString = errors.New("invalid input string. Must only consist of + or - characters.")

// stack represents a stack of pancakes and has some ease of use methods
type stack struct {
	cakes []bool
	flips int
}

// NewStack takes a string of + and - and returns a stack struct
func NewStack(in string) (*stack, error) {
	// restructure as []bool so we can work in place
	s := &stack{}
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

// flip flips the first n pancakes in-place in the slice of bools and increments the flip count.
// n <= 0 is a noOp.
func (s *stack) flip(n int) error {
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

// isHappy checks if all are true and we are done.
func (s *stack) isHappy() bool {
	for _, v := range s.cakes {
		if !v {
			return false
		}
	}
	return true
}

// lowestFlip finds the first from the bottom that needs to be flipped.
// Returns how many to flip to fix it for directly passing into stack.flip
func (s *stack) lowestFlip() int {
	for i := len(s.cakes) - 1; i >= 0; i-- {
		if !s.cakes[i] {
			return i + 1 // flip method wants how many to flip, not index. so +1
		}
	}
	return 0
}

// prepTop returns the number to pre-flip on top to make sure at least the top pancake (or more) is -, so that a deep flip actually makes n'th happy.
// In other words, it returns the number of consecutive "+"'s from the top to flip.
// Directly compatible with stack.flip
func (s *stack) prepTop() int {
	num := 0
	for _, v := range s.cakes {
		if !v {
			break
		}
		num++
	}
	return num
}

// equals is just a helper method for checking the state of those yummy pancakes to see if their current state is equal to the passed in slice.
func (s *stack) equals(test []bool) bool {
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
