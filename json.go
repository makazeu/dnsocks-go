package main

import (
	"encoding/json"
	"io/ioutil"
)

type DNSocksConfig struct {
	ListenAddress	string `json:"listenAddress"`
	ListenPort		string `json:"listenPort"`
	DnsAddress 		string `json:"dnsAddress"`
	DnsPort    		string `json:"dnsPort"`
	ProxyEnabled	bool   `json:"proxyEnabled"`
	ProxyAddress	string `json:"proxyAddress"`
	ProxyPort		string `json:"proxyPort"`
}

func ReadConfig() {
	var dnsocksConfig DNSocksConfig

	defer CommonOutput("")
	raw, err := ioutil.ReadFile(CONFIG_FILENAME)
	if err != nil {
		CommonOutput(" !!Cannot open config file!!")
		CommonOutput(" Using default config currently")
		return
	}

	json.Unmarshal(raw, &dnsocksConfig)

	if len(dnsocksConfig.ListenAddress) == 0 || 
		len(dnsocksConfig.ListenPort) == 0 ||
		len(dnsocksConfig.DnsAddress) == 0 ||
		len(dnsocksConfig.DnsPort) == 0 || 
		(dnsocksConfig.ProxyEnabled && (
			 len(dnsocksConfig.ProxyAddress) == 0 ||
			 len(dnsocksConfig.ProxyPort) == 0)) {
		CommonOutput(" !!Broken config file!!")
		CommonOutput(" Using default config currently")
		return
	}
	
	dnsConfig.listen_host 	= dnsocksConfig.ListenAddress
	dnsConfig.listen_port 	= dnsocksConfig.ListenPort
	dnsConfig.dns_address 	= dnsocksConfig.DnsAddress
	dnsConfig.dns_port 		= dnsocksConfig.DnsPort
	dnsConfig.proxy_enabled	= dnsocksConfig.ProxyEnabled
	dnsConfig.proxy_address	= dnsocksConfig.ProxyAddress
	dnsConfig.proxy_port 	= dnsocksConfig.ProxyPort
	CommonOutput(" Successfully read the config file")
}