package etrade

const (
	ETRADE_OAUTH_SERVER = "https://etws.etrade.com"
	AUTHORIZE_URL       = "https://us.etrade.com/e/t/etws/authorize"
	ETRADE_SERVER       = "https://etwssandbox.etrade.com"
	RESPONSE_FORMAT     = ".json"
	REQUEST_FORMAT      = ".json" // It can be "json" or "xml", default is xml+
	DEBUG_MODE          = 1

	REQUEST_TOKEN_URL    = ETRADE_OAUTH_SERVER + "/oauth/request_token"
	ACCESS_TOKEN_URL     = ETRADE_OAUTH_SERVER + "/oauth/access_token"
	RENEW_TOKEN_URL      = ETRADE_OAUTH_SERVER + "/oauth/renew_access_token"
	REVOKE_TOKEN_URL     = ETRADE_OAUTH_SERVER + "/oauth/revoke_access_token"
	url_str_part         = "sandbox/"
	URL_ACCOUNTLIST      = ETRADE_SERVER + "/accounts/" + url_str_part + "rest/accountlist"
	URL_ACCOUNTBALANCE   = ETRADE_SERVER + "/accounts/" + url_str_part + "rest/accountbalance"
	URL_ACCOUNTPOSITIONS = ETRADE_SERVER + "/accounts/" + url_str_part + "rest/accountpositions"
	URL_ACCOUNTALERTS    = ETRADE_SERVER + "/accounts/" + url_str_part + "rest/alerts"

	URL_OPTIONCHAINS  = ETRADE_SERVER + "/market/" + url_str_part + "rest/optionchains"
	URL_MARKETINDICES = ETRADE_SERVER + "/market/" + url_str_part + "rest/marketindices"
	URL_PRODUCTLOOKUP = ETRADE_SERVER + "/market/" + url_str_part + "rest/productlookup"
	URL_GETQUOTE      = ETRADE_SERVER + "/market/" + url_str_part + "rest/quote"
	URL_EXPIRYDATES   = ETRADE_SERVER + "/market/" + url_str_part + "rest/optionexpiredate"

	URL_ORDERLIST      = ETRADE_SERVER + "/order/" + url_str_part + "rest/orderlist"
	URL_PL_EQ_ORDER    = ETRADE_SERVER + "/order/" + url_str_part + "rest/placeequityorder"
	URL_PL_OP_ORDER    = ETRADE_SERVER + "/order/" + url_str_part + "rest/placeoptionorder"
	URL_PR_EQ_ORDER    = ETRADE_SERVER + "/order/" + url_str_part + "rest/previewequityorder"
	URL_PR_OP_ORDER    = ETRADE_SERVER + "/order/" + url_str_part + "rest/previewoptionorder"
	URL_PR_CH_EQ_ORDER = ETRADE_SERVER + "/order/" + url_str_part + "rest/previewchangeequityorder"
	URL_PL_CH_EQ_ORDER = ETRADE_SERVER + "/order/" + url_str_part + "rest/placechangeequityorder"
	URL_PR_CH_OP_ORDER = ETRADE_SERVER + "/order/" + url_str_part + "rest/previewchangeoptionorder"
	URL_PL_CH_OP_ORDER = ETRADE_SERVER + "/order/" + url_str_part + "rest/placechangeoptionorder"
	URL_CANCEL_ORDER   = ETRADE_SERVER + "/order/" + url_str_part + "rest/cancelorder"
)
