package main

import (
	"fmt"

	"github.com/pion/stun"
)

func main() {
	// Parse a STUN URI
	u, err := stun.ParseURI("stun:stun.l.google.com:19302")

	if err != nil {
		panic(err)
	}

	c, err := stun.DialURI(u, &stun.DialConfig{})

	if err != nil {
		panic(err)
	}

	//Building binding requst with random transaction id
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	//Sending request to STUN server and wait for response msg
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		fmt.Println("your IP is", xorAddr.IP)
	}); err != nil {
		panic(err)
	}
}
