package main

// import "fmt"

// type Feed struct {
// 	tweetId   int
// 	timestamp int
// }

// func main() {
// 	obj := Constructor()
// 	obj.PostTweet(2, 5)
// 	obj.PostTweet(1, 3)
// 	obj.PostTweet(1, 101)
// 	obj.PostTweet(2, 13)
// 	obj.PostTweet(2, 10)
// 	obj.PostTweet(1, 2)
// 	obj.PostTweet(2, 94)
// 	obj.PostTweet(2, 505)
// 	obj.PostTweet(1, 333)
// 	obj.PostTweet(1, 22)

// 	fmt.Println(obj.GetNewsFeed(2))
// 	obj.Follow(2, 1)
// 	fmt.Println(obj.GetNewsFeed(2))

// }

// type Twitter struct {
// 	followers       map[int]map[int]struct{}
// 	feeds           map[int][]*Feed
// 	currentTimstamp int
// }

// func Constructor() Twitter {
// 	return Twitter{
// 		followers: make(map[int]map[int]struct{}),
// 		feeds:     map[int][]*Feed{},
// 	}
// }

// func (this *Twitter) PostTweet(userId int, tweetId int) {
// 	this.feeds[userId] = append(this.feeds[userId], &Feed{
// 		tweetId:   tweetId,
// 		timestamp: this.currentTimstamp,
// 	})
// 	this.currentTimstamp++
// }

// func (this *Twitter) GetNewsFeed(userId int) []int {
// 	feedsForUsers := this.feeds[userId]

// 	followeeIds, ok := this.followers[userId]

// 	if ok {
// 		for u := range followeeIds {
// 			userFeeds := this.feeds[u]
// 			feedsForUsers = mergeSortedArr(feedsForUsers, userFeeds)
// 		}
// 	}
// 	result := make([]int, 0, 10)
// 	l := len(feedsForUsers)
// 	for i := len(feedsForUsers) - 1; i >= 0; i-- {
// 		if l-1-i == 10 {
// 			break
// 		}
// 		result = append(result, feedsForUsers[i].tweetId)
// 	}

// 	return result

// }

// func (this *Twitter) Follow(followerId int, followeeId int) {
// 	_, ok := this.followers[followerId]
// 	if !ok {
// 		this.followers[followerId] = make(map[int]struct{})
// 	}
// 	this.followers[followerId][followeeId] = struct{}{}
// }

// func (this *Twitter) Unfollow(followerId int, followeeId int) {
// 	delete(this.followers[followerId], followeeId)

// }

// func mergeSortedArr(arr1, arr2 []*Feed) []*Feed {
// 	l1, l2 := len(arr1), len(arr2)
// 	i, j := 0, 0
// 	result := make([]*Feed, 0, l1+l2)
// 	for i < l1 && j < l2 {
// 		for i < l1 && j < l2 && arr1[i].timestamp < arr2[j].timestamp {
// 			result = append(result, arr1[i])
// 			i++
// 		}
// 		for i < l1 && j < l2 && arr1[i].timestamp > arr2[j].timestamp {
// 			result = append(result, arr2[j])
// 			j++
// 		}
// 	}
// 	if i < l1 {
// 		for k := i; k < l1; k++ {
// 			result = append(result, arr1[k])
// 		}
// 	}
// 	if j < l2 {
// 		for k := j; k < l2; k++ {
// 			result = append(result, arr2[k])
// 		}
// 	}
// 	return result
// }

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
