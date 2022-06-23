package main

import (
	"crypto/ecdsa"
	"log"
	"net"
	"net/rpc"
	"os"
)

var logFile *log.Logger = GetlogFile()

func GetlogFile() *log.Logger {
	f, _ := os.OpenFile("rpc.log", os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	defer f.Close()

	return log.New(f, "[INFO]", log.LstdFlags)
}

type Wa Wallet

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
	Alias      string
	Timestamp  int64
}

func main() {

	rpc.Register(new(Wa))
	ln, err := net.Listen("tcp", ":7000")
	if err != nil {
		logFile.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		defer conn.Close()
		go rpc.ServeConn(conn)
	}

}
