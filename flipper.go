// Package flipper contains my personal solution to the 'Revenge of the Pancakes' coding challenge.
package flipper

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Flipper is the solver for the pancake problem. Handles the input and output and solves.
type Flipper struct {
	length  int
	problem []string
	current int
	reader  *bufio.Reader
	writer  *bufio.Writer
}

// NewFlipper creates a new flipper struct with buffered io using the passed in input and output sinks.
func NewFlipper(in io.Reader, out io.Writer) *Flipper {
	return &Flipper{
		length:  0,
		problem: nil,
		current: 0,
		reader:  bufio.NewReader(in),
		writer:  bufio.NewWriter(out),
	}
}

// newFlipper internal ease of use function just calls NewFlipper using os.Stdin and os.Stdout.
func newFlipper() *Flipper {
	return NewFlipper(os.Stdin, os.Stdout)
}

// ReadProblem reads the problem from the io.Reader input and stores it in the struct members.
func (f *Flipper) ReadProblem() error {
	_, err := fmt.Fscanf(f.reader, "%d\n", &f.length)
	if err != nil {
		return err
	}
	f.problem = make([]string, f.length)
	for i := 0; i < f.length; i++ {
		var line string
		_, err = fmt.Fscanf(f.reader, "%s\n", &line)
		if err != nil {
			return err
		}
		f.problem[i] = line
	}
	return nil
}

// SolveNext preps and solves a single input line and prints the solution to the io.Writer output buffer.
func (f *Flipper) SolveNext() error {
	// It solves by iteratively padding the top with as many - as it can with 0-1 flips, then flipping from the bottommost -.
	// Internal comment for solution just so it doesn't show up in godoc.
	s, err := NewStack(f.problem[f.current])
	if err != nil {
		return err
	}
	for !s.IsHappy() {
		err := s.Flip(s.PrepTop())
		if err != nil {
			return err
		}
		if s.IsHappy() {
			break
		}
		err = s.Flip(s.LowestFlip())
		if err != nil {
			return err
		}
	}
	f.current++
	fmt.Fprintf(f.writer, "Case #%d: %d\n", f.current, s.flips)
	return nil
}

// SolveAll just iterates and solves all the problems in the set.
func (f *Flipper) SolveAll() error {
	defer f.writer.Flush()
	for i := 0; i < f.length; i++ {
		err := f.SolveNext()
		if err != nil {
			return err
		}
	}
	return nil
}

// Flush just calls the internal wrtier buffer's Flush method.
func (f *Flipper) Flush() error {
	return f.writer.Flush()
}

// DoEverything is just an exported function that creates a default flipper, reads the input, solves, and outputs.
func DoEverything() error {
	f := newFlipper()

	err := f.ReadProblem()
	if err != nil {
		return err
	}
	err = f.SolveAll()
	if err != nil {
		return err
	}
	return nil
}
