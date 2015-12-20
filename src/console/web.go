package console

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	//	"runtime/pprof"
	"crypto/sha1"
	"driver"
	"github.com/coocood/jas"
	"io"
	"reflect"
	"sort"
)

func Init() {
	fmt.Println("console init")
	go webUI()
}

func sha1s(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func redisHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, fmt.Sprintf("Cur redis actives : %d", driver.RedisPool.ActiveCount()))
}

type Hello struct{}

func (*Hello) GetWorld(ctx *jas.Context) { // `GET /v1/hello`
	var req WeixinTokenRequest
	_ = DoToken(req)
	ctx.Data = ACCESS_TOKEN
	//response: `{"data":"hello world","error":null}`
}

func (*Hello) GetTuling(ctx *jas.Context) {
     var req TulingRequest
     req.Info = "新闻"
     resp := DoTulingQuery(req)
     resp.Print()
     ctx.Data = "success"
}

type Weixin struct{}

func (*Weixin) Get(ctx *jas.Context) {
	fmt.Println("Get weixin")
	echostr := ctx.RequireString("echostr")
	nonce := ctx.RequireString("nonce")
	timestamp := ctx.RequireString("timestamp")
	sig := ctx.RequireString("signature")

	token := "test"

	params := []string{nonce, timestamp, token}
	sort.Sort(sort.StringSlice(params))
	data := ""
	for _, v := range params {
		data += v
	}
	sha1Sig := string(sha1s(data))

	fmt.Println("echostr:", echostr, "nonce:", nonce, "timestamp:", timestamp, "sig:", sig)
	fmt.Println("sha1str:", data)
	fmt.Println("sha1sig:", sha1Sig)

	ctx.Data = echostr
	if sig == sha1Sig {
		cfg := &jas.Config{}
		cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
			len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
			return len
		}
		ctx.SetConfig(cfg)
	}
}

func (*Weixin) Post(ctx *jas.Context) {
	fmt.Println("Post weixin msg")
	b := make([]byte, 2048)
	_, _ = ctx.Body.Read(b)
	//	fmt.Println("body:",string(b))
	//	fmt.Println("len:", n)

	var msg WeixinMsg
	msg.ParseMsg(b)
	msg.Print()

	switch msg.
	var msgResp WeixinMsgTextResponse
	msgResp.ToUserName = "o6OKFxDhoecsCZzIMxCerGwpQE9c"
	msgResp.FromUserName = "gh_5b1770a0f8f4"
	msgResp.CreateTime = msg.CreateTime
	msgResp.MsgType = msg.MsgType
	msgResp.Content = "你逗我玩吗"
	rs := SerialToXML(msgResp)
	ctx.Data = string(rs)
	cfg := &jas.Config{}
	cfg.HijackWrite = func(writer io.Writer, ctx *jas.Context) int {
		len, _ := writer.Write([]byte(reflect.ValueOf(ctx.Data).String()))
		return len
	}
	ctx.SetConfig(cfg)

	fmt.Println(ctx.Data)
	//	ctx.Data = "Post"
}

/*
func goHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func heapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("heap")
	p.WriteTo(w, 1)
}

func threadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("threadcreate")
	p.WriteTo(w, 1)
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("block")
	p.WriteTo(w, 1)
}
*/
func webUI() {
	/*
		tpl, err := template.New("XML translation").Parse(html)
		if err != nil {
			panic(err)
		}

		requestMain := func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("main page")

			files, err := ListSubDir(".")
			err = tpl.ExecuteTemplate(w, "main", files)
			if err != nil {
				log.Fatal(err)
			}
		}

		requestCpu := func(w http.ResponseWriter, req *http.Request) {
			dir, err := ioutil.ReadDir(".")
			if err != nil {
				io.WriteString(w, "非法目录\n")
			}

			start := time.Now()
			startSec := start.UnixNano() / 1000000

			wg := new(sync.WaitGroup)
			for _, fi := range dir {
				if !fi.IsDir() && strings.HasPrefix(fi.Name(), "fanyi_") {
					doTrans("", fi.Name(), false, wg)
				}
			}
			wg.Wait()

			end := time.Now()
			endSec := end.UnixNano() / 1000000

			result := fmt.Sprintf("翻译文件消耗%d毫秒\n", int(endSec-startSec))
			err = tpl.ExecuteTemplate(w, "tip", result)
			if err != nil {
				log.Fatal(err)
			}
		}
	*/
	/*
		http.HandleFunc("/", goHandler)
		http.HandleFunc("/heap", heapHandler)
		http.HandleFunc("/thread", threadHandler)
		http.HandleFunc("/block", blockHandler)
	*/
	http.HandleFunc("/redis", redisHandle)

	router := jas.NewRouter(new(Hello), new(Weixin))
	router.BasePath = "/"
	fmt.Println(router.HandledPaths(true))
	//output: `GET /v1/hello`
	http.Handle(router.BasePath, router)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

var html string = `{{define "main"}}{{/* 服务器状态页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p>
        <form>
            <tr>使用说明：</tr><br>
            <tr>1. 点击"CPU状态"，查看CPU状态</tr><br>
            <tr>2. 点击"内存状态"，查看内存状态</tr><br>
         </form>
    </p>
    <p>
        <a href="/cpu"><input type="submit" value="CPU状态"/></a>
        <a href="/mem"><input type="submit" value="内存状态"/></a>
    </p>
</body>
</html>{{end}}`
