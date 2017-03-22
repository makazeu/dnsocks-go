package main

import(
	_"fmt"
	"net"
	_"encoding/hex"
)

func DNSocks(conn *net.UDPConn, length int, addr *net.UDPAddr, data []byte){
	//fmt.Println(hex.Dump(data[0:length]))
	domain	:= GetDomainName(length, data[0:length])
	query 	:= UDP2TCP(length, data[0:length])

	/* Debug Output */
	if DEBUG_MODE {
		LogOutput(addr.String() + " -> " + domain)
	}
	
	response := forwardQuery(query)

	/* Reply to client */
	sendResponse(conn, addr, response)
}

func forwardQuery(data []byte) []byte {
	response := make([]byte, SIZE_TCP_REPLY)

	// Establish a TCP connection
	serverAddr := dnsConfig.dns_address + ":" + dnsConfig.dns_port
	severConn, err := net.DialTimeout("tcp", serverAddr, SEC_TIMEOUT)
	if err != nil {
		ErrorOutput(err)
		return response[0:1]
	}
	defer severConn.Close()

	_, err = severConn.Write(data)
	if err != nil {
		ErrorOutput(err)
		return response[0:1]
	}

	n, err := severConn.Read(response)
	if err != nil {
		ErrorOutput(err)
		return response[0:1]
	}

	return response[2:n]
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, response []byte) {
    _,err := conn.WriteToUDP(response, addr)
    if err != nil {
        ErrorOutput(err)
    }
}