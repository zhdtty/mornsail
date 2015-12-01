package console

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	//	"runtime/pprof"
	"driver"
	"io"
)

func Init() {
	fmt.Println("console init")
	go webUI()
}

func redisHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, fmt.Sprintf("Cur redis actives : %d", driver.RedisPool.ActiveCount()))
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
	err := http.ListenAndServe(":5555", nil)
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
