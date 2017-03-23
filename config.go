package main

import (
	"time"
	"net"
)

const VERSION = "0.3.1"

const (
	SIZE_UDP_QUERY = 1024
	SIZE_TCP_REPLY = 4096

	SEC_TIMEOUT = 7 * time.Second
	
	DEBUG_MODE = true

	CONFIG_FILENAME = "config.json"
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
	dnsConfig.listen_host = "0.0.0.0"
	dnsConfig.listen_port = "53"

	// resolver2.opendns.com
	dnsConfig.dns_address = "208.67.220.220"
	dnsConfig.dns_port = "5353"

	// dns.hinet.net
	/*
	dnsConfig.dns_address = "168.95.1.1"
	dnsConfig.dns_port = "53"
	*/

	//dns.hutchcity.com
	/*
	dnsConfig.dns_address = "202.45.84.58"
	dnsConfig.dns_port = "53"
	*/

	dnsConfig.proxy_enabled = false
	dnsConfig.proxy_address = "127.0.0.1"
	dnsConfig.proxy_port = "1080"

	CommonOutput("")
	CommonOutput("─────I N F O R M A T I O N ─────")
	ReadConfig()
	checkConfig()
	outputConfig()
}

func outputConfig() {
	
	if dnsConfig.proxy_enabled {
		CommonOutput(" SOCKS5 proxy enabled on " + 
			dnsConfig.proxy_address + ":" + dnsConfig.proxy_port)
	}
	CommonOutput(" Remote DNS is " + 
			dnsConfig.dns_address + ":" + dnsConfig.dns_port)
	CommonOutput("")
}

func checkConfig() {
	// DNS
	_, err := net.ResolveTCPAddr("tcp",
		dnsConfig.dns_address + ":" + dnsConfig.dns_port)
	CheckFatalError(err)

	//Proxy
	if dnsConfig.proxy_enabled {
		_, err := net.ResolveTCPAddr("tcp",
		dnsConfig.dns_address + ":" + dnsConfig.dns_port)
		CheckFatalError(err)
	}
}
