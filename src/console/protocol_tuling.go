package console

import (
	"encoding/json"
	"fmt"
)

type TulingRequest struct {
	Key    string
	Info   string
	UserId string
	Loc    string
	Lon    string
	Lat    string
}

func (this *TulingRequest) ToRequest() string {
	var req string = "key="
	req += this.Key
	req += "&info="
	req += this.Info
	if this.UserId != "" {
		req += "&userid="
		req += this.UserId
	}
	if this.Loc != "" {
		req += "&loc="
		req += this.Loc
	}
	if this.Lon != "" {
		req += "&lon="
		req += this.Lon
	}
	if this.Lat != "" {
		req += "&lat="
		req += this.Lat
	}
	return req
}

func (this *TulingRequest) ToTicket() string {
	return ""
}

type TulingNews struct {
	Article   string
	Source    string
	Icon      string
	DetailUrl string
}
type TulingTrain struct {
	TrainNum  string
	Start     string
	Terminal  string
	StartTime string
	EndTime   string
	Icon      string
	DetailUrl string
}
type TulingAir struct {
	Flight    string
	StartTime string
	EndTime   string
	Icon      string
}
type TulingFood struct {
	Name      string
	Icon      string
	Info      string
	DetailUrl string
}
type TulingData struct {
        Article   string
        Source    string
        Icon      string
        DetailUrl string

        TrainNum  string
        Start     string
        Terminal  string
        StartTime string
        EndTime   string

        Flight    string

        Name      string
        Info      string
}

type TulingResponse struct {
	Code       int32
	Text       string
	Url        string
	List  	   []TulingData
}

func (this *TulingResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

func (this *TulingResponse) Print() {
        fmt.Println("code:", this.Code)
	fmt.Println("text:", this.Text)
	fmt.Println("url:", this.Url)
}