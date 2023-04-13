package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func openFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}

func getExpression(pattern string, ignore bool) (*regexp.Regexp, error) {
	ignorePrefix := ""
	if ignore {
		ignorePrefix = "(?i)"
	}

	compiledExpession, err := regexp.Compile(ignorePrefix + pattern)
	if err != nil {
		return nil, err
	}

	return compiledExpession, nil
}

func getNumberOfIntersections(file []string, expression *regexp.Regexp) int {
	result := 0
	for _, str := range file {
		match := expression.Match([]byte(str))
		if match {
			result++
		}
	}

	return result
}

func reg(file []string, expression *regexp.Regexp, after, before int, number, invert bool) {
	for i, str := range file {
		match := expression.Match([]byte(str))
		if invert && !match {
			echo(file, i, after, before, number)
		} else if !invert && match {
			echo(file, i, after, before, number)
		}
	}
}

func echo(file []string, i, after, before int, number bool) {
	startPoint := 0
	endPoint := len(file)
	if i-after > 0 {
		startPoint = i - after
	}

	if i+before < len(file) {
		endPoint = i + before
	}

	if endPoint != len(file) {
		endPoint += 1
	}

	for line := startPoint; line < endPoint; line++ {
		if number {
			fmt.Printf("%d: ", line+1)
		}
		fmt.Printf("%s\n", file[line])
	}
}

func main() {
	pattern := flag.String("e", "", "паттерн")
	path := flag.String("f", "", "файл")
	after := flag.Int("A", 0, "вывод N строк после совпадения")
	before := flag.Int("B", 0, "вывод N строк до совпадения")
	inTheMiddle := flag.Int("C", 0, "вывод N строк в районе совпадения")
	count := flag.Bool("c", false, "вывести количество строк с совпадением")
	ignore := flag.Bool("i", false, "игнорировать различия регистра")
	invert := flag.Bool("v", false, "инвертировать вывод")
	// fixed := *flag.Bool("F", false, "точное совпадение со строкой")
	number := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	file, err := openFile(*path)
	if err != nil {
		panic(err)
	}
	expression, err := getExpression(*pattern, *ignore)
	if err != nil {
		panic(err)
	}

	if *count {
		result := getNumberOfIntersections(file, expression)
		if *invert {
			result = len(file) - result
		}
		fmt.Println(result)
	} else {
		if *after == 0 && *before == 0 && *inTheMiddle != 0 {
			*after = *inTheMiddle / 2
			*before = *inTheMiddle / 2
		}
		reg(file, expression, *after, *before, *number, *invert)
	}

}
