package etrade

// Market module

// Look Up Product

type Product struct {
	CompanyName  string `json:"companyName"`
	Exchange     string `json:"exchange"`
	SecurityType string `json:"type"`
	Symbol       string `json:"symbol"`
}

func (c *OauthClient) ProductLookup(company string, kind string) (*[]Product, error) {
	var productLookupResponse struct {
		ProductLookupResponse struct {
			ProductList []Product `json:"productList"`
		} `json:"productLookupResponse"`
	}

	var output []Product
	url := URL_PRODUCTLOOKUP + RESPONSE_FORMAT
	opts := map[string]string{
		"company": company,
		"type":    kind,
	}

	err := c.GetUnmarshal(url, opts, &productLookupResponse)
	if err != nil {
		return &output, err
	}
	output = productLookupResponse.ProductLookupResponse.ProductList
	return &output, nil
}

// Get Quote

type QuoteData struct {
	All struct {
		AdjNonAdjFlag    bool        `json:"adjNonAdjFlag"`
		AnnualDividend   interface{} `json:"annualDividend"`
		Ask              interface{} `json:"ask"`
		AskExchange      string      `json:"askExchange"`
		AskSize          interface{} `json:"askSize"`
		AskTime          string      `json:"askTime"`
		Beta             interface{} `json:"beta"`
		Bid              interface{} `json:"bid"`
		BidExchange      string      `json:"bidExchange"`
		BidSize          interface{} `json:"bidSize"`
		BidTime          string      `json:"bidTime"`
		ChgClose         interface{} `json:"chgClose"`
		ChgClosePrcn     interface{} `json:"chgClosePrcn"`
		CompanyName      string      `json:"companyName"`
		DaysToExpiration interface{} `json:"daysToExpiration"`
		DirLast          string      `json:"dirLast"`
		Dividend         interface{} `json:"dividend"`
		Eps              interface{} `json:"eps"`
		EstEarnings      interface{} `json:"estEarnings"`
		ExDivDate        string      `json:"exDivDate"`
		ExchgLastTrade   string      `json:"exchgTradeLast"`
		Fsi              string      `json:"fsi"`
		High             interface{} `json:"high"`
		High52           interface{} `json:"high52"`
		HighAsk          interface{} `json:"highAsk"`
		HighBid          interface{} `json:"highBid"`
		LastTrade        interface{} `json:"lastTrade"`
		Low              interface{} `json:"low"`
		Low52            interface{} `json:"low52"`
		LowAsk           interface{} `json:"lowAsk"`
		LowBid           interface{} `json:"lowBid"`
		NumTrades        interface{} `json:"numTrades"`
		Open             interface{} `json:"open"`
		OpenInterest     interface{} `json:"openInterest"`
		OptionStyle      string      `json:"optionStyle"`
		OptionUnderlier  string      `json:"optionUnderlier"`
		PrevClose        interface{} `json:"prevClose"`
		PrevDayVolume    interface{} `json:"prevDayVolume"`
		PrimaryExchange  string      `json:"primaryExchange"`
		SymbolDesc       string      `json:"symbolDesc"`
		TodayClosed      interface{} `json:"todayClosed"`
		TotalVolume      interface{} `json:"totalVolume"`
		Upc              interface{} `json:"upc"`
		Volume10Day      interface{} `json:"volume10Day"`
	} `json:"all"`
	Product  Product `json:"product"`
	DateTime string  `json:"dateTime"`
}

func (c *OauthClient) GetQuote(symbol string, detailFlag ...string) (*QuoteData, error) {
	q := QuoteData{}
	url := URL_GETQUOTE + "/" + symbol + RESPONSE_FORMAT
	opts := map[string]string{}
	if detailFlag != nil {
		opts["detailFlag"] = detailFlag[0]
	} else {
		opts["detailFlag"] = "ALL"
	}

	var quoteResponse struct {
		QuoteResponse struct {
			Data QuoteData `json:"quoteData"`
		} `json:"quoteResponse"`
	}

	err := c.GetUnmarshal(url, opts, &quoteResponse)
	if err != nil {
		return &q, err
	}
	q = quoteResponse.QuoteResponse.Data

	return &q, nil

}
