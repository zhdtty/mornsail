package console

import (
	"fmt"
	"github.com/coocood/jas"
	"io"
	"reflect"
	"sort"
	//	"util"
	"strings"
)

var WEIXIN_TOKEN string = "test"
var WEIXIN_MAX_NEWS_COUNT int = 10

type Weixin struct{}

func (*Weixin) Get(ctx *jas.Context) {
	fmt.Println("Get weixin")
	echostr := ctx.RequireString("echostr")
	nonce := ctx.RequireString("nonce")
	timestamp := ctx.RequireString("timestamp")
	sig := ctx.RequireString("signature")

	params := []string{nonce, timestamp, WEIXIN_TOKEN}
	sort.Sort(sort.StringSlice(params))
	data := ""
	for _, v := range params {
		data += v
	}
	sha1Sig := string(sha1s(data))

	ctx.Data = echostr
	if sig == sha1Sig {
		cfg := &jas.Config{}
		cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
			len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
			return len
		}
		ctx.SetConfig(cfg)
	}
}

func (*Weixin) Post(ctx *jas.Context) {
	fmt.Println("Post weixin msg")
	b := make([]byte, 2048)
	_, _ = ctx.Body.Read(b)
	fmt.Println(string(b))
	msg := &WeixinMsg{}
	msg.ParseMsg(b)
	msg.Print()

	switch msg.MsgType {
	case "text":
		var result string
		if len(msg.Content) > 0 {
			if strings.HasPrefix(msg.Content, "搜索") {
				result = string(TianxingHttpRequest(msg))
			} else if strings.HasSuffix(msg.Content, "天气") || strings.HasSuffix(msg.Content, "天气！") {
				result = string(HeweatherHttpRequest(msg))
				if result == "" {
					result = string(TulingHttpRequest(msg))
				}
			} else {
				result = string(TulingHttpRequest(msg))
			}
		}
		ctx.Data = result
		fmt.Println(result)
		cfg := &jas.Config{}
		cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
			len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
			return len
		}
		ctx.SetConfig(cfg)
	case "image":
		{
		}
	case "voice":
		//		text := BaiduVoiceHttpRequest(msg) //baidu http voice
		text := msg.Recognition

		fmt.Println("recognition:", text)
		if text == "" {
			ctx.Data = "success"
		} else {
			msg.MsgType = "text"
			msg.Content = text
			var result string
			if len(msg.Content) > 0 {
				if strings.HasPrefix(msg.Content, "搜索") {
					result = string(TianxingHttpRequest(msg))
				} else if strings.HasSuffix(msg.Content, "天气") || strings.HasSuffix(msg.Content, "天气！") {
					result = string(HeweatherHttpRequest(msg))
					if result == "" {
						result = string(TulingHttpRequest(msg))
					}
				} else {
					result = string(TulingHttpRequest(msg))
				}
			}
			ctx.Data = result
		}
		cfg := &jas.Config{}
		cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
			len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
			return len
		}
		ctx.SetConfig(cfg)
	case "video":
		{
		}
	case "shortvideo":
		{
		}
	case "location":
		{
		}
	case "link":
		{
		}
	}
}
