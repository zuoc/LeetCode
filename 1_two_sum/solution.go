// 查表法~
package _1_two_sum

func twoSum(nums []int, target int) []int {
	tb := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if j, ok := tb[target-nums[i]]; ok {
			return []int{j, i}
		}
		tb[nums[i]] = i
	}
	return []int{}
}
