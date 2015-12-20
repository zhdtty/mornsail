package console

import (
        "fmt"
	"encoding/json"
)

type WeixinTokenRequest struct {
	GrantType string
	AppId string
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
     Expires_In int
}

func (this *WeixinTokenResponse) ParseJson(jsonObj []byte) {
        if err := json.Unmarshal(jsonObj, this); err != nil {
                fmt.Println(err)
        }
}