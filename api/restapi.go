package api

import (
	"cschain-bond/logger"
	"cschain-bond/types"
	"cschain-bond/utils"
	"encoding/json"
	"fmt"
)

const (
	baseUrl = "http://10.1.4.248:3000"
)

// QueryNfts common function for all denomId
func QueryNfts(denomId string) types.NftData {
	url := baseUrl + "/nfts?denomId=" + denomId
	bz, err := utils.GetFromUrl(url)
	if err != nil {
		logger.Error(fmt.Sprintf("fetch: %v\n", err))
		return types.NftData{}

	}

	var res types.Result
	if err := json.Unmarshal(bz, &res); err != nil {
		logger.Error(fmt.Sprintf("unmarshalJson: %v\n", err))
		return types.NftData{}
	}

	if len(res.Data.Data) == 0 {
		logger.Error("the result of queryNfts is nil")
		return types.NftData{}
	}
	return res.Data.Data[0]
}
