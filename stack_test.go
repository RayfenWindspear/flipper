package main

import (
	"testing"
)

func TestFlip(t *testing.T) {
	s := &stack{}
	s.cakes = []bool{false, true, false, true, false}
	err := s.flip(5)
	if err != nil {
		t.Fatal(err)
	}
	if !s.stackEquals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(4)
	if err != nil {
		t.Fatal(err)
	}
	if !s.stackEquals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(3)
	if err != nil {
		t.Fatal(err)
	}
	if !s.stackEquals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(2)
	if err != nil {
		t.Fatal(err)
	}
	if !s.stackEquals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}

	err = s.flip(1)
	if err != nil {
		t.Fatal(err)
	}
	if !s.stackEquals([]bool{true, true, false, false, true}) {
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
