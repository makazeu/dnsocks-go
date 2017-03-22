package main

import (
	"time"
	"net"
)

type DNS_Config struct {
	/* Local DNS */
	listen_host string
	listen_port string

	/* Remote DNS */
	dns_address string
	dns_port 	string

	/* SOCKS5 Proxy */
	proxy_enabled bool
	proxy_address string
	proxy_port	string
}

var dnsConfig DNS_Config

func InitConfig() {
	dnsConfig.listen_host = "127.0.0.1"
	dnsConfig.listen_port = "53"

	// resolver2.opendns.com
	dnsConfig.dns_address = "208.67.220.220"
	dnsConfig.dns_port = "5353"
	
	dnsConfig.proxy_enabled = false

	checkConfig()
}

const (
	SIZE_UDP_QUERY = 1024
	SIZE_TCP_REPLY = 4096

	SEC_TIMEOUT = 10 * time.Second
	
	DEBUG_MODE = true
)

func checkConfig() {
	// DNS
	_, err := net.ResolveTCPAddr("tcp",
		dnsConfig.dns_address + ":" + dnsConfig.dns_port)
	CheckFatalError(err)

	//Proxy
}
