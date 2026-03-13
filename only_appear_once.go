package main

func FindNumOnlyAppearOnce(nums []int) int {
	if len(nums)&1 == 0 {
		return 0 // 数组 length应该为奇数
	}

	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num ^= nums[i]
	}

	return num
}
