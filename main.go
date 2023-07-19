package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	tcommon "github.com/gochain/gochain/v4/common"
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
	var rpcProviderURL = "http://192.168.100.191:8545"
	client := webGetBN(rpcProviderURL)
	b, _ := client.GetBlockNumber()
	fmt.Println(b)
	privateKey, err := crypto.HexToECDSA("e5ff5392711a137f3a4ac680e85ed29cb896427e89b4e0aa582b785722a84c49")
	//0280b75f9a8df2b0330593b9cfbbaac84c0522ef9d2a7568a2595a44835edf60
	//e5ff5392711a137f3a4ac680e85ed29cb896427e89b4e0aa582b785722a84c49
	//db9d9ec53ef566aa9aed60e50c79483f34f5cb289c28b5d0c0b37046b6961347
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// bloc, _ := client.GetBlockByNumber(context.Background(), big.NewInt(2), true)
	// fmt.Println(bloc.TxCount())

	bal, _ := client.GetBalance(context.Background(), address, new(big.Int).SetUint64(b))
	fmt.Println(bal)

	TxCount, _ := client.GetTransactionCount(context.Background(), tcommon.HexToAddress(address))
	fmt.Println(TxCount)

}

func webGetBN(url string) web3.Client {
	//var rpcProviderURL = "http://127.0.0.1:8545"
	client, _ := web3.Dial(url)
	return client
}
