package console

import (
       "github.com/coocood/jas"
)

var WEIXIN_TOKEN string = "test"

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

//        fmt.Println("echostr:", echostr, "nonce:", nonce, "timestamp:", timestamp, "sig:", sig)
//        fmt.Println("sha1str:", data)
//        fmt.Println("sha1sig:", sha1Sig)

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
        n, _ = ctx.Body.Read(b)
	if n <= 0 {
              fmt.Println("body:",string(b))
	}

        var msg WeixinMsg
        msg.ParseMsg(b)
        msg.Print()

        switch msg.MsgType {
	       case "text": {
	               var msgResp WeixinMsgTextResponse
        	       msgResp.ToUserName = "o6OKFxDhoecsCZzIMxCerGwpQE9c"
        	       msgResp.FromUserName = "gh_5b1770a0f8f4"
		       msgResp.CreateTime = msg.CreateTime
        	       msgResp.MsgType = msg.MsgType
        	       msgResp.Content = "你逗我玩吗"
        	       rs := SerialToXML(msgResp)
        	       ctx.Data = string(rs)
        	       cfg := &jas.Config{}
        	       cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
                       		       len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
                		       return len
        	       }
        	       ctx.SetConfig(cfg)
	       }
	       case "image":{}
	       case "voice":{}
	       case "video":{}
	       case "shortvideo":{}
	       case "location":{}
	       case "link":{}
	}
}