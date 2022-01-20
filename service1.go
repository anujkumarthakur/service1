package servicefirst

import (
	"encoding/json"
	"fmt"
	"sort"
)

func ServiceFirst(words []string) {
	if len(words) > 0 {
		MostUsedWordsInJson(words)
	} else {
		fmt.Println("Empty Text")
	}
}

func MostUsedWordsInJson(words []string) {
	wc := make(map[string]int)
	for _, word := range words {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	keys := make([]string, 0, len(wc))
	for k := range wc {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	jsonData := make(map[string]int)
	for _, k := range keys {
		//fmt.Println(k, wc[k])
		jsonData[k] = wc[k]
	}
	data, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Error From Marshal Indent Json Data :", err)
		return
	}
	fmt.Println(string(data))
}
