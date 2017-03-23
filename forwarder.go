package main

import(
	_"fmt"
	"net"
	_"encoding/hex"
	"time"
)

func DNSocks(conn *net.UDPConn, length int, addr *net.UDPAddr, data []byte){
	//fmt.Println(hex.Dump(data[0:length]))
	domain	:= GetDomainName(length, data[0:length])
	query 	:= UDP2TCP(length, data[0:length])

	debugInfo := addr.String() + " -> " + domain
	
	response, err := forwardQuery(query)
	if err != nil {
		response = response[:1]
		debugInfo += " (failed) "
		ErrorOutput(err)
	} else {
		debugInfo += " (ok)"
	}

	/* Debug Output */
	if DEBUG_MODE {
		LogOutput(debugInfo)
	}

	/* Reply to client */
	sendResponse(conn, addr, response)
}

func forwardQuery(data []byte) (response []byte, err error) {
	response = make([]byte, SIZE_TCP_REPLY)

	// Establish a TCP connection
	var severConn net.Conn
	if dnsConfig.proxy_enabled {
		severConn, err = DialSocks5(
			dnsConfig.proxy_address + ":" + dnsConfig.proxy_port,
			dnsConfig.dns_address,
			dnsConfig.dns_port,
			SEC_TIMEOUT)
	} else {
		serverAddr := dnsConfig.dns_address + ":" + dnsConfig.dns_port
		severConn, err = net.DialTimeout("tcp", serverAddr, SEC_TIMEOUT)
	}
	
	if err != nil {
		return 
	}
	defer severConn.Close()

	// timeout config
	severConn.SetWriteDeadline(time.Now().Add(SEC_TIMEOUT))
	severConn.SetReadDeadline(time.Now().Add(SEC_TIMEOUT))

	// request
	_, err = severConn.Write(data)
	if err != nil {
		return
	}

	// response
	n, err := severConn.Read(response)
	if err != nil {
		return
	}

	response = response[2:n]
	return
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, response []byte) {
    _,err := conn.WriteToUDP(response, addr)
    if err != nil {
        ErrorOutput(err)
    }
}