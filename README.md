## 功能  Features
- 以TCP协议转发DNS请求  Forwards DNS queries via TCP
- 支持基于SOCKS5的代理  Supports proxy based on SOCKS5 

## 配置文件  Config file
config.json
- ``listenAddress`` : 本地DNS监听地址 (127.0.0.1 / 0.0.0.0)
- ``listenPort`` : 本地DNS监听端口
- ``dnsAddress`` : 目标DNS的IP地址
- ``dnsPort`` : 目标DNS的端口
- ``proxyEnabled`` : 是否开启SOCKS5代理 (true / false)
- ``proxyAddress`` : SOCKS5代理服务器地址
- ``proxyPort`` : SOCKS5代理服务器端口

```
{
	"listenAddress"	:	"0.0.0.0",
	"listenPort"	:	"53",
	"dnsAddress"	:	"208.67.220.220",
	"dnsPort"		:	"5353",
	"proxyEnabled"	:	false,
	"proxyAddress"	:	"127.0.0.1",
	"proxyPort"	:	"1080"
}
```