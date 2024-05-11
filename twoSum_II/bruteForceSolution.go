package twoSum_II

func BruteForceSolution(nums []int, target int) (result []int) {
	for i, num1 := range nums {
		for j, num2 := range nums[:i] {
			if num1+num2 == target {
				return []int{j + 1, i + 1}
			}
		}
	}
	return []int{}
}
