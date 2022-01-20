package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

func GetMaxWordsCount(top int, m map[string]int) map[string]int {

	result := make(map[string]int)

	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}

	if len(a) < top {
		top = len(a)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	count := 0
	for _, k := range a {
		for _, s := range n[k] {
			if count > top {
				continue
			}
			result[s] = k
			count++
		}
	}

	return result
}

func main() {

	http.HandleFunc("/api/gettopwords", TopWordsHandler)
	http.ListenAndServe(":8080", nil)

}

func TopWordsHandler(w http.ResponseWriter, r *http.Request) {

	var Request struct {
		Text string `json:"text"`
	}

	json.NewDecoder(r.Body).Decode(&Request)

	wordsMap := make(map[string]int)

	for _, v := range strings.Split(Request.Text, " ") {
		c, _ := wordsMap[v]
		wordsMap[v] = c + 1
	}

	data := GetMaxWordsCount(10, wordsMap)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
