package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	lines, err := ReadFile("./test.txt")
	if err != nil {
		t.Error(err)
	}

	for _, v := range lines {
		fmt.Println(v)
	}
}