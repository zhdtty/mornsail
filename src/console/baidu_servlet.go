package console

import (
	"fmt"
	"github.com/coocood/jas"
	"io/ioutil"
)

type Baidu struct{}

func (*Baidu) GetCallback(ctx *jas.Context) {
	fmt.Println("get baidu callback")
	resJson, err := ioutil.ReadAll(ctx.Body)
	if err != nil {
		fmt.Println("baidu body body error")
		return
	}
	ctx.Body.Close()
	//	b := DecodeJson(resJson)
	resp := &BaiduVoiceResponse{}
	resp.ParseJson(resJson)

	msg := BDVoiceQueue.GetAndDelete()
	if msg == nil || resp.Err_No != 0 {
		return
	}
	msg.Content = resp.Result[0]

	var result string
	result = string(TulingHttpRequest(msg))
	fmt.Println("Tuling result after aidu callback : ", result)
}
