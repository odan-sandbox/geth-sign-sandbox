package main

import "encoding/hex"
import "fmt"
import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/crypto/secp256k1"

func main() {
  private_key, _ := hex.DecodeString("29a1ad7fe945e254474891bb530114b8706cdf89844bce9445570743182a69b1")
  msg, _ := hex.DecodeString("abc")
  digest := crypto.Keccak256(msg)

  signature, err := secp256k1.Sign(digest, private_key)

  fmt.Printf("Hello World\n %v %v", err, signature)
}