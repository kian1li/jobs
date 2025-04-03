package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"realestatetransfer"
	"realestatetransfer/chaincode"
)

func main() {
	// interface check
	var _ realestatetransfer.RealestateTransfer =  (*chaincode.RealestateTransferCC)(nil)

	err := shim.Start(new(chaincode.RealestateTransferCC))
	if err != nil {
		fmt.Printf("Error in chaincode process: %s", err)
	}
}
