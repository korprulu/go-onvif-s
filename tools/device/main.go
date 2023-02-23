package main

import (
	"net"

	"github.com/rs/zerolog/log"
)

func main() {
	ipaddr, _ := net.ResolveUDPAddr("udp", "239.255.255.250:3702")
	conn, err := net.ListenMulticastUDP("udp4", nil, ipaddr)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	buf := make([]byte, 8069)

	n, srcAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	log.Info().Msg(string(buf[:n]))

	_, err = conn.WriteToUDP([]byte("hello world"), srcAddr)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
