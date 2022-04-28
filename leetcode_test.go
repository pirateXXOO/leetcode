package main

import (
	"testing"
)

func TestIfEuqal(t *testing.T) {
	if !IfEqual([]int{1, 2, 3}, []int{3, 2, 1}) {
		t.Errorf("TestFailed")
	}
	if !IfEqual([]int{1, -1, 0}, []int{1, 0, -1}) {
		t.Errorf("TestFailed")
	}
	if !IfEqual([]int{-1, 2, -1}, []int{2, -1, -1}) {
		t.Errorf("TestFailed")
	}
	if !IfEqual([]int{0, -3, 3}, []int{-3, 3, 0}) {
		t.Errorf("TestFailed")
	}
	if !IfEqual([]int{-1, -2, 3}, []int{-1, -2, 3}) {
		t.Errorf("TestFailed")
	}
}

// func TestThreeSum(t *testing.T) {
// 	ans := ThreeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
// 	should := [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-1, -2, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}
// 	// res := [][]int{{-1, 0, 1}, {-1, 2, -1}}
// 	if !IfEqual(ans, should) {
// 		t.Error("Failed")
// 	}

// }
