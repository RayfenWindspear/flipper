package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// competition io template (returns added) from:
// https://www.codementor.io/tucnak/using-golang-for-competitive-programming-h8lhvxzt3
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) (n int, err error) { return fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) (n int, err error)  { return fmt.Fscanf(reader, f, a...) }

// end template

var errFlipTooMany = errors.New("specified too many pancakes to flip")

// stack represents a stack of pancakes and has some ease of use methods
type stack struct {
	cakes []bool
	flips int
}

// flip flips the first n pancakes
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

// isHappy checks if all are true and we are done
func (s *stack) isHappy() bool {
	for _, v := range s.cakes {
		if !v {
			return false
		}
	}
	return true
}

// lowestFlip finds the first from the bottom that needs to be flipped. Returns how many to flip to fix it
func (s *stack) lowestFlip() int {
	for i := len(s.cakes) - 1; i >= 0; i-- {
		if !s.cakes[i] {
			return i + 1 // flip method wants how many to flip, not index. so +1
		}
	}
	return 0
}

// prepTop returns the number to pre-flip on top to make sure at least the top pancake (or more) is -, so that a deep flip actually makes n'th happy
func (s *stack) prepTop() int {
	num := 0
	for _, v := range s.cakes {
		if v {
			num++
		} else {
			break
		}
	}
	return num
}

// flipper is the solver for the pancake problem. Handles the input and output and solves.
type flipper struct {
	length  int
	problem []string
	current int
}

// readProblem reads the problem from os.Stdin and stores it in the struct members.
func (f *flipper) readProblem() error {
	_, err := scanf("%d\n", &f.length)
	if err != nil {
		return err
	}
	f.problem = make([]string, f.length)
	for i := 0; i < f.length; i++ {
		var line string
		_, err = scanf("%s\n", &line)
		if err != nil {
			return err
		}
		f.problem[i] = line
	}
	return nil
}

// solveNext preps and solves a single input line, the one indicated by the index 'current', and prints the solution to os.Stdout.
// It solves by
func (f *flipper) solveNext() error {
	// restructure as []bool so we can work in place
	s := &stack{}
	s.cakes = make([]bool, len(f.problem[f.current]))
	for i, v := range f.problem[f.current] {
		// byte values for + and - are 43 and 45 respectively when the string is ranged over, but stack's zero vals are already false, so just detect +
		if v == 43 {
			s.cakes[i] = true
		}
	}
	for !s.isHappy() {
		err := s.flip(s.prepTop())
		if err != nil {
			return err
		}
		if s.isHappy() {
			break
		}
		err = s.flip(s.lowestFlip())
		if err != nil {
			return err
		}
	}
	f.current++
	printf("Case #%d: %d\n", f.current, s.flips)
	return nil
}

// solveAll just iterates and solves all the problems in the set.
func (f *flipper) solveAll() error {
	for i := 0; i < f.length; i++ {
		err := f.solveNext()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	defer writer.Flush()

	f := &flipper{}

	err := f.readProblem()
	if err != nil {
		panic(err)
	}
	err = f.solveAll()
	if err != nil {
		panic(err)
	}
}
