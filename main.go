package main

import (
	"log"
	"crypto/ecdsa"
)

const (
	TestAmount = 0.00001
)

func main() {
	privateKey := getPrivateKey()
	publicKey := getPublicKey(privateKey)

	myWalletAddr := calcWalletAddress(publicKey)

	receiverAddr := getReceiverAddress()

	tx := NewTransaction(publicKey, receiverAddr, TestAmount)

	conn, err := ConnectBitcoinNetwork()

	log.Printf("sending transaction: %s", tx)
	msg, err := conn.SendTX(tx)

	log.Printf("got message: %s", msg)
}

func getPrivateKey() *ecdsa.PrivateKey {
	// TODO: if private key in config file
	// construct from config file
	// else
	return GeneratePrivateKey()
}

func getPublichKey(p *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return nil
}

func calcWalletAddress(pub *ecdsa.PublicKey) string {
	return ""
}

func getReceiverAddress() string {

}