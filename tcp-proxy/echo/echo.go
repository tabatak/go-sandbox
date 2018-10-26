package main

import (
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("echo server start. (%s)", addr)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			defer log.Printf("[echo] connection closed.")

			for {
				buf := make([]byte, 1024)
				size, err := conn.Read(buf)
				if err != nil {
					return
				}
				data := buf[:size]
				// log.Println("[echo] ", data)
				log.Printf("\n%s", string(data))
				conn.Write(data)
			}
		}(conn)
	}
}
