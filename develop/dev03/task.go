package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
//var (
//	kFlag = flag.String("k", "", "указание колонки для сортировки")
//	nFlag = flag.String("n", "", "сортировать по числовому значению")
//	rFlag = flag.String("k", "", "сортировать в обратном порядке")
//	uFlag = flag.String("k", "", "не выводить повторяющиеся строки")
//)

func ReadFile(path string, uFlag bool) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	rd := bufio.NewReader(f)
	lines := make([]string, 0)
	uniqueMap := make(map[string]struct{})

	for err == nil {
		var data []byte
		data, _, err = rd.ReadLine()
		line := string(data)

		if uFlag {
			if _, ok := uniqueMap[line]; !ok {
				uniqueMap[line] = struct{}{}
				lines = append(lines, line)
			}
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func Sort(lines []string, nFlag, rFlag bool) []string {
	if nFlag {
		sort.Slice(lines, func(i, j int) bool {
			numi, _ := strconv.Atoi(lines[i])
			numj, _ := strconv.Atoi(lines[j])
			return numi < numj && !rFlag
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			return lines[i] < lines[j] && !rFlag
		})
	}

	return lines
}

func main() {
	//kFlag := flag.String("k", "", "указание колонки для сортировки")
	nFlag := flag.Bool("n", false, "сортировать по числовому значению")
	rFlag := flag.Bool("r", false, "сортировать в обратном порядке")
	uFlag := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	lines, err := ReadFile("test.txt", *uFlag)
	if err != nil {
		log.Fatal(err)
	}

	lines = Sort(lines, *nFlag, *rFlag)
	for _, v := range lines {
		fmt.Println(v)
	}
}
