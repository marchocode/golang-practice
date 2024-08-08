package datastruct

// https://leetcode.cn/leetbook/read/array-and-string/cxqdh/
// 使用二分法搜索插入位置
func SearchInsert(nums []int, target int) int {

	var first, second, middle int = 0, len(nums) - 1, 0

	for first <= second {

		middle = (first + second) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			second = middle - 1
		} else {
			first = middle + 1
		}

	}

	return first

}

// https://leetcode.cn/problems/two-sum/
// 求两个数之和
// 简答版本，暴力求解 时间复杂度 O(N^2)
func twoSum(nums []int, target int) []int {

	re := make([]int, 2)

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if target == (nums[i] + nums[j]) {

				re[0] = i
				re[1] = j

				return re
			}
		}

	}

	return re
}

// https://leetcode.cn/problems/two-sum/
// 求两个数之和
// 借助哈希表
func twoSumHash(nums []int, target int) []int {

	re := make([]int, 2)
	h := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		h[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {

		c := target - nums[i]

		if v, ok := h[c]; ok {

			if i == v {
				continue
			}

			re[0] = i
			re[1] = v

			return re
		}

	}

	return re
}

// https://leetcode.cn/problems/find-the-integer-added-to-array-i/
// 3131找出与数组相加的整数I
// 首次解决
func addedInteger(nums1 []int, nums2 []int) int {

	var re int = 0

	for i := 0; i < len(nums1); i++ {

		if nums1[i] == nums2[i] {
			continue
		}

		if nums1[i] < nums2[i] {
			re += (nums2[i] - nums1[i])
		} else {
			re -= (nums1[i] - nums2[i])
		}
	}

	return re / len(nums1)
}

func Merge(intervals [][]int) [][]int {

	li := make([][]int, 1)
	li[0] = intervals[0]
	index := 0

	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] <= li[index][1] {

			ex := false

			if intervals[i][1] >= li[index][1] {
				li[index][1] = intervals[i][1]
				ex = true
			}

			if intervals[i][0] <= li[index][0] {
				li[index][0] = intervals[i][0]
				ex = true
			}

			if ex {
				continue
			}
		}

		index++
		li = append(li, intervals[i])
	}

	return li
}
