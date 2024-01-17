package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/**
	 * Your RandomizedSet object will be instantiated and called as such:
	 * obj := Constructor();
	 * param_1 := obj.Insert(val);
	 * param_2 := obj.Remove(val);
	 * param_3 := obj.GetRandom();
	 */
	obj := Constructor()
	fmt.Println(obj.Insert(0))

	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())

}

type RandomizedSet struct {
	indexes    map[int]int
	nums       []int
	currentLen int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 0),
		indexes: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexes[val]
	if ok {
		return false
	}
	this.indexes[val] = this.currentLen
	this.currentLen++
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indexes[val]
	if !ok {
		return false
	}
	if index != this.currentLen-1 {
		old := this.nums[this.currentLen-1]
		this.swap(index, this.currentLen-1)
		this.indexes[old] = index
	}
	delete(this.indexes, val)
	this.nums = this.nums[:this.currentLen-1]
	this.currentLen--
	return true
}

func (this *RandomizedSet) swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.currentLen)
	return this.nums[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
