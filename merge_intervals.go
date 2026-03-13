package main

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	res := make([][]int, 0)

	// 先按照区间的左端点排序
	sort.Slice(
		intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)

	for _, interval := range intervals {

		// 首次 append 或者是当前 append 的区间的左端点 大于已经合并的右端点，将是一次追加操作
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
			continue
		}

		// 已经合并的最后区间用来做比较
		last := res[len(res)-1]

		// 当前 append 的区间的左端点 小于等于 已经合并的右端点，将是一次合并操作
		if interval[0] <= last[1] {
			last[1] = max(last[1], interval[1])
			continue
		}

	}
	return res
}
