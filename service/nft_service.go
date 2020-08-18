package service

import (
	"cschain-bond/api"
	"cschain-bond/dao"
	"cschain-bond/logger"
	"cschain-bond/models"
	"cschain-bond/types"
	"cschain-bond/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func UploadByNFT() {
	bz, err := utils.GetFromUrl("http://www.szse.cn/api/report/ShowReport/data?SHOWTYPE=JSON&CATALOGID=scsj_zqydgk&TABKEY=tab1&txtQueryDate=2020-08")
	if err != nil {
		logger.Error(fmt.Sprintf("fetch: %v\n", err))
		return
	}

	bonds, err := utils.BzToBonds(bz)
	if err != nil {
		logger.Error(fmt.Sprintf("unmarshalJson: %v\n", err))
		return
	}

	var tokenData types.TokenData
	utils.Parse2TokenData(bonds, &tokenData)
	tokenDataBz, err := json.Marshal(tokenData)
	if err != nil {
		logger.Error(fmt.Sprintf("unmarshalJson: %v\n", err))
		return
	}

	// use sdk to IssueDenom, MintNFT
	denomID := strings.ToLower(utils.RandStringOfLength(4))
	denomName := strings.ToLower(utils.RandStringOfLength(4))
	schema := utils.GetScheme()
	if err = issueDenom(denomID, denomName, schema); err != nil {
		logger.Error(fmt.Sprintf("issueDenom failed: %v\n", err))
		return
	}

	tokenID := strings.ToLower(utils.RandStringOfLength(7))
	tokenName := strings.ToLower(utils.RandStringOfLength(7))
	if err = mintNFT(denomID, tokenID, tokenName, string(tokenDataBz)); err != nil {
		logger.Error(fmt.Sprintf("mintNFT failed: %v\n", err))
		return
	}
}

func DataCollation(denomId string) {
	// queryNft by restApi
	nftData := api.QueryNfts(denomId)

	nameIdMap := bondRepurchaseMarket2Map()
	txs := make([]models.BondTransaction, 0)

	// construct ever row of data, then push ever data in slice
	nftId := nftData.NftId
	owner := nftData.Owner
	tokenUri := nftData.TokenUri

	var tokenData types.TokenData
	if err := json.Unmarshal([]byte(nftData.TokenDataStr), &tokenData); err != nil {
		logger.Error(fmt.Sprintf("unmarshalJson: %v\n", err))
		return
	}

	visible := tokenData.Visible
	market := tokenData.Report.FixedValueHeader.Value
	date := tokenData.Report.Date
	for _, data := range tokenData.Report.Data {
		var amount float64
		if len(data[0]) > 0 {
			var err error
			amount, err = strconv.ParseFloat(data[0], 64)
			if err != nil {
				logger.Error("parseFloat failed", logger.String("err", err.Error()))
				return
			}
		}

		// TODO out_bond exception of data
		// get bondCategoryName or repurchaseCategory from data ex[]
		var bondCategoryName, repurchaseCategory string
		if len(data[1]) != 0 {
			bondCategoryName = data[1]
		}

		if len(data[2]) != 0 {
			repurchaseCategory = data[2]
		}

		tx := models.BondTransaction{
			NftId: nftId,
			// TODO SourceType ,Visible when demand is determined
			SourceType:         0,
			DenomId:            denomId,
			Owner:              owner,
			Uri:                tokenUri,
			Visible:            visible,
			Amount:             amount,
			Market:             nameIdMap[market],
			StartDate:          utils.String2Time(date.StartDate),
			EndDate:            utils.String2Time(date.EndDate),
			PeriodCategory:     date.Period,
			BondCategory:       nameIdMap[bondCategoryName],
			RepurchaseCategory: nameIdMap[repurchaseCategory],
		}
		txs = append(txs, tx)
	}

	BatchInsert(txs)
}

// bondAndRepurchase2Map convert the result of db(3 tables) to map
// ex[{1 国债 1 3} {2 地方政府债 0 2001} {3 政策性金融债 2 5}] {2 深交所} -> [国债:1,地方政府债:2,政策性金融债:3,深交所:2]
func bondRepurchaseMarket2Map() map[string]int {
	var res = make(map[string]int)

	var b dao.BondVarietyDao
	for _, v := range b.FindAll() {
		res[v.Name] = v.ID
	}

	var r dao.RepurchaseVarietyDao
	for _, v := range r.FindAll() {
		res[v.Name] = v.ID
	}

	var m dao.MarketVarietyDao
	for _, v := range m.FindAll() {
		res[v.Market] = v.Id
	}
	return res
}

// BatchInsert batch insert into table bond_transaction, using placeholder(?) to prevent sql injection
func BatchInsert(bondTxs []models.BondTransaction) {
	db := utils.GetConnection()
	defer db.Close()

	tx := db.Begin()
	sql := `INSERT INTO bond_transaction(
					nft_id,source_type,denom_id,owner,uri,
					visible,amount,market,start_date,
					end_date,period_category,bond_category,repurchase_category) 
			VALUES`
	var vals []interface{}
	const rowSql = "(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var inserts []string

	for _, bondTx := range bondTxs {
		inserts = append(inserts, rowSql)
		vals = append(vals,
			bondTx.NftId, bondTx.SourceType, bondTx.DenomId, bondTx.Owner, bondTx.Uri,
			bondTx.Visible, bondTx.Amount, bondTx.Market, bondTx.StartDate,
			bondTx.EndDate, bondTx.PeriodCategory, bondTx.BondCategory, bondTx.RepurchaseCategory)
	}

	sql = sql + strings.Join(inserts, ",")
	err := tx.Exec(sql, vals...).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	tx.Commit()
}
