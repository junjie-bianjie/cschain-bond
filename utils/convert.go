package utils

import (
	"cschain-bond/types"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func BzToBonds(bz []byte) types.Bonds {
	var bonds types.Bonds
	err := json.Unmarshal(bz, &bonds)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshalJson: %v\n", err)
		os.Exit(1)
	}

	return bonds
}

func ParseToResult(bonds types.Bonds, tokenData *types.TokenData) {
	// init the result of 深交所
	tokenData.Visible = true
	tokenData.Report.Header = []string{"成交金额", "债券类别", "回购类别"}
	tokenData.Report.FixedValueHeader = types.FixedValueHeader{
		Header: "市场",
		Value:  "深交所",
	}

	// parse bonds to result
	tokenData.Report.Data = make([][]string, 0)
	d := &tokenData.Report.Data
	var isBond bool
	var yearAndMoth string
	for _, bond := range bonds {
		yearAndMoth = bond.Metadata.Subname
		for _, data := range bond.Data {
			if data.Lbmc == "ABS" {
				break
			}

			if data.Lbmc == "债券现货" {
				isBond = true
				continue
			}

			if data.Lbmc == "债券回购" {
				isBond = false
				continue
			}

			formatLbmc := strings.ReplaceAll(data.Lbmc, "&nbsp;&nbsp;", "")
			formatCjje := strings.ReplaceAll(data.Cjje, ",", "")
			if isBond {
				*d = append(*d, []string{formatCjje, formatLbmc, ""})
			} else {
				*d = append(*d, []string{formatCjje, "", formatLbmc})
			}
		}
	}

	// assign the date
	if len(yearAndMoth) > 0 {
		startDate, endDate := GetMonthStartAndEnd(yearAndMoth)
		tokenData.Report.Date = types.Date{
			StartDate: startDate,
			EndDate:   endDate,
			Period:    "M",
		}
	}
}
