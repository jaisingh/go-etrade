package etrade

// Account module code

// Account details for the list of accounts
type AccountDetail struct {
	AccountDesc      string      `json:"accountDesc"`
	AccountId        int         `json:"accountId"`
	MarginLevel      string      `json:"marginLevel"`
	NetAccountValue  interface{} `json:"netAccountValue"`
	RegistrationType string      `json:"registrationType"`
}

// Get the list of account from etrade, return an array of AccountDetails
func (c *OauthClient) GetAccountList() (*[]AccountDetail, error) {
	a := []AccountDetail{}
	var tempAcctList struct {
		JSONListResponse struct {
			Response []AccountDetail `json:"response"`
		} `json:"json.accountListResponse"`
	}

	err := c.GetUnmarshal(URL_ACCOUNTLIST+RESPONSE_FORMAT, map[string]string{}, &tempAcctList)

	if err != nil {
		return &a, err
	}
	a = tempAcctList.JSONListResponse.Response
	return &a, nil
}
