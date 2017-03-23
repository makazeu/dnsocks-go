package main

import (
	"net"
	"time"
	"errors"
	"strconv"
	"strings"
)

func DialSocks5(proxy, target, port string, timeout time.Duration) (conn net.Conn, err error) {
	// Dial TCP with timeout
	conn, err = net.DialTimeout("tcp", proxy, timeout)
	if err != nil {
		return
	}
	
	// RFC1928 - SOCKS Protocol Version 5
	// Version identifier/method selection
	request := []byte {
		5,	// protocol version
		1,	// method identifier
		0,  // no authentication required
	}
	response, err := sendReceive(conn, request)
	if err != nil {
		return
	} else if len(response) != 2 {
		err = errors.New("Bad response from proxy server.")
		return
	} else if response[0] != 5 {
		err = errors.New("SOCKS5 not supported by proxy server.")
		return
	} else if response[1] != 0 {
		err = errors.New("Authentication required.")
		return
	}

	// Requests
	request = []byte {
		5,	// VER - protocol version
		1,	// CMD - connect
		0,	// RSV - reserved
		1,	// ATYP - IPv4 address
	}

	convByte, err := getAddressBytes(target)
	if err != nil {
		return
	}
	request = append(request, convByte...)

	convByte, err = getPortBytes(port)
	if err != nil {
		return
	}
	request = append(request, convByte...)
	
	response, err = sendReceive(conn, request)
	if err != nil {
		return
	} else if len(response) != 10 {
		err = errors.New("Bad response from proxy server.")
	} else if response[1] != 0 {
		err = errors.New("SOCKS5 connection failed.")
	}

	return
}

func sendReceive(conn net.Conn, request []byte) (response []byte, err error) {
	_, err = conn.Write(request)
	if err != nil {
		return
	}

	response, err = readAll(conn)
	return
} 

func readAll(conn net.Conn) (response []byte, err error) {
	response = make([]byte, 1024)
	n, err := conn.Read(response)
	response = response[:n]
	return
}

func getAddressBytes(address string) (result []byte, err error) {
	result = make([]byte, 4)
	strs := strings.Split(address, ".")
	if len(strs) != 4 {
		err = errors.New("Bad proxy address - " + address)
		return
	}

	for i,s := range strs {
		n, err1 := strconv.Atoi(s)
		if err1 != nil {
			err = err1
			return
		}

		if n < 0 || n > 255 {
			err = errors.New("Bad proxy address - " + address)
			return
		}

		result[i] = byte(n)
	}

	return 
}

func getPortBytes(port string) (result []byte, err error) {
	result = make([]byte, 2)
	n, err := strconv.Atoi(port)
	if err != nil {
		return
	}

	if n < 0 || n > 65525 {
		err = errors.New("Bad proxy port - " + port)
		return
	}

	higher	:= n / 256
	lower	:= n % 256

	result[0] = byte(higher)
	result[1] = byte(lower)
	return 
}
