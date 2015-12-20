package console

import ()

var TULING_ADDRESS string = "www.tuling123.com"
var TULING_API_KEY = "58e4a76358d33acddb3014060ffbd952"

var OPS_TULING_QUERY string = "/openapi/api?"

func DoTulingQuery(req TulingRequest) *TulingResponse {
        req.Key = TULING_API_KEY
        params := req.ToRequest()
        result := SendHttpRequest(TULING_ADDRESS, OPS_TULING_QUERY, "GET", params, false)
	resp := &TulingResponse{}
	resp.ParseJson(result)
	return resp
}
