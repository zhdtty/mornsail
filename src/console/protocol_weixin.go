package console

import (
	"encoding/json"
	"fmt"
)

type WeixinTokenRequest struct {
	GrantType string
	AppId     string
	AppSecret string
}

func (this *WeixinTokenRequest) ToRequest() string {
	var req string = "grant_type="
	req += this.GrantType
	req += "&appid="
	req += this.AppId
	req += "&secret="
	req += this.AppSecret
	return req
}
func (this *WeixinTokenRequest) ToTicket() string {
	return ""
}

type WeixinTokenResponse struct {
	Access_Token string
	Expires_In   int
}

func (this *WeixinTokenResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

type WeixinMediaLoadRequest struct {
	AccessToken string
	MediaId     string
}

func (this *WeixinMediaLoadRequest) ToRequest() string {
	var req string = "access_token="
	req += this.AccessToken
	req += "&media_id="
	req += this.MediaId
	return req
}
func (this *WeixinMediaLoadRequest) ToTicket() string {
	return ""
}

type WeixinMediaLoadResponse struct {
}

func (this *WeixinMediaLoadResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

//Custom message
type WeixinCustomMessageText struct {
}

type WeixinCustomMessageRequest struct {
	AccessToken string
}

func (this *WeixinCustomMessageRequest) ToRequest() string {
	var req string = "access_token="
	req += this.AccessToken
	return req
}
func (this *WeixinCustomMessageRequest) ToTicket() string {
	return ""
}

type WeixinCustomMessageResponse struct {
}

func (this *WeixinCustomMessageResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}
