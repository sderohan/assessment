package service

import (
	"sort"
	"strings"

	"github.com/sderohan/assessment.git/payload"
)

func GetTopTenWords(text string) ([]*payload.TopWordResponse, error) {

	// array to hold the word count result
	var result []*payload.TopWordResponse

	// dictionary to hold the word count
	dictionary := make(map[string]int)

	// split the string into array of word
	words := strings.Split(text, " ")

	// count the words
	for _, word := range words {
		dictionary[word]++
	}

	// make the object of the word and count
	for word, count := range dictionary {
		wordCount := &payload.TopWordResponse{
			Word:  word,
			Count: count,
		}
		result = append(result, wordCount)
	}

	// sort the word count based on the value
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	// always return top 10
	if len(result) >= 10 {
		result = result[:10]
	}

	return result, nil
}
