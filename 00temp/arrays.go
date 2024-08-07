package temp

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
