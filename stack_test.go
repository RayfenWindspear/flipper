package flipper

import (
	"testing"
)

func TestStackEquals(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, false, false, false, false, false}
	test := []bool{false, false, false, false, false, false}

	if !s.Equals(test) {
		t.Errorf("equals failed valid %+v, %+v\n", s.cakes, test)
	}
	test[5] = true
	if s.Equals(test) {
		t.Errorf("equals failed invalid %+v, %+v\n", s.cakes, test)
	}
	if s.Equals([]bool{false, false, false, false, false}) {
		t.Errorf("equals failed invalid short size %+v, %+v\n", s.cakes, test)
	}
	if s.Equals([]bool{false, false, false, false, false, false, false}) {
		t.Errorf("equals failed invalid long size %+v, %+v\n", s.cakes, test)
	}
	s.cakes = []bool{}
	if !s.Equals([]bool{}) {
		t.Errorf("equals failed 0 length %+v, %+v\n", s.cakes, test)
	}

	for l := 0; l <= 10; l++ {
		s.cakes = make([]bool, l)
		// test a bunch of mutations... just cuz
		for i := 0; i < l; i++ {
			test = make([]bool, l)
			s.cakes[i] = true
			if s.Equals(test) {
				t.Errorf("equals failed. should be false %+v, %+v\n", s.cakes, test)
			}
			for j := 0; j < l; j++ {
				test[j] = true
				check := s.Equals(test)
				if i == j {
					// test should succeed iff i == j
					if !check {
						t.Errorf("equals failed. should be true %+v, %+v\n", s.cakes, test)
					}
				} else {
					if check {
						t.Errorf("equals failed. should be false %+v, %+v\n", s.cakes, test)
					}
				}
			}
		}
	}
}

func TestNewStack(t *testing.T) {
	s, err := NewStack("++++")
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, true, true, true}) {
		t.Errorf("NewStack failed %+v\n", s.cakes)
	}
	s, err = NewStack("+-+-")
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, false, true, false}) {
		t.Errorf("NewStack failed %+v\n", s.cakes)
	}
	s, err = NewStack("----")
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, false, false, false}) {
		t.Errorf("NewStack failed %+v\n", s.cakes)
	}
	s, err = NewStack("----+")
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, false, false, false, true}) {
		t.Errorf("NewStack failed %+v\n", s.cakes)
	}

	s, err = NewStack("----+a")
	if err != errInvalidString {
		t.Errorf("NewStack failed invalid string")
	}
}

func testStackEqualsStringCases(t *testing.T, s *Stack, cases []string) {
	for _, test := range cases {
		ok, err := s.EqualsString(test)
		if err != nil {
			t.Fatal(err)
		}
		if ok {
			t.Errorf("EqualsString failed: %s, %+v\n", test, s.cakes)
		}
	}
}

func TestStackEqualsString(t *testing.T) {
	s, err := NewStack("++++")
	if err != nil {
		t.Fatal(err)
	}
	test := "++++"
	ok, err := s.EqualsString(test)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("EqualsString failed: %s, %+v\n", test, s.cakes)
	}
	cases := []string{
		"+++-",
		"----",
		"+-+-",
		"-+-+",
		"+---",
	}
	testStackEqualsStringCases(t, s, cases)

	test = "+++++"
	ok, err = s.EqualsString(test)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Errorf("EqualsString failed size diff: %s, %+v\n", test, s.cakes)
	}

	s, err = NewStack("-----")
	if err != nil {
		t.Fatal(err)
	}
	cases = []string{
		"++++-",
		"----+",
		"+-+-+",
		"-+-+-",
		"+----",
	}
	testStackEqualsStringCases(t, s, cases)
}

func TestFlip(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, true, false, true, false}
	err := s.Flip(5)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.Flip(4)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.Flip(3)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.Flip(2)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.Flip(1)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.Flip(6)
	if err != errFlipTooMany {
		t.Fatal("uncaught overflow error")
	}
}

func TestIsHappy(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, true, false, true, false}
	if s.IsHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, false, true, false}
	if s.IsHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, true, true, false}
	if s.IsHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{false, false, false, false, false}
	if s.IsHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, true, true, true}
	if !s.IsHappy() {
		t.Errorf("valid condition not detected")
	}
}

func TestLowestFlip(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, true, false, true, false}
	n := s.LowestFlip()
	b := 5
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, false, false, false, true}
	n = s.LowestFlip()
	b = 4
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, false, true, true}
	n = s.LowestFlip()
	b = 3
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, false, true, true, true}
	n = s.LowestFlip()
	b = 2
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, true, true, true}
	n = s.LowestFlip()
	b = 1
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, true}
	n = s.LowestFlip()
	b = 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}
}

func TestPrepTop(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, false, false, false, false}
	n := s.PrepTop()
	b := 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, true, true, true}
	n = s.PrepTop()
	b = 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, false, true, true, true}
	n = s.PrepTop()
	b = 1
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, false, true, true}
	n = s.PrepTop()
	b = 2
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, false, true}
	n = s.PrepTop()
	b = 3
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, false}
	n = s.PrepTop()
	b = 4
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, true}
	n = s.PrepTop()
	b = 5
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}
}
