package _1_two_sum

import "testing"

func Test_twoSum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	if result[0] == 0 || result[1] == 1 {
		t.Logf("success")
	} else {
		t.Errorf("failed")
	}
}

func Test_twoSum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)

	if result[0] == 1 || result[1] == 2 {
		t.Logf("success")
	} else {
		t.Errorf("failed")
	}
}
