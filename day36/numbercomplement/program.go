package program

import "fmt"

// func findComplement(num int) int {
// 	result := []int{}
// 	for num >= 2 {
// 		remain := num % 2
// 		result = append(result, remain)
// 		num = num / 2
// 	}
// 	result = append(result, num)
// 	fmt.Println(result)
// 	total := 0
// 	for i := 0; i < len(result); i++ {
// 		if result[i] == 0 {
// 			total += int(math.Pow(2, float64(i)))
// 		}
// 	}
// 	return total
// }

func findComplement(num int) int {
	res := 0
	base := 1
	for num > 0 {
		if num%2 == 0 {
			res += base
		}

		fmt.Println("base", base)
		num = num >> 1
		fmt.Println("num", num)

		if num == 0 {
			break
		}
		base = base << 1
	}
	return res
}
