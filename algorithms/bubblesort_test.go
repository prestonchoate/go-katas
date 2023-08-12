package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{1, 4, 23, 65, 213, 4, 0, 5}
	sorted := BubbleSort(nums)

	expected := []int{0, 1, 4, 4, 5, 23, 65, 213}

	if !reflect.DeepEqual(expected, sorted) {
		t.Fatalf("expected %q, got %q", expected, sorted)
	}
	t.Log("Bubble sort successful")
}
