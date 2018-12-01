package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	//data := "abc123!?$*&()'-=@~"
	data := "VMC:GS_IIB53T8CNeJJdUqzn9V_JnRtQadwWCbl"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)
	fmt.Println(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(uDec)
	fmt.Println(string(uDec))
}
