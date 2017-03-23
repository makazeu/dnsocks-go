package main

import (
	"net"
)

func RunDNS() {
	serverAddr, err := net.ResolveUDPAddr("udp", 
		dnsConfig.listen_host + ":" + dnsConfig.listen_port)
	CheckFatalError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckFatalError(err)
	defer serverConn.Close()

	CommonOutput(" DNS started listening at " + serverAddr.String())
	CommonOutput("")

	for {
		handleClient(serverConn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, SIZE_UDP_QUERY)
	n, addr, err := conn.ReadFromUDP(data)

	if err != nil {
		ErrorOutput(err)
		return
	}
	
	go DNSocks(conn, n, addr, data)
}

