package reorganizstring

import (
	"fmt"
	"sort"
)

type bucket struct {
	c   byte
	num int
}

type bucketSlice []bucket

func (bs bucketSlice) Len() int           { return len(bs) }
func (bs bucketSlice) Less(i, j int) bool { return bs[i].num < bs[j].num }
func (bs bucketSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }

func reorganizeString(S string) string {
	n := len(S)
	bs := make(bucketSlice, 26)
	for i := 0; i < n; i++ {
		b := S[i] - 'a'
		if bs[b].num == 0 {
			bs[b].c = S[i]
		}
		bs[b].num++
	}
	a := sort.Reverse(bs)
	fmt.Println("a", a)

	sort.Sort(a)
	fmt.Println("axxx", a)

	if bs[0].num > (n+1)/2 {
		return ""
	}
	res := make([]byte, n)
	j := 0
	for _, bucket := range bs {
		for i := 0; i < bucket.num; i, j = i+1, j+2 {
			fmt.Println("res", res, j, bucket)
			if j >= n {
				j = 1
			}
			res[j] = bucket.c
		}
	}
	return string(res)
}

func Solution(s string, k int) string {
	mapping := make(map[string]int, len(s))
	charArr := make([]string, 0, len(s))
	for _, c := range s {
		count, ok := mapping[string(c)]
		if !ok {
			mapping[string(c)] = 1
			charArr = append(charArr, string(c))
		}
		mapping[string(c)] = count + 1
	}
	result := ""
	sort.Slice(charArr, func(i, j int) bool {
		return mapping[charArr[i]] > mapping[charArr[j]]
	})
	fmt.Println(mapping, charArr)
	i := 0
	lArr := len(charArr)
	for i < lArr && len(mapping) > 0 {
		count, ok := mapping[charArr[i]]
		if !ok {
			i++
			continue
		}
		j := i + 1
		t := 0
		for j < lArr && t < count {
			limit := 0

			if len(mapping) == 1 {
				if mapping[charArr[i]] == 1 && t == count-1 {
					return result + charArr[i]
				}
				return ""
			}
			result += charArr[i]
			mapping[charArr[i]] -= 1
			if mapping[charArr[i]] == 0 {
				delete(mapping, charArr[i])
			}
			if j == lArr {
				j = i + 1
			}
			if len(mapping) == 0 {
				return result
			}
			for limit != k && j < lArr {

				charCount, ok := mapping[charArr[j]]
				if !ok {
					j++
					continue
				}
				result += charArr[j]
				mapping[charArr[j]] = charCount - 1
				limit++
				t++
				if mapping[charArr[j]] == 0 {
					delete(mapping, charArr[j])
				}
				j++
			}
		}

	}

	if len(mapping) > 1 {
		return ""
	}
	for key, v := range mapping {
		if v > k {
			return ""
		}
		result += key
	}
	return result
}

func Solution2(s string, k int) string {
	lS := len(s)

	mapping := make(map[rune]int, len(s))
	charArr := make([]rune, 0, len(s))
	for _, c := range s {
		count, ok := mapping[c]
		if !ok {
			mapping[c] = 1
			charArr = append(charArr, c)
		}
		mapping[c] = count + 1
	}
	result := make([]byte, len(s))
	sort.Slice(charArr, func(i, j int) bool {
		return mapping[charArr[i]] > mapping[charArr[j]]
	})
	fmt.Println("mapping, CharArr", mapping, charArr)
	if mapping[charArr[0]] > (lS+1)/k {
		return ""
	}
	j := 0
	for _, c := range charArr {
		for i := 0; i < mapping[c]; i, j = i+1, j+k {
			if j > lS-1 {
				j = 1
			}
			fmt.Println("result", result)
			result[j] = byte(c)
		}
	}
	return string(result)
}
