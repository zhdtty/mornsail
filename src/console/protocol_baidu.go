package console

import (
	"encoding/json"
	"fmt"
)

type BaiduTokenRequest struct {
	GrantType string
	AppKey    string
	SecretKey string
}

func (this *BaiduTokenRequest) ToRequest() string {
	var req string = "grant_type="
	req += this.GrantType
	req += "&client_id="
	req += this.AppKey
	req += "&client_secret="
	req += this.SecretKey
	return req
}
func (this *BaiduTokenRequest) ToTicket() string {
	return ""
}

type BaiduTokenResponse struct {
	Access_Token   string
	Expires_In     int
	Refresh_Token  string
	Scope          string
	Session_Key    string
	Session_Secret string
}

func (this *BaiduTokenResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

type BaiduVoiceData struct {
	Cuid    string `json:"cuid"`
	Token   string `json:"token"`
	Format  string `json:"format"`
	Rate    int    `json:"rate"`
	Channel int    `json:"channel"`
	Speech  string `json:"speech"`
	Len     int    `json:"len"`
}
type BaiduVoiceUrl struct {
	Format   string
	Rate     int
	Channel  int
	Url      string
	Callback string
}

type BaiduVoiceRequest struct {
	Cuid  string
	Token string
	Ptc   string
	Lan   string
	Vdata BaiduVoiceData
	VUrl  BaiduVoiceUrl
}

func (this *BaiduVoiceRequest) ToRequest() string {
	/*
		var req string = "cuid="
		req += this.Cuid
		req += "&token="
		req += this.Token
		if this.Ptc != "" {
			req += "&ptc="
			req += this.Ptc
		}
		if this.Lan != "" {
			req += "&lan="
			req += this.Lan
		}
		return req
	*/
	return string("")
}

func (this *BaiduVoiceRequest) ToTicket() string {
	return ""
}

type BaiduVoiceResponse struct {
	Err_No  int
	Err_Msg string
	Sn      string
	Result  []string
}

func (this *BaiduVoiceResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

func SerialToJSON(v interface{}) []byte {
	output, err := json.MarshalIndent(v, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output
}
