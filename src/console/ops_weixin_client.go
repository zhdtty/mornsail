package console

import ()

var ADDRESS string = "api.weixin.qq.com"
var APP_ID string = "wx19e530f023d3233c"
var APP_SECRET = "d4624c36b6795d1d99dcf0547af5443d"
var ACCESS_TOKEN string = ""
var OPS_TOKEN string = "/cgi-bin/token?"

func DoToken(req WeixinTokenRequest) []byte {
        req.AppId = APP_ID
	req.AppSecret = APP_SECRET
	req.GrantType = "client_credential"
        params := req.ToRequest()
        result := SendHttpRequest(ADDRESS, OPS_TOKEN, "GET", params, true)
	var resp WeixinTokenResponse
	resp.ParseJson(result)
	ACCESS_TOKEN = resp.Access_Token
	return []byte(resp.Access_Token)
}
