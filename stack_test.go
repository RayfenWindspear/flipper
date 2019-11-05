package main

import (
	"testing"
)

func TestStackEquals(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, false, false, false, false, false}
	test := []bool{false, false, false, false, false, false}

	if !s.equals(test) {
		t.Errorf("equals failed valid %+v, %+v\n", s.cakes, test)
	}
	test[5] = true
	if s.equals(test) {
		t.Errorf("equals failed invalid %+v, %+v\n", s.cakes, test)
	}
	if s.equals([]bool{false, false, false, false, false}) {
		t.Errorf("equals failed invalid short size %+v, %+v\n", s.cakes, test)
	}
	if s.equals([]bool{false, false, false, false, false, false, false}) {
		t.Errorf("equals failed invalid long size %+v, %+v\n", s.cakes, test)
	}
	s.cakes = []bool{}
	if !s.equals([]bool{}) {
		t.Errorf("equals failed 0 length %+v, %+v\n", s.cakes, test)
	}

	for l := 0; l <= 10; l++ {
		s.cakes = make([]bool, l)
		// test a bunch of mutations... just cuz
		for i := 0; i < l; i++ {
			test = make([]bool, l)
			s.cakes[i] = true
			if s.equals(test) {
				t.Errorf("equals failed. should be false %+v, %+v\n", s.cakes, test)
			}
			for j := 0; j < l; j++ {
				test[j] = true
				check := s.equals(test)
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

func TestFlip(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, true, false, true, false}
	err := s.flip(5)
	if err != nil {
		t.Fatal(err)
	}
	if !s.equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(4)
	if err != nil {
		t.Fatal(err)
	}
	if !s.equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(3)
	if err != nil {
		t.Fatal(err)
	}
	if !s.equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(2)
	if err != nil {
		t.Fatal(err)
	}
	if !s.equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(1)
	if err != nil {
		t.Fatal(err)
	}
	if !s.equals([]bool{true, true, false, false, true}) {
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
	s.cakes = []bool{false, false, false, false, false}
	if s.isHappy() {
		t.Errorf("not a valid win condition")
	}
	s.cakes = []bool{true, true, true, true, true}
	if !s.isHappy() {
		t.Errorf("valid condition not detected")
	}
}

func TestLowestFlip(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, true, false, true, false}
	n := s.lowestFlip()
	b := 5
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, false, false, false, true}
	n = s.lowestFlip()
	b = 4
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, false, true, true}
	n = s.lowestFlip()
	b = 3
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, false, true, true, true}
	n = s.lowestFlip()
	b = 2
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, true, true, true}
	n = s.lowestFlip()
	b = 1
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, true}
	n = s.lowestFlip()
	b = 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}
}

func TestPrepTop(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, false, false, false, false}
	n := s.prepTop()
	b := 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{false, true, true, true, true}
	n = s.prepTop()
	b = 0
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, false, true, true, true}
	n = s.prepTop()
	b = 1
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, false, true, true}
	n = s.prepTop()
	b = 2
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, false, true}
	n = s.prepTop()
	b = 3
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, false}
	n = s.prepTop()
	b = 4
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}

	s.cakes = []bool{true, true, true, true, true}
	n = s.prepTop()
	b = 5
	if n != b {
		t.Errorf("%d should be %d", n, b)
	}
}
