package console

import (
	"fmt"
)

var ADDRESS string = "api.weixin.qq.com"
var FILE_ADDRESS string = "file.api.weixin.qq.com"

//var APP_ID string = "wx19e530f023d3233c"            //for test
//var APP_SECRET = "d4624c36b6795d1d99dcf0547af5443d" //for test
var APP_ID string = "wx41b36bc973def2b9"
var APP_SECRET = "ee96b7fc2c33ff8999f6d1bcbd4da0af"
var ACCESS_TOKEN string = ""

var WEIXIN_LOAD_TOKEN string = "/cgi-bin/token?"
var WEIXIN_LOAD_MEDIA string = "/cgi-bin/media/get?"
var WEIXIN_CUSTOM_MESSAGE = "/cgi-bin/message/custom/send?"

func DoToken(req WeixinTokenRequest) []byte {
	req.AppId = APP_ID
	req.AppSecret = APP_SECRET
	req.GrantType = "client_credential"
	params := req.ToRequest()
	result := SendHttpRequest(ADDRESS, WEIXIN_LOAD_TOKEN, "GET", params, true)
	var resp WeixinTokenResponse
	resp.ParseJson(result)
	ACCESS_TOKEN = resp.Access_Token
	fmt.Println("weixin access token :", ACCESS_TOKEN)
	return []byte(resp.Access_Token)
}

func DoLoadMedia(req WeixinMediaLoadRequest) []byte {
	req.AccessToken = ACCESS_TOKEN
	params := req.ToRequest()
	return SendHttpRequestNotJson(FILE_ADDRESS, WEIXIN_LOAD_MEDIA, "GET", params, false)
}
