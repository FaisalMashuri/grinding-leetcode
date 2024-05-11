package twoSum_II

func TwoPointerSolution(nums []int, target int) (result []int) {
	left := 0
	right := len(nums) - 1
	//
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left + 1, right + 1}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}

	//sum := nums[left] + nums[right]
	//// Jika penjumlahan sama dengan target, kembalikan indeks
	//if sum == target {
	//	return []int{left + 1, right + 1}
	//} else if sum < target {
	//	// Jika penjumlahan lebih kecil dari target, geser pointer kiri ke kanan
	//	left++
	//} else {
	//	// Jika penjumlahan lebih besar dari target, geser pointer kanan ke kiri
	//	right--
	//}
	//return []int{}
}
