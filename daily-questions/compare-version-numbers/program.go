package compareversionnumbers

import "strings"

func compareVersion(version1 string, version2 string) int {
	version1Arr := strings.Split(version1, ".")
	version2Arr := strings.Split(version2, ".")
	l1, l2 := len(version1), len(version2)

}

