package console

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	//	"runtime/pprof"
	"crypto/md5"
	"crypto/sha1"
	"driver"
	"github.com/coocood/jas"
	"io"
	"io/ioutil"
	//	"reflect"
	//	"sort"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
	"timer"
	"tool"
)

var SEND_MAIL_KEY string = "zhdtty200172"
var SEND_MAIL_MIN_NUMS int = 10
var SEND_MAIL_MAX_NUMS int = 10000

func Init() {
	fmt.Println("console init")
	go webUI()

	var req WeixinTokenRequest
	_ = DoToken(req)
	var req1 BaiduTokenRequest
	_ = DoBaiduToken(req1)
	timer.SvrTimer.AddIntervalTimer(3600, func() {
		var req WeixinTokenRequest
		_ = DoToken(req)
		var req1 BaiduTokenRequest
		_ = DoBaiduToken(req1)
	}, true)

	//	tool.ReadFileMails("1.txt")
	//	tool.ReadMailListFromRedis("mail_test")
	//	tool.SendMailForTest()
	//tool.ReadFileSendMail("1.txt", "")
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

func ToMd5(sig string) string {
	bData := []byte(sig)
	sMd5 := md5.Sum(bData)
	var bMd5Sum []byte
	bMd5Sum = sMd5[0:16]
	strMd5 := hex.EncodeToString(bMd5Sum)
	return strMd5
}

type DirInfo struct {
	Name string
	Size int64
	//      Mode    int
	ModTime time.Time
}

func ListSubDir(dirPath string) (files []DirInfo, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		//              if fi.IsDir() {
		dirInfo := DirInfo{}
		dirInfo.Name = fi.Name()
		dirInfo.Size = fi.Size() / 1024
		//                      dirInfo.Mode = int(fi.Mode())
		dirInfo.ModTime = fi.ModTime()
		files = append(files, dirInfo)
		//              }
	}
	return files, nil
}

func webUI() {
	tpl, err := template.New("XML translation").Parse(html1)
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

	requestFile := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("file page")

		files, err := ListSubDir(".")
		err = tpl.ExecuteTemplate(w, "file", files)
		if err != nil {
			log.Fatal(err)
		}
	}

	requestUpload := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("upload file")

		if "POST" != req.Method {
			files, err := ListSubDir(".")
			err = tpl.ExecuteTemplate(w, "main", files)
			if err != nil {
				log.Fatal(err)
			}
		}

		file, head, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fW, err := os.Create(head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()

		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		fmt.Println(head.Filename)

		tablekey := "mail_set"
		prefix := strings.Split(head.Filename, ".")
		if len(prefix) > 0 {
			if prefix[0] == "male" || prefix[0] == "female" {
				tablekey += "." + prefix[0]
			}
		}
		mailAry, err := tool.ReadFileMails(head.Filename, tablekey)
		if err != nil {
			mailAry = append(mailAry, fmt.Sprintf("%v", err))
		}
		err = tpl.ExecuteTemplate(w, "tipary", mailAry)
		if err != nil {
			log.Fatal(err)
		}
	}

	requestMd5 := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("calc MD5")

		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		if len(req.Form["name"]) < 1 {
			io.WriteString(w, "参数错误!\n")
			return
		}

		name := template.HTMLEscapeString(req.Form.Get("name"))
		nameMd5 := ToMd5(name)

		err = tpl.ExecuteTemplate(w, "tip", nameMd5)
		if err != nil {
			log.Fatal(err)
		}
	}

	requestSendMail := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("send mail")

		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		keyLen := len(req.Form["key"])
		//numLen := len(req.Form["nums"])
		titleLen := len(req.Form["title"])
		contentLen := len(req.Form["content"])
		if keyLen < 1 || titleLen < 1 || contentLen < 1 {
			io.WriteString(w, "参数错误!\n")
			return
		}

		nums := SEND_MAIL_MIN_NUMS

		user := template.HTMLEscapeString(req.Form.Get("user"))
		passwd := template.HTMLEscapeString(req.Form.Get("passwd"))
		host := template.HTMLEscapeString(req.Form.Get("host"))
		recvs := template.HTMLEscapeString(req.Form.Get("recvs"))
		key := template.HTMLEscapeString(req.Form.Get("key"))
		numstr := template.HTMLEscapeString(req.Form.Get("nums"))
		title := template.HTMLEscapeString(req.Form.Get("title"))
		content := template.HTMLEscapeString(req.Form.Get("content"))
		if key != SEND_MAIL_KEY {
			io.WriteString(w, "暂无权限!\n")
			return
		}

		tmpNums, err := strconv.Atoi(numstr)
		if err == nil && tmpNums >= SEND_MAIL_MIN_NUMS && tmpNums <= SEND_MAIL_MAX_NUMS {
			nums = tmpNums
		}
		resTip := string("已发送")
		if user != "" && passwd != "" && host != "" && recvs != "" {
			recvsAry := strings.Split(recvs, ";")
			_ = tool.SendMailCustomRequest(host, user, passwd, title, content, recvsAry, nums)
		} else {
			//tool.TestSendMailRequest(title, content, nums)
			//resTip = string("不支持默认参数发送")
			//tool.ReleaseSendMailRequest(title, content, "mail_test", nums)
			tool.TimerSendMailRequest(title, content, "mail_set", nums)
		}

		err = tpl.ExecuteTemplate(w, "tip", resTip)
		if err != nil {
			log.Fatal(err)
		}
	}

	/*
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
	http.HandleFunc("/main", requestMain)
	http.HandleFunc("/upload", requestUpload)
	http.HandleFunc("/md5", requestMd5)
	http.HandleFunc("/send", requestSendMail)
	http.HandleFunc("/file", requestFile)

	http.HandleFunc("/redis", redisHandle)

	//	dir := "./"
	//	http.Handle("/", http.FileServer(http.Dir(dir)))

	router := jas.NewRouter(new(Hello), new(Weixin), new(Baidu))
	router.BasePath = "/"
	//	fmt.Println(router.HandledPaths(true))
	//output: `GET /v1/hello`
	http.Handle(router.BasePath, router)

	err = http.ListenAndServe(":80", nil)
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

var html1 string = `{{define "main"}}{{/* 文件状态页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p>
        <form>
            <tr>使用说明：懒得写了！一个简易的邮箱抓取服务，上传的文件目标用户为男，则命名成”male.txt“，为女则为”female.txt“</tr><br>
         </form>
    </p>
    <p>
        <form method="post" action="/upload" enctype="multipart/form-data">
            <input type="file" id="file" name="file"/>
            <input type="submit" value="上传并抓取邮箱"/>
         </form>
    </p>
    <p>
        <form method="post" action="/md5">
             计算文字md5码:<input type="text" name="name" /><input type="submit" value="计算" />
        </form>
    </p>
    <p> 群发邮件 (说明1：密钥(zhd), 数量(10~100), 标题自定义, 内容支持富文本(可用http://www.vemmis.com/bjq/index.html这个地址编辑))</br>
        (说明2：目前不支持默认参数发送(默认是我的邮箱)，你想用什么邮箱，自己去查看smtp的服务器地址和密码)
        <form method="post" action="/send">
             <tr>密钥：<input type="text" name="key" /> 数量：<input type="text" name="nums" /></tr></br>
             <tr>标题：<input type="text" name="title" /> 内容：<input type="text" name="content" /></tr></br>
             <tr>发送方参数（可选）</tr></br>
             <tr>邮箱地址：<input type="text" name="user" /></tr></br>
             <tr>邮箱smtp服务地址：<input type="text" name="host" /></tr></br>
             <tr>smtp密码：<input type="text" name="passwd" /></tr></br>
             <tr>接收者邮箱序列（格式：123@qq.com;456@qq.com）<input type="text" name="recvs" /></tr></br>
             <tr><input type="submit" value="群发" /></tr>
        </form>
    </p>
</body>
</html>{{end}}
{{define "tip"}}{{/* 文件状态页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p>
        <a href="/main"><input type="submit" value="返回"/></a>
    </p>
    <table>
    <tr>
        <td>{{.}}</td>
    </tr>
    </table>
</body>
</html>{{end}}
{{define "tipary"}}{{/* 文件状态页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p>
        <a href="/main"><input type="submit" value="返回"/></a>
    </p>
    <table>
{{range .}}
    <tr>
        <td>{{.}}</td>
    </tr>
{{end}}
    </table>
</body>
</html>{{end}}
{{define "file"}}{{/* 文件状态页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <table>
{{range .}}
    <tr>
        <td width="20">{{.Name}}</td><td width="20">{{.Size}}kb</td><td>{{.ModTime}}</td>
        <td><a href="/{{.Name}}">下载</a></td>
    </tr>
{{end}}
    </table>
</body>
</html>{{end}}`
