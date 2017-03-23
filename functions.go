package main

import (
	"strings"
)

func UDP2TCP(length int, data []byte) []byte {
	return append(
		[]byte{
			0,
			byte(length),
		},
		data...
	)
}

func GetDomainName(length int, query []byte) string {
	if length < 12 {
		return "Bad Request"
	}

	domain	:= ""
	pos 	:= 12

	if string(query[1:5]) == "\x01\x00\x00\x01" {
		pos = 11
	}

	for pos < length {
		//fmt.Println(rune(query[pos]), int(query[pos]))
		if rune(query[pos]) > 0 {
			domain += ( string(query[pos+1 : pos+1+int(query[pos])]) + "." )
			pos += (int(query[pos]) + 1)
		} else {
			break
		}
	}

	if strings.HasSuffix(domain, ".") {
		domain = domain[0: len(domain) - 1]
	}

	return domain
}

/* function to verify fatal error */
func CheckFatalError(err error) {
	if err != nil {
		FatalLogger(err)
	}
}


func Welcome() {
	CommonOutput("╔═════════════════╗")
	CommonOutput("║▁▂▃▄▅▆▇███▇▆▅▄▃▂▁║")
	CommonOutput("║            DNSocks-Go            ║")
	CommonOutput("║              v" + VERSION + "              ║")
	CommonOutput("╠─────────────────╣")
	CommonOutput("║         © 2017 Makazeu           ║")
	CommonOutput("╚═════════════════╝")
}