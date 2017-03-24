## 功能  Features
- 以TCP协议转发DNS请求  Forwards DNS queries via TCP
- 支持基于SOCKS5的代理  Supports proxy based on SOCKS5 

## 下载 Downloads
- 您可以通过``git clone``等途径获取本项目的源代码，然后使用``go build``进行编译
- 您也可以直接[下载](https://github.com/zyfworks/dnsocks-go/releases)编译后的二进制文件使用，不需要相关编译环境

## 配置文件  Configuration
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