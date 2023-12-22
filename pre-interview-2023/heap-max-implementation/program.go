package main

import (
	"container/heap"
	"fmt"
)

type FoodRatings struct {
	foods        map[string]*Foods
	foodCuisines map[string]string
}

type Food struct {
	Name    string
	Cuisine string
	Rating  int
	HeapIdx int
}

func main() {
	foods := &Foods{}

	heap.Push(foods, &Food{
		Name:    "Pizza",
		Cuisine: "Korea",
		Rating:  5,
	})
	heap.Push(foods, &Food{
		Name:    "aaaa",
		Cuisine: "Korea",
		Rating:  6,
	})
	heap.Push(foods, &Food{
		Name:    "aaaa",
		Cuisine: "Korea",
		Rating:  9,
	})
	heap.Push(foods, &Food{
		Name:    "aaaa",
		Cuisine: "Korea",
		Rating:  11,
	})
	heap.Push(foods, &Food{
		Name:    "aaaa",
		Cuisine: "Korea",
		Rating:  20,
	})
	heap.Push(foods, &Food{
		Name:    "aaaa",
		Cuisine: "Korea",
		Rating:  1,
	})
	fmt.Println("highest rating", foods.Peak())
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	return FoodRatings{}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {

}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return ""
}

var _ (heap.Interface) = (*Foods)(nil)

type Foods []*Food

// Less implements heap.Interface.
func (p Foods) Less(i int, j int) bool {

	if p[i].Rating == p[j].Rating {
		return p[i].Name < p[j].Name
	}
	return p[i].Rating > p[j].Rating

}

// Pop implements heap.Interface.
func (p *Foods) Pop() any {
	n := len(*p)
	product := (*p)[n-1]
	*p = (*p)[:n-1]
	fmt.Println("Pop", product)
	return product
}

func (p *Foods) Peak() any {
	return (*p)[0]
}

// Push implements heap.Interface.
func (p *Foods) Push(x any) {
	product, ok := x.(*Food)
	if !ok {
		return
	}
	product.HeapIdx = len(*p)
	*p = append(*p, product)
}

func (p Foods) Len() int {
	return len(p)
}

func (p Foods) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].HeapIdx, p[j].HeapIdx = i, j

}
