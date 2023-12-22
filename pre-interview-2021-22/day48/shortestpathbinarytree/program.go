package program

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	// pathStart, strStart := FindPathNode(root, startValue, true)
	// pathDest, strDest := FindPathNode(root, destValue, false)

	// i, j := FindLowestCommonAncestor(pathStart, pathDest)
	// result := ""

	// result += strStart[i:]
	// if j+1 < len(pathDest) {
	// 	result += strDest[j:]
	// }
	// // for k := j + 1; k < len(pathDest); k++ {
	// // 	result += pathDest[k].MoveDirection
	// // }
	var (
		result, pathStart, pathDest string
		startAcestor                = map[int]interface{}{
			root.Val: nil,
		}
		destAcestor = map[int]interface{}{
			root.Val: nil,
		}
		mapping = map[string]int{
			"start":0,
			"dist":0,
		}
	)
	Find1(root, startValue, destValue, &result, &pathStart, &pathDest, startAcestor, destAcestor, mapping)
	return result
}

func FindLowestCommonAncestor(path1, path2 []int) (int, int) {
	mapping := make(map[int]int)
	for i, item := range path1 {
		mapping[item] = i
	}
	i, j := 0, 0

	for k := 1; k < len(path2); k++ {
		if index, ok := mapping[path2[k]]; ok {
			i = index
			j = k
		}
	}
	return i, j
}

func FindPathNode(root *TreeNode, val int, isStart bool) ([]int, string) {
	path := []int{
		root.Val,
	}
	return Find(root, val, path, "", isStart)
}

func Find(root *TreeNode, val int, path []int, stringPath string, isStart bool) ([]int, string) {
	if root == nil {
		return nil, ""
	}

	if root.Val == val {
		return path, stringPath
	}

	if root.Left != nil {
		if isStart {
			stringPath += "U"
		} else {
			stringPath += "L"
		}
		result, str := Find(root.Left, val, append(path, root.Val), stringPath, isStart)

		if len(result) > 0 {
			return result, str
		}
		path = path[:len(path)]
		stringPath = stringPath[:len(stringPath)-1]
	}
	if root.Right != nil {
		if isStart {
			stringPath += "U"
		} else {
			stringPath += "R"
		}
		result, str := Find(root.Right, val, append(path, root.Right.Val), stringPath, isStart)
		if len(result) > 0 {
			return result, str
		}
		path = path[:len(path)]
		stringPath = stringPath[:len(stringPath)-1]
	}
	return nil, ""
}

func Find1(root *TreeNode, start int, dest int, result, pathStart, pathDest *string, startAcestor map[int]interface{}, destAcestor map[int]interface{}, mapping map[string]int) {
	if root == nil {
		return
	}

	if mapping["start"]==0 && root.Val == start {
		mapping["start"] = true
		if mapping["dest"] {
			return
		}
	}
	if !mapping["dest"] && root.Val == dest {
		mapping["dest"] = true
		if mapping["start"] {
			return
		}
	}
	if root.Left != nil && (!mapping["start"] || !mapping["dest"]) {
		if !mapping["start"] {
			startAcestor[root.Left.Val] = nil
		}
		if !mapping["dest"] {
			destAcestor[root.Left.Val] = nil
		}
		Find1(root.Left, start, dest, result, pathStart, pathDest, startAcestor, destAcestor, mapping)
		if !mapping["start"] {
			delete(startAcestor, root.Left.Val)
		} else if *pathDest == "" || {
			*pathStart = "U" + *pathStart
		}
		if !mapping["dest"] {
			delete(destAcestor, root.Left.Val)
		} else if *pathStart == "" {
			*pathDest = "L" + *pathDest
		}
	}
	if root.Right != nil && (!mapping["start"] || !mapping["dest"]) {
		if !mapping["start"] {
			startAcestor[root.Right.Val] = nil
		}
		if !mapping["dest"] {
			destAcestor[root.Right.Val] = nil
		}
		Find1(root.Right, start, dest, result, pathStart, pathDest, startAcestor, destAcestor, mapping)
		if !mapping["start"] {
			delete(startAcestor, root.Right.Val)
		} else {
			*pathStart += "U"
		}
		if !mapping["dest"] {
			delete(destAcestor, root.Right.Val)
		} else {
			*pathDest = "R" + *pathDest
		}
	}
	_, ok1 := startAcestor[root.Val]
	_, ok2 := destAcestor[root.Val]
	if ok1 && ok2 && *result == "" && mapping["start"] && mapping["dest"] {
		fmt.Println("", *pathStart, *pathDest)
		*result = *pathStart + *pathDest
	}

}
