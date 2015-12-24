package console

import (
	"encoding/json"
	"fmt"
)

type HeweatherRequest struct {
	Key    string
	City   string
	CityId string
	CityIp string
}

func (this *HeweatherRequest) ToRequest() string {
	var req string = "key="
	req += this.Key
	if this.City != "" {
		req += "&city="
		req += this.City
	}
	if this.CityId != "" {
		req += "&cityid="
		req += this.CityId
	}
	if this.CityIp != "" {
		req += "&cityip="
		req += this.CityIp
	}
	return req
}

func (this *HeweatherRequest) ToTicket() string {
	return ""
}

/*-------------------------sub info ----------------------------*/
type BaseUpdateTime struct {
	Loc string //数据更新的当地时间
	Utc string //数据更新的UTC时间
}
type Astro struct {
	Sr string //日出时间
	Ss string //日落时间
}
type Aqi struct { //空气质量指数
	Aqi  string //空气质量指数
	Pm25 string //PM2.5 1小时平均值(ug/m³)
	Pm10 string //PM10 1小时平均值(ug/m³)
	So2  string //二氧化硫1小时平均值(ug/m³)
	No2  string //二氧化氮1小时平均值(ug/m³)
	Co   string //一氧化碳1小时平均值(ug/m³)
	O3   string //臭氧1小时平均值(ug/m³)
	Qlty string //空气质量类别
}
type Wind struct {
	Spd string //风速(Kmph)
	Sc  string //风力等级
	Deg string //风向(角度)
	Dir string //风向(方向)
}
type Cond struct {
	Code string //天气代码
	Txt  string //天气描述
}
type CondAll struct {
	Code_d string //白天天气代码
	Txt_d  string //白天天气描述
	Code_n string //晚上天气代码
	Txt_n  string //晚上天气描述
}
type Tmp struct {
	Max string //最该温度(摄氏度)
	Min string //最低温度(摄氏度)
}

/*----------------------------------------------------------------*/

type HeweatherBase struct { //城市基本信息
	City   string         //城市名称
	Id     string         //城市ID
	Cnty   string         //国家名称
	Lat    string         //纬度
	Lon    string         //经度
	Update BaseUpdateTime //更新时间
	//	Loc  string `json:"update>loc"` //数据更新的当地时间
	//	Utc  string `json:"update>utc"` //数据更新的UTC时间
}

type HeweatherAqi struct { //空气质量指数
	City Aqi //空气质量指数
	/*
		Aqi  string `json:"city>aqi"`  //空气质量指数
		Pm25 string `json:"city>pm25"` //PM2.5 1小时平均值(ug/m³)
		Pm10 string `json:"city>pm10"` //PM10 1小时平均值(ug/m³)
		So2  string `json:"city>so2"`  //二氧化硫1小时平均值(ug/m³)
		No2  string `json:"city>no2"`  //二氧化氮1小时平均值(ug/m³)
		Co   string `json:"city>co"`   //一氧化碳1小时平均值(ug/m³)
		O3   string `json:"city>o3"`   //臭氧1小时平均值(ug/m³)
		Qlty string `json:"city>qlty"` //空气质量类别
	*/
}

type HeweatherSuggestionInfo struct {
	Brf string //简介
	Txt string //详情
}
type HeweatherSuggestion struct { //生活指数
	Drsg  HeweatherSuggestionInfo //穿衣指数
	Uv    HeweatherSuggestionInfo //紫外线指数
	Cw    HeweatherSuggestionInfo //洗车指数
	Trav  HeweatherSuggestionInfo //旅游指数
	Flu   HeweatherSuggestionInfo //感冒指数
	Sport HeweatherSuggestionInfo //运动指数
}

type HeweatherNow struct { //实况天气
	Tmp  string //当前温度(摄氏度)
	Fl   string //体感温度
	Wind Wind   //风力状况
	Cond Cond   //天气状况
	/*
		Spd  string `json:"wind>spd"`  //风速(Kmph)
		Sc   string `json:"wind>sc"`   //风力等级
		Deg  string `json:"wind>deg"`  //风向(角度)
		Dir  string `json:"wind>dir"`  //风向(方向)
		Code string `json:"cond>code"` //天气代码
		Txt  string `json:"cond>txt"`  //天气描述
	*/
	Pcpn string //降雨量(mm)
	Hum  string //湿度(%)
	Pres string //气压
	Vis  string //能见度(km)
}

type HeweatherHourlyForecast struct {
	Date string //当地日期和时间
	Tmp  string //当前温度(摄氏度)
	Wind Wind
	/*
		Spd  string `json:"wind>spd"` //风速(Kmph)
		Sc   string `json:"wind>sc"`  //风力等级
		Deg  string `json:"wind>deg"` //风向(角度)
		Dir  string `json:"wind>dir"` //风向(方向)
	*/
	Pop  string //降水概率
	Hum  string //湿度(%)
	Pres string //气压
}

type HeweatherDailyForecast struct {
	Date string //当地日期
	/*
		Sr     string `json:"astro>sr"`    //日出时间
		Ss     string `json:"astro>ss"`    //日落时间
		Max    string `json:"tmp>max"`     //最该温度(摄氏度)
		Min    string `json:"tmp>min"`     //最低温度(摄氏度)
		Spd    string `json:"wind>spd"`    //风速(Kmph)
		Sc     string `json:"wind>sc"`     //风力等级
		Deg    string `json:"wind>deg"`    //风向(角度)
		Dir    string `json:"wind>dir"`    //风向(方向)
		Code_d string `json:"cond>code_d"` //白天天气代码
		Txt_d  string `json:"cond>txt_d"`  //白天天气描述
		Code_n string `json:"cond>code_n"` //夜间天气代码
		Txt_n  string `json:"cond>txt_n"`  //夜间天气描述
	*/
	Astro Astro   //日出日落时间
	Tmp   Tmp     //温度情况
	Wind  Wind    //风力状况
	Cond  CondAll //天气状况
	Pcpn  string  //降雨量(mm)
	Pop   string  //降水概率
	Hum   string  //湿度(%)
	Pres  string  //气压
	Vis   string  //能见度(km)
}

type HeweatherBody struct {
	Status          string
	Basic           HeweatherBase
	Aqi             HeweatherAqi
	Now             HeweatherNow
	Suggestion      HeweatherSuggestion
	Hourly_Forecast []HeweatherHourlyForecast
	Daily_Forecast  []HeweatherDailyForecast
}
type HeweatherResponse struct {
	Data []HeweatherBody `json:"HeWeather data service 3.0"`
}

func (this *HeweatherResponse) ParseJson(jsonObj []byte) {
	if err := json.Unmarshal(jsonObj, this); err != nil {
		fmt.Println(err)
	}
}

func (this *HeweatherResponse) ToString() string {
	if len(this.Data) <= 0 {
		return string("")
	}
	body := this.Data[0]
	str := "更新时间：" + body.Basic.Update.Loc + "\r\n"
	str += body.Basic.City + "天气：" + body.Now.Cond.Txt + " " + body.Now.Wind.Dir + " " + body.Now.Wind.Sc + "\r\n"
	str += "当前温度：" + body.Now.Tmp + "°C\r\n" + "体感温度：" + body.Now.Fl + "°C\r\n"
	str += "能见度：" + body.Now.Vis + "\r\n湿度：" + body.Now.Hum + "\r\n降雨量：" + body.Now.Pcpn + "\r\n气压：" + body.Now.Pres + "\r\n\r\n"
	if body.Aqi.City.Pm25 != "" {
		str += "空气质量：" + body.Aqi.City.Pm25 + body.Aqi.City.Qlty + "\r\n" +
			"Pm10:" + body.Aqi.City.Pm10 + "\r\n" +
			"二氧化硫(ug/m3/h):" + body.Aqi.City.So2 + "\r\n二氧化氮(ug/m3/h):" + body.Aqi.City.No2 + "\r\n" +
			"一氧化碳(ug/m3/h):" + body.Aqi.City.Co + "\r\n臭氧(ug/m3/h):" + body.Aqi.City.O3 + "\r\n\r\n"
	}
	str += "穿衣指数：" + body.Suggestion.Drsg.Brf + "\r\n" +
		"感冒指数：" + body.Suggestion.Flu.Brf + "\r\n" +
		"运动指数：" + body.Suggestion.Sport.Brf + "\r\n" +
		"紫外线指数：" + body.Suggestion.Uv.Brf + "\r\n" +
		"洗车指数：" + body.Suggestion.Cw.Brf + "\r\n" +
		"旅游指数：" + body.Suggestion.Trav.Brf + "\r\n"
	return str
}
