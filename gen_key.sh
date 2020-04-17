#!/bin/bash

openssl ecparam -genkey -name secp256k1 -out ./out/key-pair.pem
openssl ec -in ./out/key-pair.pem -outform PEM -out ./out/private.pem