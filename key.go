package main

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/anaskhan96/base58check"
	secp256k1 "github.com/haltingstate/secp256k1-go"
	"golang.org/x/crypto/ripemd160"
)

type KeyPair struct {
	PublicKey  KeyType `json:"public"`
	PrivateKey KeyType `json:"private"`
}

type KeyType struct {
	HexString string `json:"hex"`
	Bytes     []byte `json:"-"`
	Base58    string `json:"b58"`
}

func GenerateKeyPair() KeyPair {
	// pubkey is compressed 33 bytes
	// private key is 32 bytes
	pub, priv := secp256k1.GenerateKeyPair()
	log.Printf("generated secp256k1 keypair:")
	log.Printf("pub(%d) =%x", len(pub), pub)
	log.Printf("priv(%d)=%x", len(priv), priv)

	uncompressPubkey := secp256k1.UncompressPubkey(pub)
	r160hash := ripemd160.New().Sum(sha256.New().Sum(uncompressPubkey))
	base58pub, err := base58check.Encode("80", fmt.Sprintf("%X", r160hash))
	if err != nil {
		log.Fatal("could not base58check encode public key ripemd160 hash")
	}

	pubkey := KeyType{
		HexString: fmt.Sprintf("%x", uncompressPubkey),
		Bytes:     uncompressPubkey,
		Base58:    base58pub,
	}

	base58priv, err := base58check.Encode("80", fmt.Sprintf("%X", priv))
	if err != nil {
		log.Fatal("could not base58check encode private key")
	}

	privkey := KeyType{
		HexString: fmt.Sprintf("%x", priv),
		Bytes:     priv,
		Base58:    base58priv,
	}

	return KeyPair{pubkey, privkey}
}
