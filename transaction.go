package main

import (
	secp256k1 "github.com/haltingstate/secp256k1-go"
)

type Script struct {
	Length int
	Data   ScriptBytes
}

type ScriptBytes []byte

type TXIn struct {
	PreviousHash  string
	PreviousIndex int
	ScriptSig     Script
}

type TXOut struct {
	Value        int
	ScriptPubKey Script
}

type Output struct {
	Value           int
	ReceiverAddress string
}

type Transaction struct {
	Input   TXIn
	Outputs []TXOut
}

func (t *Transaction) Bytes() []byte {
	return nil
}

const (
	Version  = 0x01000000
	Locktime = 0x00000000
)

// NewTransaction create a new transaction with one input, multiple outputs
// for simplicity only support single input
func NewTransaction(keypair KeyPair, prevHash string, prevIndex int, scriptPubKey Script, outputs []Output) Transaction {
	txin := TXIn{prevHash, prevIndex, scriptPubKey}
	txouts := []TXOut{}

	for _, output := range outputs {
		txouts = append(txouts, TXOut{output.Value, newScriptPubKey(output.ReceiverAddress)})
	}

	txForSign := newTransaction(txin, txouts)

	sig := secp256k1.Sign(txForSign.Bytes(), keypair.PrivateKey.Bytes)

	scriptSig := newScriptSig(sig, keypair.PublicKey.Bytes)
	txin = TXIn{prevHash, prevIndex, scriptSig}

	return newTransaction(txin, txouts)
}

func newTransaction(txin TXIn, txouts []TXOut) Transaction {
	return Transaction{txin, txouts}
}

func newScriptSig(sig []byte, pubkey []byte) Script {
	return Script{}
}

func newScriptPubKey(receiver string) Script {
	return Script{}
}
