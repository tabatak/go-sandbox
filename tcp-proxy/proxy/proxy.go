package main

import (
	"log"
	"net"
)

func main() {

	//とりあえず右から左へ
	addr := "127.0.0.1:8081"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("proxy server start. (%s)", addr)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			defer log.Printf("[proxy] client connection closed.")

			dstAddr := "127.0.0.1:8888"
			dstConn, err := net.Dial("tcp", dstAddr)
			if err != nil {
				log.Fatal(err)
			}
			defer dstConn.Close()

			for {
				buf := make([]byte, 1024)
				size, err := conn.Read(buf)
				if err != nil {
					return
				}
				data := buf[:size]
				// log.Println("[proxy] ", data)
				log.Printf("\n%s", string(data))

				dstConn.Write(data)

				rBuf := make([]byte, 1024)
				dstConn.Read(rBuf)

				conn.Write(rBuf)
			}
		}(conn)
	}

}
