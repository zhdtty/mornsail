package console

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strings"
)

//game server interface
//login
var Account_login string = "/player/account_login?"         //get
var Account_rand_name string = "/player/account_rand_name?" //get
var Account_create string = "/player/account_create?"       //post
var Account_load_char string = "/player/account_load_char?" //get

//card
var Card_load string = "/player/card_load?" //get

//battle
var Battle_report_store string = "/player/battle_report_store?" //get
var Battle_report_load string = "/player/battle_report_load?"   //get

//rank
var Rank_load string = "/player/rank_load?" //get

//test
var Test string = "/player/test_response_time?" //get

func EncryptXOR(instr string, key string) []byte {
	i := 0
	ki := 0
	bout := make([]byte, len(instr))
	inLen := len(instr)
	keyLen := len(key)
	for i < inLen {
		bout[i] = instr[i] ^ key[ki]
		ki++
		if ki >= keyLen {
			ki = 0
		}
		i++
	}
	return bout
}

func DecryptXOR(instr string, key string) []byte {
	return EncryptXOR(instr, key)
}

func EncryptParam(param string) string {
	//        fmt.Println("url:", param)
	val := EncryptXOR(param, GAME_SECURIT_KEY)
	oparam := "p="
	oparam += url.QueryEscape(base64.StdEncoding.EncodeToString(val))
	return oparam
}

func DecryptResult(ret []byte) []byte {
	strRet := string(ret)
	strRet = strings.Replace(strRet, " ", "+", -1)
	//        fmt.Println(strRet)
	cryptRet, err := base64.StdEncoding.DecodeString(strRet)
	if err != nil {
		return []byte("")
	}
	nRet := DecryptXOR(string(cryptRet), GAME_SECURIT_KEY)
	return DecodeJson(nRet)
}

func SendHttpRequest(gameAddress string, inf string, method string, req string, https bool) []byte {
	bUrl := GetUrl(gameAddress, inf, req, https)
	fmt.Println("url:", bUrl.String())
	if method == "GET" {
		resJson := RestGet(&bUrl)
		return DecodeJson(resJson)
	} else if method == "POST" {
		var v url.Values
		resJson := RestPost(&bUrl, v)
		return DecodeJson(resJson)
	} else {
		fmt.Println("Invalid http request method! method:", method)
		var result []byte
		return result
	}
}

func SendHttpRequestNotJson(gameAddress string, inf string, method string, req string, https bool) []byte {
	bUrl := GetUrl(gameAddress, inf, req, https)
	fmt.Println("url:", bUrl.String())
	if method == "GET" {
		return RestGet(&bUrl)
	} else if method == "POST" {
		var v url.Values
		return RestPost(&bUrl, v)
	} else {
		fmt.Println("Invalid http request method! method:", method)
		var result []byte
		return result
	}
}

func PostHttpRequestWithBody(gameAddress string, inf string, req string, bodyType string, body io.Reader, https bool) []byte {
	bUrl := GetUrl(gameAddress, inf, req, https)
	fmt.Println("url with body:", bUrl.String())
	resJson := RestPostWithBody(&bUrl, bodyType, body)
	fmt.Println("url with body res json:", string(resJson))
	return DecodeJson(resJson)
}
