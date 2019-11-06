package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// flipper is the solver for the pancake problem. Handles the input and output and solves.
type flipper struct {
	length  int
	problem []string
	current int
	reader  *bufio.Reader
	writer  *bufio.Writer
}

// NewFlipper creates a new flipper struct with buffered io using the passed in input and output sinks.
func NewFlipper(in io.Reader, out io.Writer) *flipper {
	return &flipper{
		length:  0,
		problem: nil,
		current: 0,
		reader:  bufio.NewReader(in),
		writer:  bufio.NewWriter(out),
	}
}

// newFlipper internal ease of use function just calls NewFlipper using os.Stdin and os.Stdout.
func newFlipper() *flipper {
	return NewFlipper(os.Stdin, os.Stdout)
}

// readProblem reads the problem from flipper.reader and stores it in the struct members.
func (f *flipper) readProblem() error {
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

// solveNext preps and solves a single input line, the one indicated by the index 'current', and prints the solution to os.Stdout.
// It solves by iteratively padding the top with as many "-"" as it can with 0-1 flip, then flipping from the bottommost "-"
func (f *flipper) solveNext() error {
	s, err := NewStack(f.problem[f.current])
	if err != nil {
		return err
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
	fmt.Fprintf(f.writer, "Case #%d: %d\n", f.current, s.flips)
	return nil
}

// solveAll just iterates and solves all the problems in the set.
func (f *flipper) solveAll() error {
	defer f.writer.Flush()
	for i := 0; i < f.length; i++ {
		err := f.solveNext()
		if err != nil {
			return err
		}
	}
	return nil
}

// DoEverything is just an exported function that creates a default flipper, reads the input, solves, and outputs.
func DoEverything() error {
	f := newFlipper()

	err := f.readProblem()
	if err != nil {
		return err
	}
	err = f.solveAll()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := DoEverything()
	if err != nil {
		panic(err)
	}
}
