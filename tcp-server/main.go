package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	"golang.org/x/net/netutil"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":50001")
	if err != nil {
		log.Fatal("errror")
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("errror")
	}

	listner := netutil.LimitListener(l, 10)

	log.Printf("info: tcpserver start. (%s)", tcpAddr)

	for {
		conn, err := listner.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				if ne.Temporary() {
					continue
				}
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("info: accept TCP failed. err='%v'", err)
				break
			}
		}

		// 端末のIPアドレスを取得
		ipAddress := strings.Split(conn.RemoteAddr().String(), ":")[0]

		//受信データを読み込む
		buf := bufio.NewReader(conn)
		readData := make([]byte, 1536)
		for {
			n, err := buf.Read(readData)
			if err != nil {
				if err.Error() != "EOF" {
					log.Printf("info : read failed. err='%v'", err)
				}
				return
			} else if n == 0 {
				return
			}
			log.Printf("info: receive_tcp source_ip_address='%s' bytes=%d", ipAddress, n)
			log.Printf("info: payload %X", readData[:n])

		}
	}
}
