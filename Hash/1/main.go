package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if idx, ok := m[target-n]; ok {
			return []int{i, idx}
		}
		m[n] = i
	}
	return nil
}
