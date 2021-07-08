#!/bin/sh

solc --abi --bin -o build --overwrite Transfer.sol
abigen --bin="build/Transfer.bin" --abi="build/Transfer.abi" --type="Transfer" --pkg="main" --out="transfer.go"