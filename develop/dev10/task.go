package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Client struct {
	host     string
	port     string
	connTime time.Duration
}

func NewClient(h, p, cTime string) *Client {
	t, err := strconv.Atoi(cTime)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		host:     h,
		port:     p,
		connTime: time.Duration(t) * time.Second,
	}
}

func (cl *Client) dial(ctx context.Context, cancel context.CancelFunc) {
	conn, err := net.DialTimeout("tcp", cl.host+":"+cl.port, cl.connTime)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			_, err := fmt.Fprintf(conn, "timeout")
			if err != nil {
				log.Print(err)
			}
			log.Println("timeout")
			err = conn.Close()
			if err != nil {
				log.Print(err)
			}
			return
		default:
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("message: ")
			text, err := rd.ReadString('\n')
			if err != nil {
				log.Print("read error...")
			}
			_, err = fmt.Fprintf(conn, text+"\n")
			if err != nil {
				log.Print(err)
			}
			fb, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("from server :" + fb)
		}
	}
}

func main() {
	t := flag.String("timeout", "10", "timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("should be host and port")
	}
	port := args[1]
	host := args[0]

	client := NewClient(host, port, *t)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, client.connTime)

	client.dial(ctx, cancel)

	fmt.Println("Connection closed...")
}
