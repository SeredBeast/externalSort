package test

import (
	"fmt"
	"testing"
)

func TestStringComparing(t *testing.T) {
	s1 := "abcd"
	s2 := "abce"
	res := s1 < s2
	fmt.Println(res)
}

func TestMergeSort(t *testing.T) {
	x1 := []int{1, 3, 5, 7, 109}
	x2 := []int{102, 104, 106, 108, 1010, 1012, 1014, 1016, 1018}

	merged := mergeSortedArrays(x1, x2)
	fmt.Println(merged)
}

func mergeSortedArrays(x1 []int, x2 []int) []int {
	result := make([]int, len(x1)+len(x2))
	i, j, k := 0, 0, 0

	for i < len(x1) && j < len(x2) {
		if x1[i] <= x2[j] {
			result[k] = x1[i]
			i++
		} else {
			result[k] = x2[j]
			j++
		}
		k++
	}

	for i < len(x1) {
		result[k] = x1[i]
		i++
		k++
	}

	for j < len(x2) {
		result[k] = x2[j]
		j++
		k++
	}

	return result
}

func drop(arr []string) []string {
	return arr[:len(arr)-1]
}

func TestDrop(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	fmt.Println(arr)
	arr = drop(arr)
	fmt.Println(arr)

}
