package flipper

import (
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

const output = `Case #1: 0
Case #2: 2
Case #3: 2
Case #4: 2
Case #5: 1
Case #6: 3
Case #7: 1
Case #8: 1
`

func TestReadProblem(t *testing.T) {
	f := NewFlipper(bytes.NewBuffer([]byte(data)), os.Stdout)

	err := f.readProblem()
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
	out := bytes.NewBuffer(nil)
	f := NewFlipper(bytes.NewBuffer([]byte(data)), out)

	err := f.readProblem()
	if err != nil {
		t.Fatal(err)
	}
	err = f.solveNext()
	if err != nil {
		t.Fatal(err)
	}
	err = f.solveNext()
	if err != nil {
		t.Fatal(err)
	}
	f.writer.Flush()

	if string(out.Bytes()) != "Case #1: 0\nCase #2: 2\n" {
		t.Errorf("Incorrect results:\n%s", string(out.Bytes()))
	}
}

func TestSolveAll(t *testing.T) {
	out := bytes.NewBuffer(nil)
	f := NewFlipper(bytes.NewBuffer([]byte(data)), out)

	err := f.readProblem()
	if err != nil {
		t.Fatal(err)
	}
	err = f.solveAll()
	if err != nil {
		t.Fatal(err)
	}

	if string(out.Bytes()) != output {
		t.Errorf("Incorrect results:\n%s", string(out.Bytes()))
	}
}
