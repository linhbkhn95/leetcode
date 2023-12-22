package shufflearray

import "math/rand"

type Solution struct {
	origin []int
	nums   []int
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
		nums:   append([]int{}, nums...),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.nums), func(i, j int) {
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	})
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
