package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gochain/web3"
)

type Block struct {
	ParentHash      common.Hash
	Sha3Uncles      common.Hash
	Miner           common.Address
	Signers         []common.Address
	Voters          []common.Address
	Signer          []byte
	StateRoot       common.Hash
	TxsRoot         common.Hash
	ReceiptsRoot    common.Hash
	LogsBloom       *types.Bloom
	Difficulty      *big.Int
	TotalDifficulty *big.Int
	Number          *big.Int
	GasLimit        uint64
	GasUsed         uint64
	Timestamp       time.Time
	ExtraData       []byte
	MixHash         common.Hash
	Nonce           types.BlockNonce
	Hash            common.Hash

	// Only one of TxHashes or TxDetails will be populated.
	TxHashes  []common.Hash
	TxDetails []*Transaction

	Uncles []common.Hash
}

type Transaction struct {
	Nonce    uint64
	GasPrice *big.Int // wei
	GasLimit uint64
	To       *common.Address
	Value    *big.Int // wei
	Input    []byte
	From     common.Address
	V        *big.Int
	R        *big.Int
	S        *big.Int
	Hash     common.Hash

	BlockNumber      *big.Int
	BlockHash        common.Hash
	TransactionIndex uint64
}

func main() {
	// change to your rpc provider
	var rpcProviderURL = "http://localhost:8545"
	client := webGetBN(rpcProviderURL)
	b, _ := client.GetBlockNumber()
	fmt.Println(b)
	bloc, _ := client.GetBlockByNumber(context.Background(), big.NewInt(2), true)
	fmt.Println(bloc.TxCount())
	acc, _ := client.GetAccounts()
	for _, v := range acc {
		//account := common.BytesToAddress(v)
		//fmt.Println(account.Hex())
		fmt.Println(v.Hex())
		bal, _ := client.GetBalance(context.Background(), v.Hex(), new(big.Int).SetUint64(b))
		fmt.Println(bal)
		TxCount, _ := client.GetTransactionCount(context.Background(), v)
		fmt.Println(TxCount)
	}
}

func webGetBN(url string) web3.Client {
	//var rpcProviderURL = "http://127.0.0.1:8545"
	client, _ := web3.Dial(url)
	return client
}
