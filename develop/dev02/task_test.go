package main

import "testing"

type TestCase struct {
	src string
	ans string
}

func TestUnpackString(t *testing.T) {
	cases := []TestCase{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for i, c := range cases {
		res, err := UnpackString(c.src)
		if err != nil {
			t.Error(err)
		}

		if res != c.ans {
			t.Errorf("testCase: %d: want: %s; have: %s", i, c.ans, res)
		}
	}

}
