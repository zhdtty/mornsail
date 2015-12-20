package console

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var LB_SECURIT_KEY string = "His7m$#4uO2(~rP9Yu^9N"
var GAME_SECURIT_KEY string = "kL@2U%va~T23&E2o$*gd2%(k9T2mv#6s"
var rid int = 0

func GetUrl(address string, uinterface string, data string, https bool) bytes.Buffer {
	var bUrl bytes.Buffer
	if https {
	     bUrl.WriteString("https://")
	} else {
	     bUrl.WriteString("http://")
	}
	bUrl.WriteString(address)
	bUrl.WriteString(uinterface)
	bUrl.WriteString(data)
	return bUrl
}

func GetUrlWithoutData(address string, uinterface string, https bool) bytes.Buffer {
	var bUrl bytes.Buffer
        if https {
             bUrl.WriteString("https://")
        } else {
             bUrl.WriteString("http://")
        }
	bUrl.WriteString(address)
	bUrl.WriteString(uinterface)
	return bUrl
}

func RestGet(bufUrl *bytes.Buffer) []byte {
	res, err := http.Get(bufUrl.String())
	if err != nil {
		var bErr []byte
		fmt.Println("rest get error : ", err)
		return bErr
	}
	resJson, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		var bErr []byte
		fmt.Println("rest read get body error : ", err)
		return bErr
	}
	//        fmt.Printf("%s\r\n", resJson)
	return resJson
}

func RestPost(bufUrl *bytes.Buffer, data url.Values) []byte {
	res, err := http.PostForm(bufUrl.String(), data)
	if err != nil {
		var bErr []byte
		fmt.Println("rest post error : ", err)
		return bErr
	}
	resJson, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		var bErr []byte
		fmt.Println("rest read post body error : ", err)
		return bErr
	}
	//        fmt.Printf("%s\r\n", resJson)
	return resJson
}

func toMd5(sig string) string {
	bData := []byte(sig)
	sMd5 := md5.Sum(bData)
	var bMd5Sum []byte
	bMd5Sum = sMd5[0:16]
	strMd5 := hex.EncodeToString(bMd5Sum)
	return strMd5
}

func toMd5Base64(sig string) string {
	bData := []byte(sig)
	sMd5 := md5.Sum(bData)
	var bMd5Sum []byte
	bMd5Sum = sMd5[0:16]
	strMd5 := hex.EncodeToString(bMd5Sum)
	bMd5 := []byte(strMd5)
	baseVal := base64.StdEncoding.EncodeToString(bMd5)
	return baseVal
}

func gameSecurityKey(token string) string {
	var intkey int
	intkey = 0
	count := 0
	for i := 0; i < len(token); i++ {
		word := token[i]
		if word < '0' || word > '9' {
			continue
		}
		intkey *= 10
		intkey += (int(word) - 48)
		count++
		if count >= 9 {
			break
		}
	}

	var ss string
	var choose bool
	var check int
	choose = false
	check = 0x00000001
	for i := 0; i < 32; i++ {
		choose = ((intkey & (check << uint(i))) != 0)
		if choose {
			ss += string(GAME_SECURIT_KEY[i])
		}
	}
	//    fmt.Println(ss);
	return ss
}
func ParseObjectJsonNoDecode(objJson []byte) map[string]string {
	var result map[string]string
	if err := json.Unmarshal(objJson, &result); err != nil {
		fmt.Println("parse json error", err)
		return result
	}
	return result
}

func ParseObjectJson(objJson []byte) map[string]string {
	fmt.Println(string(objJson))
	oj, err := url.QueryUnescape(string(objJson)) //base64.URLEncoding.DecodeString(string(objJson));
	fmt.Println("decode:", oj)
	if err != nil {
		fmt.Println("decode string error : ", err)
	}
	boj := []byte(oj)
	var result map[string]string
	if err := json.Unmarshal(boj, &result); err != nil {
		fmt.Println(err)
		return result
	}
	//        fmt.Println(result);
	return result
}

func DecodeJson(objJson []byte) []byte {
	oj, err := url.QueryUnescape(string(objJson))
	if err != nil {
		fmt.Println("decode string error : ", err)
	}
	boj := []byte(oj)
	return boj
}
