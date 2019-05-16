package main

import (
	"log"
	"net"
)

// Run はUDPサーバを起動する。
func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":50000")
	if err != nil {
		log.Fatal("errror")
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("errror")
	}

	log.Printf("info: udpserver start. (%s)", udpAddr)

	buf := make([]byte, 1536)
	for {
		n, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("info: read data failed. err='%v'", err)
			return
		}
		if n == 0 {
			continue
		}
		log.Printf("info: receive_udp source_ip_address='%s' bytes=%d", remote.IP.String(), n)
		log.Printf("info: payload %X", buf[:n])

	}
}
