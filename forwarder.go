package main

import(
	"fmt"
	"net"
	_"time"
	"encoding/hex"
)

func DNSocks(conn *net.UDPConn, len int, addr *net.UDPAddr, data []byte){
	query := UDP2TCP(len, data[0:len])
	
	response := forwardQuery(query)

	/* Reply to client */
	fmt.Println(hex.Dump(response))
	sendResponse(conn, addr, response)
	//time.Sleep(time.Second * 10)
}

func forwardQuery(data []byte) []byte {
	response := make([]byte, SIZE_TCP_REPLY)

	// Establish a TCP connection
	serverAddr := dnsConfig.dns_address + ":" + dnsConfig.dns_port
	severConn, err := net.DialTimeout("tcp", serverAddr, SEC_TIMEOUT)
	if err != nil {
		ErrorOutput(err)
		return response
	}

	_, err = severConn.Write(data)
	if err != nil {
		ErrorOutput(err)
		return response
	}

	len, err := severConn.Read(response)
	if err != nil {
		ErrorOutput(err)
		return response
	}

	return response[2:len]
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, response []byte) {
    _,err := conn.WriteToUDP(response, addr)
    if err != nil {
        ErrorOutput(err)
    }
}