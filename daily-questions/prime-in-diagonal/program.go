package primeindiagonal

func diagonalPrime(nums [][]int) int {
	max := 0
	l := len(nums)
	for i := 0; i < len(nums); i++ {

		if nums[i][i] > max && nums[i][i] != 1 && isPrime(nums[i][i]) {
			max = nums[i][i]
		}

		if nums[l-i-1][i] > max && nums[l-i-1][i] != 1 && isPrime(nums[l-i-1][i]) {
			max = nums[l-i-1][i]
		}
	}
	return max
}

func isPrime(num int) bool {
	if num < 4 {
		return true
	}
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
