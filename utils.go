package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

//IntToHex Converts an int64 to byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
