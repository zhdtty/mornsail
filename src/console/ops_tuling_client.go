package console

import (
	"fmt"
	"util"
)

var TULING_ADDRESS string = "www.tuling123.com"
var TULING_API_KEY = "58e4a76358d33acddb3014060ffbd952"

var TULING_QUERY string = "/openapi/api?"

const (
	CODE_TEXT  int32 = 100000
	CODE_LINK  int32 = 200000
	CODE_NEWS  int32 = 302000
	CODE_TRAIN int32 = 305000
	CODE_AIR   int32 = 306000
	CODE_FOOD  int32 = 308000
)

func DoTulingQuery(req TulingRequest) *TulingResponse {
	req.Key = TULING_API_KEY
	params := req.ToRequest()
	result := SendHttpRequest(TULING_ADDRESS, TULING_QUERY, "GET", params, false)
	resp := &TulingResponse{}
	resp.ParseJson(result)
	return resp
}

func tulingMuxResponse(reqMsg *WeixinMsg, resp *TulingResponse) []byte {
	if resp.Code == CODE_TEXT {
		var respMsg WeixinMsgTextResponse
		respMsg.ToUserName = reqMsg.FromUserName
		respMsg.FromUserName = reqMsg.ToUserName
		respMsg.CreateTime = int(util.GetCurrentSecond())
		respMsg.MsgType = "text"
		respMsg.Content = resp.Text
		return SerialToXML(respMsg)
	}

	var respMsg WeixinMsgNewsResponse
	respMsg.ToUserName = reqMsg.FromUserName
	respMsg.FromUserName = reqMsg.ToUserName
	respMsg.CreateTime = int(util.GetCurrentSecond())
	respMsg.MsgType = "news"
	cnt := 0
	if resp.Code == CODE_LINK {
		fmt.Println("code:", resp.Code)
		var wxnew WeixinNews
		wxnew.Title = resp.Text
		wxnew.Url = resp.Url
		respMsg.Items = append(respMsg.Items, wxnew)
		cnt++
	} else if resp.Code == CODE_NEWS {
		fmt.Println("code:", resp.Code)
		for _, v := range resp.List {
			var wxnew WeixinNews
			wxnew.Title = v.Article
			wxnew.Description = v.Source
			wxnew.PicUrl = v.Icon
			wxnew.Url = v.DetailUrl
			respMsg.Items = append(respMsg.Items, wxnew)
			cnt++
			if cnt >= WEIXIN_MAX_NEWS_COUNT {
				break
			}
		}
	} else if resp.Code == CODE_TRAIN {
		fmt.Println("code:", resp.Code)
		for _, v := range resp.List {
			var wxnew WeixinNews
			desc := v.Start + "-" + v.Terminal + " " + v.TrainNum + " " + v.StartTime + "~" + v.EndTime
			wxnew.Title = desc
			wxnew.Description = desc
			wxnew.PicUrl = v.Icon
			wxnew.Url = v.DetailUrl
			respMsg.Items = append(respMsg.Items, wxnew)
			cnt++
			if cnt >= WEIXIN_MAX_NEWS_COUNT {
				break
			}
		}
	} else if resp.Code == CODE_AIR {
		fmt.Println("code:", resp.Code)
		for _, v := range resp.List {
			var wxnew WeixinNews
			desc := v.Flight + " " + v.StartTime + "~" + v.EndTime
			wxnew.Title = desc
			wxnew.Description = desc
			wxnew.PicUrl = v.Icon
			respMsg.Items = append(respMsg.Items, wxnew)
			cnt++
			if cnt >= WEIXIN_MAX_NEWS_COUNT {
				break
			}
		}
	} else if resp.Code == CODE_FOOD {
		fmt.Println("code:", resp.Code)
		for _, v := range resp.List {
			var wxnew WeixinNews
			wxnew.Title = v.Name + "\r\n" + v.Info
			wxnew.Description = v.Info
			wxnew.PicUrl = v.Icon
			wxnew.Url = v.DetailUrl
			respMsg.Items = append(respMsg.Items, wxnew)
			cnt++
			if cnt >= WEIXIN_MAX_NEWS_COUNT {
				break
			}
		}
	}
	respMsg.ArticleCount = cnt
	return SerialToXML(respMsg)
}

func TulingHttpRequest(reqMsg *WeixinMsg) []byte {
	if reqMsg.MsgType != "text" {
		return []byte("")
	}
	var tlReq TulingRequest
	tlReq.Info = reqMsg.Content
	tlResp := DoTulingQuery(tlReq)
	rs := tulingMuxResponse(reqMsg, tlResp)
	return rs
}
