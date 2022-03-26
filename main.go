package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	signerKey, _ := crypto.GenerateKey()
	key, _ := crypto.GenerateKey()

	signerAddr := crypto.PubkeyToAddress(signerKey.PublicKey).String()
	addr := crypto.PubkeyToAddress(key.PublicKey).String()

	fmt.Println("signerKey ", common.Bytes2Hex(crypto.FromECDSA(signerKey)))
	fmt.Println("signerAddr ", signerAddr)
	fmt.Println("key ", common.Bytes2Hex(crypto.FromECDSA(key)))
	fmt.Println("addr ", addr)

	tx := types.NewTransaction(0, common.Address{},
		big.NewInt(0), 100000, big.NewInt(100000000000),
		common.Hex2Bytes("604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf3"))

	txByte, _ := rlp.EncodeToBytes(tx)

	result := make(map[string]interface{})
	result["gasPrice"] = 100000000000
	result["gasLimit"] = 100000
	result["signerAddress"] = strings.ToLower(signerAddr)
	result["transaction"] = "0x" + common.Bytes2Hex(txByte)
	result["address"] = strings.ToLower(addr)

	s, _ := json.MarshalIndent(result, "  ", "")

	fmt.Println(string(s))
}
