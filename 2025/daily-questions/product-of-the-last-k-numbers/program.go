package main

import "fmt"

type ProductOfNumbers struct {
	products     []int64
	maxZeroIndex int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		products:     make([]int64, 0),
		maxZeroIndex: -1,
	}
}

/*

] = k
products[0] = k
product[0] *= k

*/

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.maxZeroIndex = len(this.products)
		this.products = append(this.products, 1)
		return
	}
	var lastProduct int64 = 1
	if len(this.products) != 0 {
		lastProduct = this.products[len(this.products)-1]
	}
	this.products = append(this.products, lastProduct*int64(num))
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if this.maxZeroIndex > -1 && k >= (len(this.products)-this.maxZeroIndex) {
		return 0
	}
	if len(this.products) == 1 {
		return int(this.products[0])
	}
	if k == len(this.products) {
		return int(this.products[len(this.products)-1])
	}

	return int(this.products[len(this.products)-1] / this.products[len(this.products)-k-1])
}

// 1 , 2 ,3
//  1 2 6
// k = 2 => 6/(1) = 6

func main() {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)

	fmt.Println(obj.GetProduct(2))
	fmt.Println(obj.GetProduct(3))
	fmt.Println(obj.GetProduct(4))
	obj.Add(8)

	fmt.Println(obj.GetProduct(2))

}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
