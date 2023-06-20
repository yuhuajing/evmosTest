package main

import (
	"fmt"

	"github.com/gochain/web3"
)

func main() {

	// change to your rpc provider
	var rpcProviderURL = "http://127.0.0.1:8545"
	client := webGetBN(rpcProviderURL)
	b, _ := client.GetBlockNumber()
	fmt.Println(b)
	return
}

func webGetBN(url string) web3.Client {
	//var rpcProviderURL = "http://127.0.0.1:8545"
	client, _ := web3.Dial(url)
	return client

}
