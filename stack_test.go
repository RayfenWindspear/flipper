package flipper

import (
	"sync"
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

	// bad string
	ok, err = s.EqualsString("+-+-0")
	if ok || err != errInvalidString {
		t.Errorf("EqualsString failed invalid string\n")
	}
}

func TestFlip(t *testing.T) {
	s := &Stack{}
	s.cakes = []bool{false, true, false, true, false}
	count := 0
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}
	err := s.Flip(5)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}
	count = 1
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}

	err = s.Flip(4)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, false, true, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}
	count = 2
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}

	err = s.Flip(3)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}
	count = 3
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}

	err = s.Flip(2)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{false, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}
	count = 4
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}

	err = s.Flip(1)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Equals([]bool{true, true, false, false, true}) {
		t.Errorf("Bad flip check %+v\n", s.cakes)
	}
	count = 5
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
	}

	err = s.Flip(6)
	if err != errFlipTooMany {
		t.Fatal("uncaught overflow error")
	}
	count = 5
	if s.Count() != count {
		t.Errorf("incorrect count %d, should be %d", s.Count(), count)
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

func solveCase(t *testing.T, in string, sol int) {
	s, err := NewStack(in)
	if err != nil {
		t.Fatal(err)
	}
	flips, err := s.Solve()
	if err != nil {
		t.Fatal(err)
	}
	if flips != sol {
		t.Errorf("solution incorrect: %d, expected: %d\n", flips, sol)
	}
}

func TestSolve(t *testing.T) {
	cases := []string{
		"-",
		"-+",
		"+-",
		"+++",
		"--+-",
	}
	solutions := []int{
		1,
		1,
		2,
		0,
		3,
	}
	if len(cases) != len(solutions) {
		t.Errorf("bad Solve cases/solutions")
	}

	for i := range cases {
		solveCase(t, cases[i], solutions[i])
	}
}

func breakStuff(s *Stack, length int, cakes bool, kill chan bool) {
	for {
		copee := make([]bool, length)
		if cakes {
			for i := range copee {
				copee[i] = true
			}
		}
		if length == 4 {
			copee[3] = false
		}
		select {
		case <-kill:
			return
		default:
			s.cakes = copee
		}
	}
}

func TestSolveErrorReturnsHack(t *testing.T) {
	// as written, it's not possible to have Solve return an error... unless we break it using concurrency!
	routines := 4
	wg := &sync.WaitGroup{}

	s, err := NewStack("+++-")
	if err != nil {
		t.Fatal(err)
	}
	kill := make(chan bool)

	// break when PrepTop makes Solve fail.
	// when race condition triggers, PrepTop gets set by breakStuff and returns 5, which is too many.
	// it cannot trigger on LowestFlip, as the race version has no -
	for i := 0; i < routines/2; i++ {
		go breakStuff(s, 5, true, kill)
	}
	for i := 0; i < routines/2; i++ {
		go breakStuff(s, 4, true, kill)
	}
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// if it causes a panic, we gotta try again :(
				TestSolveErrorReturnsHack(t)
				wg.Done()
			}
		}()
		for {
			_, err := s.Solve()
			if err == errFlipTooMany {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < routines; i++ {
		kill <- true
	}

	// break when LowestFlip makes Solve fail.
	// when race condition triggers, LowestFlip will return 5, which is too many.
	// cannot trigger PrepTop, as there are no +
	for i := 0; i < routines; i++ {
		go breakStuff(s, 5, false, kill)
	}
	for i := 0; i < routines/2; i++ {
		go breakStuff(s, 4, true, kill)
	}
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// if it causes a panic, we gotta try again :(
				TestSolveErrorReturnsHack(t)
				wg.Done()
			}
		}()
		for {
			_, err := s.Solve()
			if err == errFlipTooMany {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < routines; i++ {
		kill <- true
	}
}
