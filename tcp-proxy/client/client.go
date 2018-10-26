package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go sendHTTP(i)

		time.Sleep(2 * time.Second)
	}
}

func sendTCP(i int) {
	addr := "127.0.0.1:8081"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for j := 0; j < 100; j++ {
		fmt.Fprintf(conn, "req%d\n", i*1000+j)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("response=" + message)

		time.Sleep(100 * time.Millisecond)
	}
}

func sendHTTP(i int) {
	values := url.Values{}
	values.Add("index", fmt.Sprintf("%d", i))
	http.Get(fmt.Sprintf("http://127.0.0.1:8081/?%s", values.Encode()))
}
