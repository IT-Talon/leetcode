package main

func main() {

}

// 示例 1：
//
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
// 偷窃到的最高金额 = 1 + 3 = 4 。
func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	pb := make([]int, len(nums))
	pb[0] = nums[0]
	pb[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		pb[i] = max(pb[i-2]+nums[i], pb[i-1])
	}
	return pb[len(pb)-1]
}
