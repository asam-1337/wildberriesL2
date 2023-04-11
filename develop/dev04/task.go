package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

loop:
	for _, word := range words {
	innerLoop:
		for k, v := range anagrams {
			if len(k) == len(word) {
				for _, let := range word {
					if !strings.Contains(k, string(let)) {
						continue innerLoop
					}
				}

				anagrams[k] = append(v, word)
				continue loop
			}
		}

		anagrams[word] = append(anagrams[word], word)

	}

	for k, v := range anagrams {
		if len(v) <= 1 {
			delete(anagrams, k)
			continue
		}

		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	return anagrams
}

func main() {
	sl := []string{"пятак", "листок", "слиток", "пятка", "тяпка", "столик", "кулик"}
	fmt.Println(getAnagrams(sl))
}
