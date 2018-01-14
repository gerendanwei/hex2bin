package main

import (
	"strconv"
	"log"
)

func Hex2bin(hex string) (str string) {
	if len(hex)%2 != 0 {
		return
	}
	for i := 0; i < len(hex)-1; i += 2 {
		//str += strconv.QuoteRuneToGraphic(rune(hexdec(hex[i:i+2])))
		ascii := rune(hexdec(hex[i:i+2]))
		str += strconv.QuoteRuneToASCII(ascii)[1:2]
	}
	return
}

func Bin2hex(str string) (hex string) {
	for i := 0; i < len(str); i++ {
		hex += dechex(CharCodeAt(str, i))
	}
	return hex
}

func CharCodeAt(str string, i int) int {
	/*	value, _, _, err := strconv.UnquoteChar(CharAt(str, i), '"')
		if err != nil {
			log.Fatalln(err.Error())
			return 0
		}
		return int(value)*/
	return int(str[i])
}

func dechex(number int) string {
	return strconv.FormatInt(int64(number), 16)
}

func hexdec(hex string) int64 {
	i, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		log.Println(err.Error())
	}
	return i
}
