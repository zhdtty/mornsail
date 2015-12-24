package console

import (
	"bytes"
	"container/list"
	"fmt"
	//	"reflect"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var BAIDU_ADDRESS string = "openapi.baidu.com"
var BAIDU_VOICE_ADDRESS string = "vop.baidu.com"

var BAIDU_APP_ID = "6647463"
var BAIDU_API_KEY = "CX8plDjHa7VnGEwO2t9iKwqF"
var BAIDU_SECRET_KEY = "ba4bd90d941486b5779a7dc3c0c3242c"
var BAIDU_ACCESS_TOKEN string = ""
var BAIDU_CALLBACK_URL string = "http://120.24.5.207/baidu/callback"

var BAIDU_OAUTH_TOKEN string = "/oauth/2.0/token?"

var BAIDU_VOICE_API string = "/server_api"

type VoiceQueue struct {
	sync.Mutex
	msgs *list.List
}

func NewVoiceQueue() *VoiceQueue {
	vq := &VoiceQueue{
		msgs: list.New(),
	}
	return vq
}

func (vq *VoiceQueue) Add(msg *WeixinMsg) {
	vq.Lock()
	defer vq.Unlock()
	vq.msgs.PushBack(msg)
}
func (vq *VoiceQueue) GetAndDelete() *WeixinMsg {
	vq.Lock()
	defer vq.Unlock()
	msg := vq.msgs.Front()
	if msg != nil {
		vq.msgs.Remove(msg)
		return msg.Value.(*WeixinMsg) //reflect.ValueOf(msg.Value)
	}
	return nil
}

func DoBaiduToken(req BaiduTokenRequest) []byte {
	req.AppKey = BAIDU_API_KEY
	req.SecretKey = BAIDU_SECRET_KEY
	req.GrantType = "client_credentials"
	params := req.ToRequest()
	result := SendHttpRequest(BAIDU_ADDRESS, BAIDU_OAUTH_TOKEN, "GET", params, true)
	var resp BaiduTokenResponse
	resp.ParseJson(result)
	BAIDU_ACCESS_TOKEN = resp.Access_Token
	fmt.Println("baidu access token :", BAIDU_ACCESS_TOKEN)
	return []byte(resp.Access_Token)
}

func DoBaiduVoice(req BaiduVoiceRequest) string {
	//	params := req.ToRequest()
	body := SerialToJSON(req.Vdata)
	fmt.Println("baidu voice request json :", string(body))

	client := &http.Client{}
	url := "http://" + BAIDU_VOICE_ADDRESS + BAIDU_VOICE_API
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("content-length", fmt.Sprintf("%d", len(body)))
	resp, err := client.Do(request)
	if err != nil {
		var bErr []byte
		fmt.Println("rest post error : ", err)
		return string(bErr)
	}
	resJson, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		var bErr []byte
		fmt.Println("rest read post body error : ", err)
		return string(bErr)
	}
	var bdResp BaiduVoiceResponse
	bdResp.ParseJson(resJson)
	if bdResp.Err_No == 0 && len(bdResp.Result) > 0 {
		result := bdResp.Result[0]
		return strings.TrimSuffix(result, "，")
	}
	fmt.Println("Baidu voice errno :", bdResp.Err_No, ", errmsg :", bdResp.Err_Msg)

	return string("")
}

func BaiduVoiceHttpRequest(reqMsg *WeixinMsg) string {
	if reqMsg.MsgType != "voice" {
		return ""
	}
	var mediaReq WeixinMediaLoadRequest
	mediaReq.MediaId = reqMsg.MediaId
	voiceData := DoLoadMedia(mediaReq)
	voiceStr := base64.StdEncoding.EncodeToString(voiceData)

	var req BaiduVoiceRequest
	req.Vdata.Cuid = reqMsg.FromUserName
	req.Vdata.Token = BAIDU_ACCESS_TOKEN
	req.Vdata.Format = "amr"
	req.Vdata.Rate = 8000
	req.Vdata.Channel = 1
	req.Vdata.Speech = voiceStr
	req.Vdata.Len = len(voiceData)

	fmt.Println("baidu voice use weixin access token :", ACCESS_TOKEN)
	fmt.Println("baidu voice use weixin media id :", reqMsg.MediaId)

	return DoBaiduVoice(req)
}

var BDVoiceQueue *VoiceQueue = NewVoiceQueue()
