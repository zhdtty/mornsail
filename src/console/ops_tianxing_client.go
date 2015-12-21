package console

import (
	"fmt"
	"strconv"
	"strings"
	"util"
)

var TIANXING_ADDRESS string = "api.huceo.com"
var TIANXING_API_KEY = "f1494a448f58c7df804e1c1c540d56a2"

var TIANXING_QUERY_WXNEWS string = "/wxnew/other?"
var TIANXING_QUERY_SOCIAL string = "/social/other?"
var TIANXING_QUERY_GUONEI string = "/guonei/other?"
var TIANXING_QUERY_WORLD string = "/world/other?"
var TIANXING_QUERY_TIYU string = "/tiyu/other?"
var TIANXING_QUERY_HUABIAN string = "/huabian/other?"
var TIANXING_QUERY_MEINV string = "/meinv/other?"
var TIANXING_QUERY_KEJI string = "/keji/other?"
var TIANXING_QUERY_QIWEN string = "/qiwen/other?"
var TIANXING_QUERY_HEALTH string = "/health/other?"
var TIANXING_QUERY_TRAVEL string = "/travel/other?"
var TIANXING_QUERY_APPLE string = "/apple?"

const (
	QUERY_WXNEWS  = 1 + iota
	QUERY_SOCIAL  //2
	QUERY_GUONEI  //3
	QUERY_WORLD   //4
	QUERY_TIYU    //5
	QUERY_HUABIAN //6
	QUERY_MEINV   //7
	QUERY_KEJI    //8
	QUERY_QIWEN   //9
	QUERY_HEALTH  //10
	QUERY_TRAVEL  //11
	QUERY_APPLE   //12
	QUERY_COUNT
)

var TIANXING_QUERY_STR [QUERY_COUNT]string = [...]string{
	string(""),
	TIANXING_QUERY_WXNEWS,
	TIANXING_QUERY_SOCIAL,
	TIANXING_QUERY_GUONEI,
	TIANXING_QUERY_WORLD,
	TIANXING_QUERY_TIYU,
	TIANXING_QUERY_HUABIAN,
	TIANXING_QUERY_MEINV,
	TIANXING_QUERY_KEJI,
	TIANXING_QUERY_QIWEN,
	TIANXING_QUERY_HEALTH,
	TIANXING_QUERY_TRAVEL,
	TIANXING_QUERY_APPLE,
}

func DoTianxingQuery(queryId int32, req TianxingRequest) (*TianxingResponse, error) {
	if queryId < QUERY_WXNEWS || queryId >= QUERY_COUNT {
		return nil, fmt.Errorf("Invalid query id, id : %d", queryId)
	}
	req.Key = TIANXING_API_KEY
	params := req.ToRequest()
	result := SendHttpRequest(TIANXING_ADDRESS, TIANXING_QUERY_STR[queryId], "GET", params, false)
	resp := &TianxingResponse{}
	resp.ParseJson(result)
	return resp, nil
}

func tianxingMuxResponse(reqMsg *WeixinMsg, resp *TianxingResponse) []byte {
	if resp.Code != 200 {
		var respMsg WeixinMsgTextResponse
		respMsg.ToUserName = reqMsg.FromUserName
		respMsg.FromUserName = reqMsg.ToUserName
		respMsg.CreateTime = int(util.GetCurrentSecond())
		respMsg.MsgType = "text"
		respMsg.Content = resp.Msg
		return SerialToXML(respMsg)
	}
	var respMsg WeixinMsgNewsResponse
	respMsg.ToUserName = reqMsg.FromUserName
	respMsg.FromUserName = reqMsg.ToUserName
	respMsg.CreateTime = int(util.GetCurrentSecond())
	respMsg.MsgType = "news"
	cnt := 0
	for _, v := range resp.NewsList {
		var wxnew WeixinNews
		wxnew.Title = v.Title
		wxnew.Description = v.Description
		wxnew.PicUrl = v.PicUrl
		wxnew.Url = v.Url
		respMsg.Items = append(respMsg.Items, wxnew)
		cnt++
		if cnt >= WEIXIN_MAX_NEWS_COUNT {
			break
		}
	}
	respMsg.ArticleCount = cnt
	return SerialToXML(respMsg)
}

func TianxingHttpRequest(reqMsg *WeixinMsg) []byte {
	if reqMsg.MsgType != "text" {
		return []byte("")
	}
	var queryId int32
	queryId = QUERY_WXNEWS

	content := reqMsg.Content
	content = strings.TrimPrefix(content, "@")
	ary := strings.Split(content, "-")
	aryLen := len(ary)
	if aryLen <= 0 {
		var txReq TianxingRequest
		txReq.Num = "5"
		txReq.Rand = "1"
		txResp, _ := DoTianxingQuery(queryId, txReq)
		rs := tianxingMuxResponse(reqMsg, txResp)
		return rs
	}

	qId, err := strconv.Atoi(ary[0])
	if err != nil {
		fmt.Println("字符串转换成整数失败")
		queryId = QUERY_WXNEWS

		var txReq TianxingRequest
		txReq.Num = "5"
		txReq.Rand = "1"
		txReq.Word = content
		txResp, _ := DoTianxingQuery(queryId, txReq)
		rs := tianxingMuxResponse(reqMsg, txResp)
		return rs
	}
	queryId = int32(qId)
	if queryId < QUERY_WXNEWS || queryId >= QUERY_COUNT {
		queryId = QUERY_WXNEWS
	}
	var txReq TianxingRequest
	txReq.Num = "5"
	txReq.Rand = "1"
	if aryLen == 2 {
		txReq.Word = ary[1]
	}
	txResp, _ := DoTianxingQuery(queryId, txReq)
	rs := tianxingMuxResponse(reqMsg, txResp)
	return rs
}
