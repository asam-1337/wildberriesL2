package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// UnpackString simple unpacking func with repeated runes
func UnpackString(s string) (string, error) {
	if len(s) <= 0 {
		return s, nil
	}

	var (
		err       error
		prevRune  rune
		cnt       int
		isEscaped bool
	)

	res := make([]rune, 0)
	for _, v := range s {
		if isEscaped {
			prevRune = v
			isEscaped = false
			res = append(res, prevRune)
			continue
		}

		if string(v) == `\` {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(v) {
			if prevRune == 0 {
				return "", nil
			}

			cnt, err = strconv.Atoi(string(v))
			if err != nil {
				return string(res), err
			}

			for ; cnt > 1; cnt-- {
				res = append(res, prevRune)
			}

			continue
		}

		prevRune = v
		res = append(res, prevRune)
	}
	return string(res), nil
}

func main() {
	res, err := UnpackString("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
