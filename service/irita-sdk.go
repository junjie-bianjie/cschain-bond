package service

import (
	"cschain-bond/utils"
	"fmt"
	"github.com/bianjieai/irita-sdk-go/modules/nft"
	sdk "github.com/bianjieai/irita-sdk-go/types"
)

var (
	baseTx = sdk.BaseTx{
		//From:     s.Account().Name,
		From: "v1",
		Gas:  200000,
		Memo: "test",
		Mode: sdk.Commit,
		//Password: s.Account().Password,
		Password: "YQVGsOjegu",
	}
)

func issueDenom(denomID, denomName, schema string) error {

	issueReq := nft.IssueDenomRequest{
		ID:     denomID,
		Name:   denomName,
		Schema: schema,
	}

	_, err := client.NFT.IssueDenom(issueReq, baseTx)
	if err != nil {
		// TODO handle the error
		return err
	}

	return nil
}

func mintNFT(denomID, tokenID, tokenName, tokenData string) error {
	mintReq := nft.MintNFTRequest{
		Denom: denomID,
		ID:    tokenID,
		Name:  tokenName,
		URI:   fmt.Sprintf("https://%s", utils.RandStringOfLength(10)),
		Data:  tokenData,
	}

	_, err := client.NFT.MintNFT(mintReq, baseTx)
	if err != nil {
		// TODO handle the error
		return err
	}

	return nil
}
