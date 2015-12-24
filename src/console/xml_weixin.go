package console

import (
	"encoding/xml"
	"fmt"
)

type WeixinMsgText struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int64
}

type WeixinMsgImage struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	PicUrl       string
	MediaId      string
	MsgId        int64
}

type WeixinMsgVoice struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string
	Format       string
	Recognition  string
	MsgId        int64
}

type WeixinMsgVideo struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string
	ThumbMediaId string
	MsgId        int64
}

type WeixinMsgShortVideo struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string
	ThumbMediaId string
	MsgId        int64
}

type WeixinMsgLocation struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Location_X   float64
	Location_Y   float64
	Scale        int
	Label        string
	MsgId        int64
}

type WeixinMsgLink struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Title        string
	Description  string
	Url          string
	MsgId        int64
}

//msg merge
type WeixinMsg struct {
	Root         xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MsgId        int64

	//text
	Content string

	MediaId string

	//image
	PicUrl string

	//voice
	Format      string
	Recognition string

	//video shortvideo
	ThumbMediaId string

	//location
	Location_X float64
	Location_Y float64
	Scale      int
	Label      string

	//lint
	Title       string
	Description string
	Url         string
}

func (this *WeixinMsg) ParseMsg(data []byte) {
	err := xml.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}

func (this *WeixinMsg) Print() {
	fmt.Println("ToUserName:", this.ToUserName)
	fmt.Println("FromUserName:", this.FromUserName)
	fmt.Println("CreateTime:", this.CreateTime)
	fmt.Println("MsgType:", this.MsgType)
	fmt.Println("MsgId:", this.MsgId)
	fmt.Println("Content:", this.Content)
	fmt.Println("MediaId:", this.MediaId)
	fmt.Println("PicUrl:", this.PicUrl)
	fmt.Println("Format:", this.Format)
	fmt.Println("Recognition:", this.Recognition)
	fmt.Println("ThumbMediaId:", this.ThumbMediaId)
	fmt.Println("Location_X:", this.Location_X)
	fmt.Println("Location_Y:", this.Location_Y)
	fmt.Println("Scale:", this.Scale)
	fmt.Println("Label:", this.Label)
	fmt.Println("Title:", this.Title)
	fmt.Println("Description:", this.Description)
	fmt.Println("Url:", this.Url)
}
