package maximumascendingsubarraysum


func maxAscendingSum(nums []int) int {
    maxSum := 0
    currentSum := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] > nums[i-1] {
            currentSum += nums[i]
        } else {
            currentSum = nums[i]
        }
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }
    return maxSum
}