package main

import (
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"net/rpc"
)

type Reply struct {
	Client *Wallet
}

type Args struct {
	Arg string
}

func main() {
	gob.Register(elliptic.P256())
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	args := &Args{"test"}
	reply := &Reply{}
	err = client.Call("Wal.CreateWallet", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.Client.Address)
	fmt.Printf("%x\n", reply.Client.PublicKey)
	fmt.Printf("%x\n", reply.Client.PrivateKey)

}
