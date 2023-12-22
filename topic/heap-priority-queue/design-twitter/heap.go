package main

import "container/heap"

type Feed struct {
	tweetId   int
	timestamp int
}

type HeapMin []*Feed

// Len implements heap.Interface.
func (h HeapMin) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h HeapMin) Less(i int, j int) bool {
	return h[i].timestamp < h[j].timestamp
}

// Pop implements heap.Interface.
func (h *HeapMin) Pop() any {
	n := h.Len()
	item := (*h)[n-1]

	(*h) = (*h)[n-1:]
	return item
}

// Push implements heap.Interface.
func (*HeapMin) Push(x any) {
	panic("unimplemented")
}

// Swap implements heap.Interface.
func (*HeapMin) Swap(i int, j int) {
	panic("unimplemented")
}

var _ heap.Interface = (*HeapMin)(nil)

type Twitter struct {
	followers       map[int]map[int]struct{}
	feeds           map[int][]*Feed
	currentTimstamp int
}

func Constructor() Twitter {
	return Twitter{
		followers: make(map[int]map[int]struct{}),
		feeds:     map[int][]*Feed{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.feeds[userId] = append(this.feeds[userId], &Feed{
		tweetId:   tweetId,
		timestamp: this.currentTimstamp,
	})
	this.currentTimstamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feedsForUsers := this.feeds[userId]

	followeeIds, ok := this.followers[userId]

	if ok {
		for u := range followeeIds {
			userFeeds := this.feeds[u]
			feedsForUsers = mergeSortedArr(feedsForUsers, userFeeds)
		}
	}
	result := make([]int, 0, 10)
	l := len(feedsForUsers)
	for i := len(feedsForUsers) - 1; i >= 0; i-- {
		if l-1-i == 10 {
			break
		}
		result = append(result, feedsForUsers[i].tweetId)
	}

	return result

}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, ok := this.followers[followerId]
	if !ok {
		this.followers[followerId] = make(map[int]struct{})
	}
	this.followers[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.followers[followerId], followeeId)

}

func mergeSortedArr(arr1, arr2 []*Feed) []*Feed {
	l1, l2 := len(arr1), len(arr2)
	i, j := 0, 0
	result := make([]*Feed, 0, l1+l2)
	for i < l1 && j < l2 {
		for i < l1 && j < l2 && arr1[i].timestamp < arr2[j].timestamp {
			result = append(result, arr1[i])
			i++
		}
		for i < l1 && j < l2 && arr1[i].timestamp > arr2[j].timestamp {
			result = append(result, arr2[j])
			j++
		}
	}
	if i < l1 {
		for k := i; k < l1; k++ {
			result = append(result, arr1[k])
		}
	}
	if j < l2 {
		for k := j; k < l2; k++ {
			result = append(result, arr2[k])
		}
	}
	return result
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
