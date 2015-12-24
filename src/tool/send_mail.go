package tool

import (
	"bufio"
	"driver"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
	//	"time"
)

type SenderConfig struct {
	Host   string
	User   string
	Passwd string
}

type MailInfo struct {
	Sender   SenderConfig
	To       string
	MailType string
	Subject  string
	Body     string
}

func sendMail(info MailInfo) error {
	host := strings.Split(info.Sender.Host, ":")
	auth := smtp.PlainAuth("", info.Sender.User, info.Sender.Passwd, host[0])
	var contentType string
	if info.MailType == "html" {
		contentType = "Content-Type: text/html;charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain;charset=UTF-8"
	}
	msg := "To: " + info.To + "\r\n" +
		"From: " + info.Sender.User + "<" + info.Sender.User + ">\r\n" +
		"Subject: " + info.Subject + "\r\n" + contentType + "\r\n\r\n" + info.Body
	recvs := strings.Split(info.To, ";")
	err := smtp.SendMail(info.Sender.Host, auth, info.Sender.User, recvs, []byte(msg))
	return err
}

func ReadEmailContentFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(content), nil
}

func ReadMailListFromRedis(tableKey string) ([]string, error) {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	arr, err := redis.Strings(conn.Do("smembers", tableKey))
	if err != nil {
		panic(err)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
	return arr, nil
}

func ReadMailListFromFile(filename string) ([]string, error) {
	result := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return result, errors.New("Open file failed.")
	}
	defer file.Close()
	bf := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := bf.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				return result, errors.New("ReadLine no finish")
			}
			break
		}
		if isPrefix {
			return result, errors.New("Line is too long")
		}
		str := string(line)
		result = append(result, str)
	}
	return result, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func SendMailForTest() {
	var mailInfo MailInfo
	mailInfo.Sender.User = "305661987@qq.com"
	mailInfo.Sender.Passwd = "gljhpijservdbidd"
	mailInfo.Sender.Host = "smtp.qq.com:25"
	mailInfo.Subject = "测试"
	mailInfo.MailType = "html"
	mailInfo.Body = "<p><strong>阿斯顿发撒旦法稍等<img src=\"http://img.baidu.com/hi/jx2/j_0028.gif\"/></strong><span style=\"font-size: 36px;\"><em><span style=\"text-decoration-line: underline;\"><strong><a href=\"http://www.baidu.com\" target=\"_self\" title=\"哈哈\">阿萨德阿斯蒂芬</a></strong></span></em></span></p>"

	tableKey := "mail_test"
	recvs, err := ReadMailListFromRedis(tableKey)
	if err != nil {
		fmt.Println("Invalid table key :", tableKey)
		return
	}
	for i := 0; i < len(recvs); i += 3 {
		tos := recvs[i:min(len(recvs), i+3)]
		to := strings.Join(tos, ";")
		mailInfo.To = to
		err = sendMail(mailInfo)
		if err != nil {
			fmt.Println("send mail error :", err)
		}
	}
}
