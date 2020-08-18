package service

import (
	"cschain-bond/utils"
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/types"
	"github.com/bianjieai/irita-sdk-go/types/store"
)

var (
	client sdk.IRITAClient
)

const (
	nodeURI       = "tcp://10.2.10.130:36657"
	chainID       = "cschain-bond"
	gas           = 200000
	algo          = "sm2"
	mode          = "commit"
	level         = "info"
	maxTxsBytes   = 1073741824
	gasAdjustment = 1.0

	name   = "v1"
	passwd = "YQVGsOjegu"
)

func init() {
	client = sdk.NewIRITAClient(types.ClientConfig{
		NodeURI: nodeURI,
		ChainID: chainID,
		Gas:     gas,
		Fee: []types.DecCoin{
			{Denom: "point", Amount: types.NewDec(1)},
		},
		Algo:          algo,
		KeyDAO:        store.NewMemory(nil),
		Mode:          mode,
		Timeout:       10,
		Level:         level,
		MaxTxBytes:    maxTxsBytes,
		GasAdjustment: gasAdjustment,
	})

	// import the key from privKey
	_, err := client.Key.Import(name, passwd, string(utils.GetPrivKeyArmor()))
	if err != nil {
		panic(err)
	}
}
