package program

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return (points[i][0]*points[i][0] + points[i][1]*points[i][1]) < (points[j][0]*points[j][0] + points[j][1]*points[j][1])
	})
	fmt.Sprintf()
	return points[:k+1]
}

type response struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []*Country `json:"data"`
}

type Country struct {
	Name         string   `json:"page"`
	Capital      string   `json:"capital"`
	CallingCodes []string `json:"callingCodes"`
}

/*
 * Complete the 'getPhoneNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING country
 *  2. STRING phoneNumber
 * API URL: https://jsonmock.hackerrank.com/api/countries?name=<country>
 */

func getPhoneNumbers(country string, phoneNumber string) string {
	callingCodes, err := getCallingCodes(country)
	if err != nil || len(callingCodes) == 0 {
		return "-1"
	}
	callingCode := callingCodes[0]
	return generatePhone(callingCode, phoneNumber)
}

func generatePhone(callingCode, phoneNumber string) string {
	return fmt.Sprintf("%s %s", callingCode, phoneNumber)
}

func getCallingCodes(country string) ([]string, error) {

	resp, err := http.Get("https://jsonmock.hackerrank.com/api/countries?name=" + country)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res response
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}
	if res.Total > 0 {
		return res.Data[0].CallingCodes
	}
}
