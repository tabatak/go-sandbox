package main

import (
	"log"
	"net"
)

func main() {
	addr := "0.0.0.0:60009"
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("%v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("%v", err)
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
