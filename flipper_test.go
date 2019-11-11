package flipper

import (
	"bytes"
	"io/ioutil"
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

	err := f.ReadProblem()
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

	err := f.ReadProblem()
	if err != nil {
		t.Fatal(err)
	}
	err = f.SolveNext()
	if err != nil {
		t.Fatal(err)
	}
	err = f.SolveNext()
	if err != nil {
		t.Fatal(err)
	}
	err = f.Flush()
	if err != nil {
		t.Fatal(err)
	}

	if string(out.Bytes()) != "Case #1: 0\nCase #2: 2\n" {
		t.Errorf("Incorrect results:\n%s", string(out.Bytes()))
	}
}

func TestSolveAll(t *testing.T) {
	out := bytes.NewBuffer(nil)
	f := NewFlipper(bytes.NewBuffer([]byte(data)), out)

	err := f.ReadProblem()
	if err != nil {
		t.Fatal(err)
	}
	err = f.SolveAll()
	if err != nil {
		t.Fatal(err)
	}

	if string(out.Bytes()) != output {
		t.Errorf("Incorrect results:\n%s", string(out.Bytes()))
	}
}

func mockStdin(data []byte) (func(), error) {
	tmpfile, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.Write(data); err != nil {
		return nil, err
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	oldStdin := os.Stdin
	os.Stdin = tmpfile
	ranDefer := false
	return func() {
		if !ranDefer {
			tmpfile.Close()
			os.Remove(tmpfile.Name()) // clean up
			os.Stdin = oldStdin       // Restore original Stdin
		}
	}, nil
}

func mockStdout() (*os.File, func(), error) {
	tmpfile2, err := ioutil.TempFile("", "tmpfile2")
	if err != nil {
		return nil, nil, err
	}

	oldStdout := os.Stdout
	os.Stdout = tmpfile2
	ranDefer := false
	return tmpfile2, func() {
		if !ranDefer {
			os.Remove(tmpfile2.Name()) // clean up
			os.Stdout = oldStdout      // Restore original Stdin
			ranDefer = true
		}
	}, nil
}

func TestDoEverything(t *testing.T) {
	// mock stdin
	inDefer, err := mockStdin([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	defer inDefer()

	// mock stdout
	outfile, outDefer, err := mockStdout()
	if err != nil {
		t.Fatal(err)
	}
	defer outDefer()

	err = DoEverything()
	if err != nil {
		t.Fatal(err)
	}

	// read output
	if _, err := outfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	dat, err := ioutil.ReadAll(outfile)
	if err != nil {
		t.Fatal(err)
	}
	if output != string(dat) {
		t.Errorf("output doesn't match:\n%s\n\nExpected:\n%s", string(dat), output)
	}
}

func TestMiscErrs(t *testing.T) {
	// test err return reading length
	// mock stdin
	inDefer, err := mockStdin([]byte("a"))
	if err != nil {
		t.Fatal(err)
	}
	defer inDefer()
	f := newFlipper()
	err = f.ReadProblem()
	if err == nil {
		t.Errorf("ReadProblem should have failed")
	}
	inDefer() // explicitly clean up

	// test err return reading line
	inDefer, err = mockStdin([]byte("2\n"))
	if err != nil {
		t.Fatal(err)
	}
	defer inDefer()
	f = newFlipper()
	err = f.ReadProblem()
	if err == nil {
		t.Errorf("ReadProblem should have failed")
	}
	inDefer() // explicitly clean up

	// test SolveNext NewStack err return
	f.problem = []string{"a"}
	err = f.SolveNext()
	if err != errInvalidString {
		t.Errorf("SolveNext NewStack should have failed")
	}

	err = f.SolveAll()
	if err != errInvalidString {
		t.Errorf("SolveAll should have failed")
	}

	// test DoEverything Read fail
	inDefer, err = mockStdin([]byte("3\n"))
	if err != nil {
		t.Fatal(err)
	}
	defer inDefer()
	err = DoEverything()
	if err == nil {
		t.Errorf("DoEverything should have failed reading")
	}
	inDefer() // explicitly clean up

	// test DoEverything Solve fail
	inDefer, err = mockStdin([]byte("1\nabcd\n"))
	if err != nil {
		t.Fatal(err)
	}
	defer inDefer()
	err = DoEverything()
	if err == nil {
		t.Errorf("DoEverything should have failed reading")
	}
	inDefer() // explicitly clean up
}
