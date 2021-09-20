package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Words struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	answer := make([]string, 0, 10)
	if len(text) == 0 { // пришла пустая строка
		return answer
	}

	newStrings := strings.Fields(text) // разбили строку по одному или нескольким пробелам

	// собрали мапу со словами (ключ) и количеством повторов (значение)
	words := make(map[string]int)
	for _, word := range newStrings {
		words[word]++
	}

	// структура для сортировки по количеству
	structWords := make([]Words, 0, len(words))
	// для лексикографической сортировки нужно собрать мапу, где ключ - количество повторений, значения - массив слов
	resultMap := make(map[int][]string)

	for k, v := range words {
		structWords = append(structWords, Words{k, v})
		resultMap[v] = append(resultMap[v], k)
	}
	// отсортировали по кол-ву
	sort.Slice(structWords, func(i, j int) bool {
		return structWords[i].Count > structWords[j].Count
	})

	// Тут видимо надо все же сначала сортировать а потом уже брать 10 слов
	// иначе сортировка лексикографически будет неверная
	// structWords = structWords[:10]

	for _, word := range structWords {
		app := resultMap[word.Count]
		sort.Strings(app)
		answer = append(answer, app...)
		delete(resultMap, word.Count)
	}
	// взяли 10 самых больших повторений
	answer = answer[:10]

	// удаляем пустые значения
	resAnswer := []string{}
	for _, word := range answer {
		if word != "" {
			resAnswer = append(resAnswer, word)
		}
	}

	return resAnswer
}
