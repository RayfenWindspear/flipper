package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

const data = `8
+++
++-
+-+
+--
-++
-+-
--+
---`

func TestReadProblem(t *testing.T) {
	f := &flipper{}

	reader = bufio.NewReader(bytes.NewBuffer([]byte(data)))

	err := f.readProblem()
	// restore reader to stdin
	reader = bufio.NewReader(os.Stdin)
	if err != nil {
		t.Fatal(err)
	}

	if f.length != 8 {
		t.Errorf("Invalid length %d\n", f.length)
	}

	if len(f.problem) != 8 {
		t.Errorf("Line reads incorrect %d\n", len(f.problem))
	}

	datalines := strings.Split(data, "\n")
	for i := range f.problem {
		if f.problem[i] != datalines[i+1] {
			t.Errorf("Lines read don't match up. %s : %s", f.problem[i], datalines[i+1])
		}
	}
}

func TestSolveNext(t *testing.T) {
	f := &flipper{}

	reader = bufio.NewReader(bytes.NewBuffer([]byte(data)))
	err := f.readProblem()
	// restore reader to stdin
	reader = bufio.NewReader(os.Stdin)
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	writer = bufio.NewWriter(buf)
	err = f.solveNext()
	if err != nil {
		t.Fatal(err)
	}
	err = f.solveNext()
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()

	// restore writer to stdout
	writer = bufio.NewWriter(os.Stdout)

	str := string(buf.Bytes())

	if str != "Case #1: 0\nCase #2: 2\n" {
		t.Errorf("Incorrect results:\n%s", str)
	}
}

func TestSolveAll(t *testing.T) {
	f := &flipper{}

	reader = bufio.NewReader(bytes.NewBuffer([]byte(data)))
	err := f.readProblem()
	// restore reader to stdin
	reader = bufio.NewReader(os.Stdin)
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	writer = bufio.NewWriter(buf)
	err = f.solveAll()
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()

	// restore writer to stdout
	writer = bufio.NewWriter(os.Stdout)

	str := string(buf.Bytes())

	if str != `Case #1: 0
Case #2: 2
Case #3: 2
Case #4: 2
Case #5: 1
Case #6: 3
Case #7: 1
Case #8: 1
` {
		t.Errorf("Incorrect results:\n%s", str)
	}
}

func stackEquals(s, test []bool) bool {
	if len(s) != len(test) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != test[i] {
			return false
		}
	}
	return true
}

func TestFlip(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, true, false, true, false}
	err := s.flip(5)
	if err != nil {
		t.Fatal(err)
	}
	if !stackEquals(s.cakes, []bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(4)
	if err != nil {
		t.Fatal(err)
	}
	if !stackEquals(s.cakes, []bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(3)
	if err != nil {
		t.Fatal(err)
	}
	if !stackEquals(s.cakes, []bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(2)
	if err != nil {
		t.Fatal(err)
	}
	if !stackEquals(s.cakes, []bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(1)
	if err != nil {
		t.Fatal(err)
	}
	if !stackEquals(s.cakes, []bool{true, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(6)
	if err != errFlipTooMany {
		t.Fatal("uncaught overflow error")
	}
}

func TestIsHappy(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, true, false, true, false}
	if s.isHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, false, true, false}
	if s.isHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, true, true, false}
	if s.isHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, true, true, true}
	if !s.isHappy() {
		t.Errorf("valid condition not detected")
	}
}
