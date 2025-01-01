package finddiff

import "fmt"

func findDiff(a, b map[string]interface{}, currentPath, sameKey string, result map[string]interface{}) {
	sameKeys := make([]string, 0)
	for kA, valueA := range a {
		valueB, ok := b[kA]
		if ok {
			sameKeys = append(sameKeys, kA)
		} else {
			result[fmt.Sprintf("%s.")]
		}

	}
}
