package api

import (
	"cschain-bond/types"
	"cschain-bond/utils"
	"encoding/json"
)

const (
	baseUrl = "http://10.1.4.248:3000"
)

// TODO common function for all denomId
func QueryNfts(denomId string) types.NftData {
	url := baseUrl + "/nfts?denomId=" + denomId
	bz := utils.GetFromUrl(url)

	var res types.Result
	if err := json.Unmarshal(bz, &res); err != nil {
		// TODO handle the error
		panic(err)
	}

	if len(res.Data.Data) <= 0 {
		// TODO handle the error
		panic("error")
	}
	return res.Data.Data[0]
}
