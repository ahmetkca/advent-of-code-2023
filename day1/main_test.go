package main

import "testing"

func TestIsDigit(t *testing.T) {
	ch := '1'
	digit, isDigit := IsDigit(ch)

	want1 := true
	if !isDigit {
		t.Errorf("IsDigit(%q) err = %q, want %v", ch, "Expected isDigit was different than teh returned value", want1)
	}

	want2 := 1
	if isDigit && digit != 1 {
		t.Errorf("IsDigit(%q) err = %q, want %q", ch, "Expected digit was different than teh returned value", want2)
	}

	ch = 'a'
	digit, isDigit = IsDigit(ch)

	want1 = false
	if isDigit {
		t.Errorf("IsDigit(%q) err = %q, want %v", ch, "Expected isDigit was different than teh returned value", want1)
	}

	want2 = -1
	if isDigit && digit != -1 {
		t.Errorf("IsDigit(%q) err = %q, want %q", ch, "Expected digit was different than teh returned value", want2)
	}
}

func TestGetFirstAndLastDigitPartOne(t *testing.T) {
	input := "xsdm7lhcjzk3hstcf"

	want := []int{7, 3}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigit(%q) got = %q, want %q", input, got, want)
	}
}




func TestGetFirstAndLastDigitPartTwo(t *testing.T) {

	input := "rgbfivefive3eightthree"

	want := []int{5, 3}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo2(t *testing.T) {

	input := "rg1bfivefive3eightthree"

	want := []int{1, 3}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo3(t *testing.T) {

	input := "rg1bfiavefi4vae3eiaghtthr5aee"

	want := []int{1, 5}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo4(t *testing.T) {

	input := "9fourcsjph86shfqjrxlfourninev"

	want := []int{9, 9}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo5(t *testing.T) {

	input := "ffoaurcsjphh6shfqjrxlfoufrnifnev"

	want := []int{6, 6}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo6(t *testing.T) {

	input := "xbfqgxkxmninegmhcrcmxgllktllbqpsvrfthree9"

	want := []int{9, 9}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo7(t *testing.T) {

	input := "ggqvtwofrrfdhbsdfdpzpbjjckbrg5ninebqszf"

	want := []int{2, 9}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo8(t *testing.T) {

	inputs := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	wants := []struct {
		f int
		s int
	}{
		{2, 9},
		{8, 3},
		{1, 3},
		{2, 4},
		{4, 2},
		{1, 4},
		{7, 6},
	}

	for i, input := range inputs {
		want := wants[i]

		first, last := GetFirstAndLast([]byte(input))

		if first != want.f || last != want.s {
			t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %d, %d, want = %d, %d", input, first, last, want.f, want.s)
		}
	}
}

func TestGetFirstAndLastDigitPartTwo9(t *testing.T) {

	input := "eightwothree"

	want := []int{8, 3}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo10(t *testing.T) {

	input := "seven8onertbqhthreefourctdbsrcvcsjlvcxneight"

	want := []int{7, 8}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo11(t *testing.T) {

	input := "dkxscqdrctgxzlflrlvtnkqlmrlgsrqseven8qlqrdz"

	want := []int{7, 8}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}
}

func TestGetFirstAndLastDigitPartTwo12(t *testing.T) {

	input := "3nine4fourjclspd152rpv"

	want := []int{3, 2}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo13(t *testing.T) {

	input := "pxxtd793frjfckbhstcdhsrx58mksktwoneqx"

	want := []int{7, 1}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo14(t *testing.T) {

	input := "eighthreeightwone5"

	want := []int{8, 5}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo15(t *testing.T) {

	input := "xgjjmnlvznf2nineltmsevenine"

	want := []int{2, 9}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo16(t *testing.T) {

	input := "twoneighthreeightwoneight"

	want := []int{2, 8}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo18(t *testing.T) {

	inputs := []string{
		"1twone",
		"one2one",
		"one12one",
	}

	wants := []struct {
		f int
		s int
	}{
		{1, 1},
		{1, 1},
		{1, 1},
	}

	sum := 0
	expectedSum := 33

	for i, input := range inputs {
		want := wants[i]

		first, last := GetFirstAndLast([]byte(input))

		if first != want.f || last != want.s {
			t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %d, %d, want = %d, %d", input, first, last, want.f, want.s)
		}

		sum += ((10 * first) + last)
	}

	if sum != 33 {
		t.Errorf("Expected sum = %d, got = %d", expectedSum, sum)
	}
}

func TestGetFirstAndLastDigitPartTwo45(t *testing.T) {

	input := "mkfone4ninefour"

	want := []int{1, 4}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}

func TestGetFirstAndLastDigitPartTwo46(t *testing.T) {

	input := "oneighthreeightwone2"

	want := []int{1, 2}
	got := []int{}

	first, last := GetFirstAndLast([]byte(input))

	got = append(got, first)
	got = append(got, last)

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetFirstAndLastDigitPartTwo(%q) got = %q, want %q", input, got, want)
	}

	t.Logf("got = %d, %d\n", first, last)
}
