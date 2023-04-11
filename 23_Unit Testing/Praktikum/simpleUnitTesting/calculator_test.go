package simpleunittesting

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 3)
	want := 2

	if got != want {
		t.Errorf("Subtract(5, 3) = %d; want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	got := Divide(10, 2)
	want := 5.0

	if got != want {
		t.Errorf("Divide(10, 2) = %f; want %f", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(2, 3)
	want := 6

	if got != want {
		t.Errorf("Multiply(2, 3) = %d; want %d", got, want)
	}
}
