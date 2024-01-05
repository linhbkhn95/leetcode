package main

import (
	"fmt"
)

func main() {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}
	obj := Constructor(foods, cuisines, ratings)
	fmt.Println(obj.HighestRated("korean"))
	fmt.Println(obj.HighestRated("japanese"))
	obj.ChangeRating("sushi", 16)
	fmt.Println(obj.HighestRated("japanese"))
	obj.ChangeRating("ramen", 16)
	fmt.Println(obj.HighestRated("japanese"))

	// obj := Constructor([]string{"biihw"}, []string{"okxsrcqn"}, []int{13})
	// obj.ChangeRating("biihw", 9)
	// obj.ChangeRating("biihw", 6)

}

// type FoodRatings struct {
// 	foods        map[string]*HeadMax
// 	foodCuisines map[string]string
// }

// func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
// 	l := len(foods)
// 	m := make(map[string]*HeadMax)
// 	foodC := make(map[string]string)
// 	for i := 0; i < l; i++ {
// 		foodC[foods[i]] = cuisines[i]
// 		if _, ok := m[cuisines[i]]; !ok {
// 			m[cuisines[i]] = &HeadMax{
// 				elements: make([]*Element, 0),
// 				indexing: make(map[string]int),
// 			}
// 		}
// 		m[cuisines[i]].Insert(Element{
// 			pritoty: ratings[i],
// 			name:    foods[i],
// 		})
// 	}

// 	return FoodRatings{
// 		foods:        m,
// 		foodCuisines: foodC,
// 	}
// }

// func (this *FoodRatings) ChangeRating(food string, newRating int) {
// 	c, ok := this.foodCuisines[food]
// 	if !ok {
// 		return
// 	}
// 	this.foods[c].ChangePriority(this.foods[c].GetIndex(food), newRating)
// }

// func (this *FoodRatings) HighestRated(cuisine string) string {
// 	return this.foods[cuisine].GetMax().name
// }

// /**
//  * Your FoodRatings object will be instantiated and called as such:
//  * obj := Constructor(foods, cuisines, ratings);
//  * obj.ChangeRating(food,newRating);
//  * param_2 := obj.HighestRated(cuisine);
//  */

// type Element struct {
// 	pritoty int
// 	name    string
// }

// type HeadMax struct {
// 	elements []*Element
// 	indexing map[string]int
// }

// func New(elements []*Element) *HeadMax {
// 	h := &HeadMax{
// 		elements: make([]*Element, 0, len(elements)),
// 	}
// 	idx := make(map[string]int)
// 	h.indexing = idx

// 	for _, e := range elements {
// 		h.Insert(*e)
// 	}
// 	return h
// }

// func (h *HeadMax) Left(i int) int {
// 	return 2*i + 1
// }

// func (h *HeadMax) Right(i int) int {
// 	return 2*i + 2
// }

// func (h *HeadMax) Parent(i int) int {
// 	return (i - 1) / 2
// }

// func (h *HeadMax) Insert(e Element) {
// 	h.elements = append(h.elements, &e)
// 	i := len(h.elements) - 1
// 	h.indexing[e.name] = i
// 	h.shiftUp(i)
// }

// func (h *HeadMax) swap(i, j int) {
// 	h.indexing[h.elements[i].name], h.indexing[h.elements[j].name] = j, i
// 	temp := h.elements[i]
// 	h.elements[i] = h.elements[j]
// 	h.elements[j] = temp
// }

// func (h *HeadMax) ChangePriority(index, pritory int) {
// 	var oldp = h.elements[index].pritoty
// 	h.elements[index].pritoty = pritory

// 	if pritory > oldp {
// 		h.shiftUp(index)
// 		return
// 	}
// 	h.shiftDown(index)

// }

// // Function to shift up the node in order
// // to maintain the heap property
// func (h *HeadMax) shiftUp(i int) {
// 	for h.shouldSwap(i) {
// 		h.swap(h.Parent(i), i)
// 		i = h.Parent(i)
// 	}
// }

// func (h *HeadMax) shiftDown(i int) {
// 	var maxIndex = i

// 	// Left Child
// 	l := h.Left(i)
// 	len := len(h.elements)
// 	if l < len && h.elements[l].pritoty > h.elements[maxIndex].pritoty {
// 		maxIndex = l
// 	}

// 	// Right Child
// 	r := h.Right(i)

// 	if r < len && h.elements[r].pritoty > h.elements[maxIndex].pritoty {
// 		maxIndex = r
// 	}

// 	// If i not same as maxIndex
// 	if i != maxIndex {
// 		h.swap(i, maxIndex)
// 		h.shiftDown(maxIndex)
// 	}
// }

// func (h *HeadMax) GetIndex(name string) int {
// 	return h.indexing[name]
// }

// func (h *HeadMax) GetMax() *Element {
// 	return h.elements[0]
// }

// func (h *HeadMax) shouldSwap(i int) bool {
// 	return i > 0 && (h.elements[h.Parent(i)].pritoty < h.elements[i].pritoty || h.elements[h.Parent(i)].pritoty == h.elements[i].pritoty && (strings.Compare(h.elements[h.Parent(i)].name, h.elements[i].name)) > 0)
// }
