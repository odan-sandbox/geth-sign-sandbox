package main

import "encoding/hex"
import "fmt"
import "github.com/ethereum/go-ethereum/crypto/secp256k1"

func main() {
  private_key, _ := hex.DecodeString("29a1ad7fe945e254474891bb530114b8706cdf89844bce9445570743182a69b1")
  digest, _ := hex.DecodeString("e28f5ff58ff3f1b24d6ba6e3b3e95e49589e8dd59b91296e76189d6ad2857b22")

  signature, _ := secp256k1.Sign(digest, private_key)

  fmt.Printf("%s\n%s", hex.EncodeToString(digest), hex.EncodeToString(signature))
}