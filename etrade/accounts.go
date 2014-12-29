package etrade

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Account module code

type AccountDetail struct {
	AccountDesc      string      `json:"accountDesc"`
	AccountId        int         `json:"accountId"`
	MarginLevel      string      `json:"marginLevel"`
	NetAccountValue  interface{} `json:"netAccountValue"`
	RegistrationType string      `json:"registrationType"`
}

func GetAccountDetails(c *OauthClient) (*[]AccountDetail, error) {
	a := &[]AccountDetail{}
	resp, err := c.Get(URL_ACCOUNTLIST+RESPONSE_FORMAT, map[string]string{}, &c.Config.AccessToken)
	if err != nil {
		resp.Body.Close()
		return a, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", data)
	var tempAcctList struct {
		JSONListResponse struct {
			Response []AccountDetail `json:"response"`
		} `json:"json.accountListResponse"`
	}

	err = json.Unmarshal(data, &tempAcctList)
	if err != nil {
		return a, err
	}
	a = &tempAcctList.JSONListResponse.Response
	return a, nil
}
