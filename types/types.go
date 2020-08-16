package types

type Bonds []struct {
	Data     []Data `json:"data"`
	Metadata struct {
		Subname string `json:"subname"`
	} `json:"metadata"`
}

type Data struct {
	Cjje string `json:"cjje"`
	Cjl  string `json:"cjl"`
	Lbmc string `json:"lbmc"`
}

type TokenData struct {
	Visible bool `json:"visible."`
	Report  struct {
		Header           []string         `json:"header"`
		Data             [][]string       `json:"data"`
		FixedValueHeader FixedValueHeader `json:"fixed_value_header"`
		Date             Date             `json:"date"`
	} `json:"report"`
}

type Date struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Period    string `json:"period"`
}

type FixedValueHeader struct {
	Header string `json:"header"`
	Value  string `json:"value"`
}

type Result struct {
	Code int         `json:"code"`
	Data OutSideData `json:"data"`
}

type OutSideData struct {
	Data []NftData `json:"data"`
}

type NftData struct {
	DenomId      string `json:"denom_id"`
	NftId        string `json:"nft_id"`
	Owner        string `json:"owner"`
	TokenUri     string `json:"tokenUri"`
	TokenDataStr string `json:"tokenData"`
}
