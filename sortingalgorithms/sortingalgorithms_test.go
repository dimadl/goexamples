package sortingalgorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	got := BubbleSort(numbers)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("error")
	}
}

func TestMergeSort(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	got := MergeSort(numbers)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("error %v", got)
	}
}

func BenchmarkBubbleSort(b *testing.B) {

	var arr []int
	for i := 100; i > 0; i-- {
		arr = append(arr, i)
	}

	b.ResetTimer()

	BubbleSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {

	var arr []int
	for i := 100; i > 0; i-- {
		arr = append(arr, i)
	}

	b.ResetTimer()

	MergeSort(arr)
}
