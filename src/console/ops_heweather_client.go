package console

import (
	"fmt"
	//	"strconv"
	"strings"
	"util"
)

//interface source : http://www.heweather.com/

//https://api.heweather.com/x3/weather?cityid=城市ID&key=你的认证key
//@Note: 参数，均为可选
//city城市名称、支持中英文,不区分大小写和空格,城市和国家之间用英文逗号分割北京、beijing、london,united kingdom
//cityid城市ID,参见 国内城市ID列表CN101010100
//cityip城市IP,输入所在城市的任意IP123.45.67.8
//key用户认证key

var HEWEATHER_ADDRESS string = "api.heweather.com"
var HEWEATHER_API_KEY = "ba265d5e20f24b459d234b8b4130bd9b"

var HEWEATHER_QUERY string = "/x3/weather?"

func DoHeweatherQuery(req HeweatherRequest) *HeweatherResponse {
	req.Key = HEWEATHER_API_KEY
	params := req.ToRequest()
	result := SendHttpRequest(HEWEATHER_ADDRESS, HEWEATHER_QUERY, "GET", params, false)
	fmt.Println(string(result))
	resp := &HeweatherResponse{}
	resp.ParseJson(result)
	return resp
}

func HeweatherHttpRequest(reqMsg *WeixinMsg) []byte {
	if reqMsg.MsgType != "text" {
		return []byte("")
	}
	content := reqMsg.Content
	content = strings.TrimSuffix(content, "天气")
	content = strings.TrimSuffix(content, "天气！")
	var hwReq HeweatherRequest
	hwReq.City = content
	resp := DoHeweatherQuery(hwReq)
	if len(resp.Data) > 0 && resp.Data[0].Status == "ok" {
		var respMsg WeixinMsgTextResponse
		respMsg.ToUserName = reqMsg.FromUserName
		respMsg.FromUserName = reqMsg.ToUserName
		respMsg.CreateTime = int(util.GetCurrentSecond())
		respMsg.MsgType = "text"
		respMsg.Content = resp.ToString()
		if respMsg.Content == "" {
			return []byte("")
		}
		return SerialToXML(respMsg)
	} else {
		return []byte("")
	}
}
