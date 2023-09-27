package algorithms

import (
	"reflect"
	"testing"
)

func cpSlice[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}

func assertEqual(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Exppected: %v, got: %v", expected, actual)
	}
}

func TestSorted(t *testing.T) {
	original := []int{1, 2}
	sorted := cpSlice(original)
	BubbleSort(sorted)
	assertEqual(t, original, sorted)
}

func TestUnsorted(t *testing.T) {
	original := []int{1, 4, 3, 2}
	expected := []int{1, 2, 3, 4}
	BubbleSort(original)
	assertEqual(t, expected, original)
}
