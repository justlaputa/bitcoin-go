package main

import (
	"log"
)

const (
	TestAmount = 0.00001
)

func main() {
	config := LoadConfig()

	if config.SenderKeys.PublicKey.HexString == "" {
		config.SenderKeys = GenerateKeyPair()
		config.ReceiverKeys = GenerateKeyPair()
		SaveConfig(config)
	}

	//receiverAddress := config.ReceiverKeys.PublicKey.Base58
	tx := Transaction{}

	conn, _ := ConnectBitcoinNetwork()

	log.Printf("sending transaction: %s", tx)
	msg, _ := conn.SendTX(tx)

	log.Printf("got message: %s", msg)
}
