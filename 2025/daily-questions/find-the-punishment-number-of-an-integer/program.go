package findthepunishmentnumberofaninteger

import "strconv"

func punishmentNumber(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		canFind := false
		dfs(0, 0, i, strconv.Itoa(i*i), &canFind)
		if canFind {
			res += i * i
		}

	}
	return res
}

func dfs(i, sum, target int, ar string, canFind *bool) {
	if *canFind {
		return
	}
	if i == len(ar) {
		if sum == target {
			*canFind = true
		}
	}
	for j := i; j < len(ar); j++ {
		n, _ := strconv.Atoi(ar[i : j+1])
		dfs(j+1, sum+n, target, ar, canFind)
	}
}

// func punishmentNumber(n int) int {
//     res := 0
//     for i := 1; i <= n; i++{
//         if dfs(0, 0, i, strconv.Itoa(i*i)){
//             res+=i*i
//         }
//     }
//     return res
// }

// func dfs(i, sum, target int, ar string)bool{
//     if i == len(ar){
//         return sum == target
//     }
//     res := false
//     for j := i; j < len(ar); j++{
//         n,_ := strconv.Atoi(ar[i:j+1])
//         res = res || dfs(j+1, sum+n, target, ar)
//     }
//     return res
// }
