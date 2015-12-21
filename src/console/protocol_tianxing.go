package console

import (
	"encoding/json"
	"fmt"
)

type TianxingRequest struct {
	Key  string
	Num  string
	Rand string
	Word string
	Page string
}

func (this *TianxingRequest) ToRequest() string {
	var req string = "key="
	req += this.Key
	req += "&num="
	req += this.Num
	if this.Rand != "" {
		req += "&rand="
		req += this.Rand
	}
	if this.Word != "" {
		req += "&word="
		req += this.Word
	}
	if this.Page != "" {
		req += "&page="
		req += this.Page
	}
	return req
}

func (this *TianxingRequest) ToTicket() string {
	return ""
}

type TianxingNews struct {
	Hottime     string
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type TianxingResponse struct {
	Code     int32
	Msg      string
	NewsList []TianxingNews
}

func (this *TianxingResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

func (this *TianxingResponse) Print() {
	fmt.Println("code:", this.Code)
	fmt.Println("msg:", this.Msg)
}
