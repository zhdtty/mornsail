package console

import (
	"encoding/xml"
	"fmt"
)

type WeixinMsgTextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
}

type WeixinMsgImageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string `xml:"Image>MediaId"`
}

type WeixinMsgVoiceResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string `xml:"Voice>MediaId"`
}

type WeixinMsgVideoResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string `xml:"Video>MediaId"`
	Title        string `xml:"Video>Title"`
	Description  string `xml:"Video>Description"`
	MsgId        int64
}

type WeixinMsgMusicResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Title        string `xml:"Music>Title"`
	Description  string `xml:"Music>Description"`
	MusicUrl     string `xml:"Music>MusicUrl"`
	HQMusicUrl   string `xml:"Music>HQMusicUrl"`
	ThumbMediaId string `xml:"Music>ThumbMediaId"`
	MsgId        int64
}

type WeixinNews struct {
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type WeixinMsgNewsResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	ArticleCount int
	Items        []WeixinNews `xml:"Articles>item"`
}

func SerialToXML(v interface{}) []byte {
	output, err := xml.MarshalIndent(v, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output
}
