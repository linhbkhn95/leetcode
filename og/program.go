package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println('0', '9')
	fmt.Println('a', 'z')
}

func Solution(s string) bool {
	arr := strings.Split(s, " ")

	s1, s2 := arr[0], arr[1]
	i, j := 0, 0
	for i < len(s1) {
		switch s1[i] {
		case '$':
			{
				if s2[j] > 48 && s2[j] <= 57 {
					i++
					j++
					continue
				}
				return false
			}
		case '+':
			{
				if s2[j] >= 97 && s2[j] <= 122 {
					i++
					j++
					continue
				}
				return false
			}
		case '*':
			{
				if i < len(s1)-1 && s1[i+1] == '{' {
					num, _ := strconv.Atoi(string(s1[i+2]))
					c := s2[j]
					for k := 0; k < num; k++ {
						if c != s2[j+k] {
							return false
						}
					}
					i += 4
					j += num
				} else {
					c := s2[j]
					for k := 0; k < 3; k++ {
						if c != s2[j+k] {
							return false
						}
					}
					i++
					j += 3
				}
			}
		}
	}
	return i == len(s1) && j == len(s2)
}
