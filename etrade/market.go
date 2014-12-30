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
		AdjNonAdjFlag    bool    `json:"adjNonAdjFlag"`
		AnnualDividend   float64 `json:"annualDividend"`
		Ask              float64 `json:"ask"`
		AskExchange      string  `json:"askExchange"`
		AskSize          float64 `json:"askSize"`
		AskTime          string  `json:"askTime"`
		Beta             float64 `json:"beta"`
		Bid              float64 `json:"bid"`
		BidExchange      string  `json:"bidExchange"`
		BidSize          float64 `json:"bidSize"`
		BidTime          string  `json:"bidTime"`
		ChgClose         float64 `json:"chgClose"`
		ChgClosePrcn     float64 `json:"chgClosePrcn"`
		CompanyName      string  `json:"companyName"`
		DaysToExpiration float64 `json:"daysToExpiration"`
		DirLast          string  `json:"dirLast"`
		Dividend         float64 `json:"dividend"`
		Eps              float64 `json:"eps"`
		EstEarnings      float64 `json:"estEarnings"`
		ExDivDate        string  `json:"exDivDate"`
		ExchgLastTrade   string  `json:"exchgTradeLast"`
		Fsi              string  `json:"fsi"`
		High             float64 `json:"high"`
		High52           float64 `json:"high52"`
		HighAsk          float64 `json:"highAsk"`
		HighBid          float64 `json:"highBid"`
		LastTrade        float64 `json:"lastTrade"`
		Low              float64 `json:"low"`
		Low52            float64 `json:"low52"`
		LowAsk           float64 `json:"lowAsk"`
		LowBid           float64 `json:"lowBid"`
		NumTrades        float64 `json:"numTrades"`
		Open             float64 `json:"open"`
		OpenInterest     float64 `json:"openInterest"`
		OptionStyle      string  `json:"optionStyle"`
		OptionUnderlier  string  `json:"optionUnderlier"`
		PrevClose        float64 `json:"prevClose"`
		PrevDayVolume    float64 `json:"prevDayVolume"`
		PrimaryExchange  string  `json:"primaryExchange"`
		SymbolDesc       string  `json:"symbolDesc"`
		TodayClosed      float64 `json:"todayClosed"`
		TotalVolume      float64 `json:"totalVolume"`
		Upc              float64 `json:"upc"`
		Volume10Day      float64 `json:"volume10Day"`
	} `json:"all"`
	Product  Product  `json:"product"`
	DateTime DateTime `json:"dateTime"`
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
