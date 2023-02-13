package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 4.50, b: 0.50, want: 4},
		{a: 5, b: 0, want: 5},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply (%f, %f): want %f, got %f", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiplyFail(t *testing.T) {
	t.Parallel()
	want := 14.62
	a := 4.3
	b := 3.4
	got := calculator.Multiply(4.3, 3.4)

	if want != got {
		t.Errorf("Multiply(%f, %f): want %f, got %f", a, b, want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 1.0000, b: 3.0000, want: 0.333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			if !closeEnough(tc.want, got, 0.001) {
				t.Errorf("Divide(%f, %f): want %f, got %f",
					tc.a, tc.b, tc.want, got)
			}
		}
	}
}

func TestDivideBy0(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("was expecting error for invalid input, but got nothing")
	}
}

// Helper function
func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a,
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 16, want: 4},
		{a: 10, want: 3.16227},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			if !closeEnough(tc.want, got, 0.00001) {
				t.Errorf("Sqrt(%f): want %f, got %f",
					tc.a, tc.want, got)
			}
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-64)
	if err == nil {
		t.Error("was expecting error for invalid input, but got nothing")
	}
}
