package main

import (
	"bufio"
	"fmt"
	goPs "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func cd(request []string) {
	switch len(request) {
	case 1:
		fmt.Fprintln(os.Stderr, "need 2 arguments")
	case 2:
		err := os.Chdir(request[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func pwd(request []string) {
	if len(request) == 1 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(path)
		}
	} else {
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func echo(request []string) {
	for i := 1; i < len(request); i++ {
		fmt.Printf("%s ", request[i])
	}
	fmt.Println()
}

func kill(request []string) {
	if len(request) == 1 {
		fmt.Fprintln(os.Stderr, "type process pid")
		return
	}

	if len(request) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}

	pid, err := strconv.Atoi(request[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func ps(request []string) {
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}

	sliceProc, _ := goPs.Processes()

	for _, proc := range sliceProc {

		fmt.Printf("name: %v pid: %v\n", proc.Executable(), proc.Pid())

	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		request := strings.Split(scanner.Text(), " ")
		switch request[0] {
		case "cd":
			cd(request)
		case "pwd":
			pwd(request)
		case "echo":
			echo(request)
		case "kill":
			kill(request)
		case "ps":
			ps(request)
		}
	}
}
