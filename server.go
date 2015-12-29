package main

import (
	"net"
	"log"
	"unicode/utf8"
	"bytes"
)

func defaultMessaging(conn net.Conn) {
	var buff = make([]byte, 1024)

	for {
		var strBuf bytes.Buffer
		byteLength, err := conn.Read(buff)
		if err != nil {
			log.Println("Disconnected : " + err.Error())
			break
		}

		for byteLength > 0 {
			r, size := utf8.DecodeRune(buff)
			strBuf.WriteString(string(r))
			buff = buff[size:]
			byteLength = byteLength - size
		}

		log.Print(strBuf.String())
		conn.Write([]byte("hello world"))
	}
}

func main() {
	log.Printf("Hello world!")
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(conn.RemoteAddr())
		go defaultMessaging(conn)
	}
}
