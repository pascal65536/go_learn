package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}
	return utf8.RuneCount(input), nil
}

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input       []byte
		wantLength  int
		wantErr     error
		description string
	}{
		{[]byte("hello"), 5, nil, "ASCII string"},
		{[]byte("привет"), 6, nil, "UTF-8 Cyrillic"},
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8, "Invalid UTF-8 bytes"},
		{[]byte(""), 0, nil, "Empty string"},
		{[]byte("😊"), 1, nil, "Emoji character"},
	}

	for i, test := range tests {
		gotLength, gotErr := GetUTFLength(test.input)
		if gotLength != test.wantLength || gotErr != test.wantErr {
			t.Errorf("#%d (%s): GetUTFLength(%v) = (%d, %v); want (%d, %v)",
				i, test.description, test.input, gotLength, gotErr, test.wantLength, test.wantErr)
		}
	}
}




type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},  // a < 0
	{0, "zero"},       // a == 0
	{5, "short"},      // 0 < a < 10
	{50, "long"},      // 10 <= a < 100
	{150, "very long"},// a >= 100
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Length(%d) = %s; want %s", i, test.in, size, test.out)
		}
	}
}


func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; ожидается %d", result, expected)
	}
}




func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 4, -8},
		{-3, -3, 9},
		{7, 1, 7},
	}

	for i, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("#%d: Multiply(%d, %d) = %d; want %d", i, test.a, test.b, result, test.expected)
		}
	}
}




func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hll wrld"},
		{"AEIOUaeiou", ""},
		{"sky", "sky"},
		{"Go is fun!", "G s fn!"},
		{"", ""},
		{"rhythm", "rhythm"},
	}

	for i, test := range tests {
		result := DeleteVowels(test.input)
		if result != test.expected {
			t.Errorf("#%d: DeleteVowels(%q) = %q; want %q", i, test.input, result, test.expected)
		}
	}
}
