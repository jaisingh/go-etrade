package etrade

// Market module

// Look Up Product

type Product struct {
	CompanyName  string `json:"companyName"`
	Exchange     string `json:"exchange"`
	SecurityType string `json:"securityType"`
	Symbol       string `json:"symbol"`
}

func (c *OauthClient) ProductLookup(company string, kind string) (*[]Product, error) {
	var productLookupResponse struct {
		ProductLookupResponse struct {
			ProductList []Product `json:"productList"`
		} `json:"productLookupResponse"`
	}

	var output []Product

	resp, err := c.Get(
		URL_PRODUCTLOOKUP+RESPONSE_FORMAT,
		map[string]string{
			"company": company,
			"type":    kind,
		},
		&c.Config.AccessToken)
	if err != nil {
		resp.Body.Close()
		return &output, err
	}
	defer resp.Body.Close()

	err = unmarshalResponse(resp, &productLookupResponse)
	output = productLookupResponse.ProductLookupResponse.ProductList
	return &output, nil
}

// Get Quote

func (c *OauthClient) GetQuote(symbol string, detailFlag ...string) {

}
