package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Dictionary struct {
	Key    int
	Values []string
}

func counter(stringArray []string) map[string]int {
	dict := make(map[string]int)
	for _, word := range stringArray {
		dict[word]++
	}

	return dict
}

func groupByCount(counter map[string]int) map[int][]string {
	dict := make(map[int][]string, len(counter))
	for word, count := range counter {
		dict[count] = append(dict[count], word)
	}

	return dict
}

func getSortedArrayDict(counter map[int][]string) []Dictionary {
	arrayDict := make([]Dictionary, 0, len(counter))
	for count, words := range counter {
		arrayDict = append(arrayDict, Dictionary{count, words})
	}
	sort.Slice(arrayDict, func(i, j int) bool {
		return arrayDict[i].Key > arrayDict[j].Key
	})

	return arrayDict
}

func getMostCommonWords(arrayDict []Dictionary) []string {
	var mostCommon []string
	count := 0
	for _, dict := range arrayDict {
		sort.Strings(dict.Values)
		mostCommon = append(mostCommon, dict.Values...)
		count += len(dict.Values)
		if count >= 10 {
			mostCommon = mostCommon[:10]
			break
		}
	}

	return mostCommon
}

func Top10(text string) []string {
	counter := counter(strings.Fields(text))
	intCounter := groupByCount(counter)
	arrayDict := getSortedArrayDict(intCounter)

	return getMostCommonWords(arrayDict)
}
