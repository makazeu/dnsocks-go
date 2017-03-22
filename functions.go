package main

func UDP2TCP(len int, data []byte) []byte {
	tmpStr	:= string('\x00') + string(rune(len)) + string(data)
	return []byte(tmpStr)
	//fmt.Println(hex.Dump(tmpByte))
}

/* function to verify fatal error */
func CheckFatalError(err error) {
	if err != nil {
		FatalLogger(err)
	}
}
