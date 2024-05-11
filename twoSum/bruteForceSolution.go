package twoSum

func BruteForceSolution(nums []int, target int) (result []int) {
	for i, num1 := range nums {
		for j, num2 := range nums[:i] {
			if num1+num2 == target {
				return []int{j, i}
			}
		}
	}
	return []int{}
}
