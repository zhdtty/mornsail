package tool

import (
	"bufio"
	"driver"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"html"
	"io"
	"io/ioutil"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"sync"
	"time"
	"timer"
)

var REDIS_MAIL_TITLE_PREFIX_TABLE string = "mail_title_prefix_set"      //标题额外追加的前缀
var REDIS_SEND_MAIL_CONFIG_TABLE string = "mail_config_hash"            //发送方邮箱配置集合
var REDIS_MAIL_CONFIG_INVALID_TABLE string = "mail_config_invalid_hash" //非法发送方邮箱集合
var REDIS_MAIL_HAS_SEND_TABLE string = "mail_has_send_set"              //已发送过邮件的接收方邮箱集合
var REDIS_MAIL_SEND_INVALID_TABLE string = "mail_send_invalid_set"      //非法的接收方邮箱集合
var REDIS_MAIL_ERROR_LOG = "mail_error_log_hash"                        //错误日志
var REDIS_MAIL_INVALID_ERROR_LOG = "mail_invalid_error_log_hash"        //非法错误日志
var SEND_MAIL_FREQUENCY int = 6                                         //6 f per second //Unused
var SEND_MAIL_INTERVAL int = 10                                         //5 second
var SEND_MAIL_COUNT_PER_SENDER int = 80                                 //单个账号最大邮件数量

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

const (
	ERROR_NORMAL = 1 + iota
	ERROR_INVALID
)

func getCurrentTimeString() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, min, sec)
}

func addErrorToRedis(from string, to string, errStr string, errType int) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	tablekey := REDIS_MAIL_ERROR_LOG + "." + from
	if errType == ERROR_INVALID {
		tablekey = REDIS_MAIL_INVALID_ERROR_LOG + "." + from
	}
	value := "[" + getCurrentTimeString() + "] " + errStr
	_, err := conn.Do("hset", tablekey, to, value)
	if err != nil {
		return false
	}
	return true
}

func addToHasSendMailSet(mailAdr string) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("sadd", REDIS_MAIL_HAS_SEND_TABLE, mailAdr)
	if err != nil {
		return false
	}
	return true
}

func addToRecvInvalidMailSet(srcTableKey string, mailAdr string) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	_, _ = conn.Do("srem", srcTableKey, mailAdr)
	_, err := conn.Do("sadd", REDIS_MAIL_SEND_INVALID_TABLE, mailAdr)
	if err != nil {
		return false
	}
	return true
}

func addToMailConfigInvalidSet(srcTableKey string, mailAdr string, mailAuth string) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	_, _ = conn.Do("hdel", srcTableKey, mailAdr)
	_, err := conn.Do("hset", REDIS_MAIL_CONFIG_INVALID_TABLE, mailAdr, mailAuth)
	if err != nil {
		return false
	}
	return true
}

func getMailTitlePrefixFromRedis() (string, error) {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	prefix, err := redis.String(conn.Do("srandmember", REDIS_MAIL_TITLE_PREFIX_TABLE))
	if err != nil {
		return "", err
	}
	return prefix, nil
}

func getMailConfigFromRedis() (map[string]string, error) {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
			return
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	m, err := redis.StringMap(conn.Do("hgetall", REDIS_SEND_MAIL_CONFIG_TABLE))
	if err != nil {
		return nil, err
	}
	return m, nil
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
	titlePrefix, _ := getMailTitlePrefixFromRedis()
	msg := "To: " + info.To + "\r\n" +
		"From: " + info.Sender.User + "<" + info.Sender.User + ">\r\n" +
		"Subject: " + titlePrefix + info.Subject + "\r\n" + contentType + "\r\n\r\n" + info.Body
	//	fmt.Println(msg)
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

	arr, err := redis.Strings(conn.Do("sdiff", tableKey, REDIS_MAIL_HAS_SEND_TABLE))
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

func TestSendMailRequest(title string, content string, nums int) {
	var mailInfo MailInfo
	mailInfo.Sender.User = "305661987@qq.com"
	mailInfo.Sender.Passwd = "gljhpijservdbidd"
	mailInfo.Sender.Host = "smtp.qq.com:25"
	mailInfo.MailType = "html"
	mailInfo.Subject = title
	mailInfo.Body = html.UnescapeString(content)

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

func SendMailCustomRequest(host, user, passwd, title, content string, recvs []string, nums int) error {
	var mailInfo MailInfo
	mailInfo.Sender.User = user
	mailInfo.Sender.Passwd = passwd
	mailInfo.Sender.Host = host
	mailInfo.MailType = "html"
	mailInfo.Subject = title
	mailInfo.Body = html.UnescapeString(content)

	if len(recvs) <= 0 {
		tableKey := "mail_test"
		recvObjs, err := ReadMailListFromRedis(tableKey)
		if err != nil {
			fmt.Println("Invalid table key :", tableKey)
			return err
		}
		recvs = recvObjs
	}
	for i := 0; i < len(recvs); i += 3 {
		tos := recvs[i:min(len(recvs), i+3)]
		to := strings.Join(tos, ";")
		mailInfo.To = to
		err := sendMail(mailInfo)
		if err != nil {
			fmt.Println("send mail error :", err)
		}
	}
	return nil
}

func ReleaseSendMailRequest(title string, content string, tablekey string, nums int) {
	senderMap, err := getMailConfigFromRedis()
	if err != nil {
		fmt.Println("Can't read any sender mail config from redis")
		return
	}
	recvAry, err := ReadMailListFromRedis(tablekey)
	if err != nil {
		fmt.Println("Invalid table key :", tablekey)
		return
	}

	i := 0
	sendCnt := 0
	for k, v := range senderMap {
		senderAry := strings.Split(v, ";")
		if len(senderAry) != 2 {
			continue
		}
		var mailInfo MailInfo
		mailInfo.Sender.User = k
		mailInfo.Sender.Passwd = senderAry[0]
		mailInfo.Sender.Host = senderAry[1]
		mailInfo.MailType = "html"
		mailInfo.Subject = title
		mailInfo.Body = html.UnescapeString(content)

		for i < len(recvAry) && i < nums {
			mailInfo.To = recvAry[i]
			fmt.Println("Release send mail")
			err = sendMail(mailInfo)
			fmt.Println("Release send mail over")
			if err != nil {
				fmt.Println("send mail error :", err)
				errStr := fmt.Sprintf("%v", err)
				if strings.Contains(errStr, "Mailbox not found") || strings.Contains(errStr, "Invalid User") {
					addToRecvInvalidMailSet(tablekey, mailInfo.To)
				} else if strings.Contains(errStr, "535") {
					//maybe err1: "no route to host" err2: "connection timed out"
					addToMailConfigInvalidSet(REDIS_SEND_MAIL_CONFIG_TABLE, k, v)
					addErrorToRedis(k, mailInfo.To, errStr, ERROR_INVALID)
					break
				} else {
					addErrorToRedis(k, mailInfo.To, errStr, ERROR_NORMAL)
				}
			} else {
				addToHasSendMailSet(mailInfo.To)
				sendCnt++
				if sendCnt%1 == 0 {
					fmt.Println("Send success cnt :", sendCnt)
				}
			}
			i++
			if i%SEND_MAIL_FREQUENCY == 0 {
				break
			}
		}
	}
}

//--------------------------------------定时推送----------------------------------------
type MailPerSend struct {
	sync.Mutex
	SendMap  map[string]string
	SendFlag map[string]int
	RecvAry  []string
	Offset   int

	Title    string
	Content  string
	Tablekey string
	Nums     int
	Count    int
}

func (this *MailPerSend) GetRandSender() (string, string) {
	this.Lock()
	defer this.Unlock()
	sendCnt := len(this.SendMap)
	index := rand.Intn(sendCnt)
	i := 0
	for k, v := range this.SendMap {
		if i == index {
			return k, v
		}
		i++
	}
	return "", ""
}

func (this *MailPerSend) AddSenderCnt(sendUser string) {
	this.Lock()
	defer this.Unlock()
	if val, ok := this.SendFlag[sendUser]; ok {
		if val >= SEND_MAIL_COUNT_PER_SENDER {
			delete(this.SendMap, sendUser)
		} else {
			this.SendFlag[sendUser] = val + 1
		}
		return
	}
	this.SendFlag[sendUser] = 1
}

func (this *MailPerSend) DelSender(sendUser string) {
	this.Lock()
	defer this.Unlock()
	if _, ok := this.SendMap[sendUser]; ok {
		delete(this.SendMap, sendUser)
	}
}

func (this *MailPerSend) GetRecv() string {
	this.Lock()
	defer this.Unlock()
	recvCnt := len(this.RecvAry)
	if this.Offset >= recvCnt {
		return ""
	}
	return this.RecvAry[this.Offset]
}

func (this *MailPerSend) AddOffset() {
	this.Lock()
	defer this.Unlock()
	this.Offset++
}

func (this *MailPerSend) Init(sendMap map[string]string, recvAry []string) {
	this.Lock()
	defer this.Unlock()
	this.SendFlag = make(map[string]int)
	this.Offset = 0
	this.SendMap = sendMap
	this.RecvAry = recvAry
}

func NewMailPerSend() *MailPerSend {
	mps := &MailPerSend{
		SendMap:  make(map[string]string),
		SendFlag: make(map[string]int),
		Offset:   0,
		Title:    "",
		Content:  "",
		Tablekey: "",
		Nums:     0,
		Count:    0,
	}
	return mps
}

var mailPerSend *MailPerSend = NewMailPerSend()

func doSendMail() {
	ret := 0
	sendUser, sendHostPwd := mailPerSend.GetRandSender()
	if sendUser == "" {
		ret = 2
	}
	senderAry := strings.Split(sendHostPwd, ";")
	if len(senderAry) != 2 {
		ret = 1
		mailPerSend.DelSender(sendUser)
	}
	recv := mailPerSend.GetRecv()
	if recv == "" {
		ret = 2
	}
	if ret == 0 {
		var mailInfo MailInfo
		mailInfo.Sender.User = sendUser
		mailInfo.Sender.Passwd = senderAry[0]
		mailInfo.Sender.Host = senderAry[1]
		mailInfo.MailType = "html"
		mailInfo.Subject = mailPerSend.Title
		mailInfo.Body = html.UnescapeString(mailPerSend.Content)

		mailInfo.To = recv
		err := sendMail(mailInfo)
		if err != nil {
			fmt.Println("send mail error :", err)
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "Mailbox not found") || strings.Contains(errStr, "550 Invalid User") {
				addToRecvInvalidMailSet(mailPerSend.Tablekey, mailInfo.To)
				mailPerSend.AddOffset()
			} else if strings.Contains(errStr, "535") {
				//maybe err1: "no route to host" err2: "connection timed out"
				addToMailConfigInvalidSet(REDIS_SEND_MAIL_CONFIG_TABLE, sendUser, sendHostPwd)
				addErrorToRedis(sendUser, mailInfo.To, errStr, ERROR_INVALID)
				mailPerSend.DelSender(sendUser)
			} else {
				addErrorToRedis(sendUser, mailInfo.To, errStr, ERROR_NORMAL)
			}
		} else {
			addToHasSendMailSet(mailInfo.To)
			mailPerSend.AddSenderCnt(sendUser)
			mailPerSend.AddOffset()
			mailPerSend.Count += 1

			fmt.Println("[", sendUser, " to ", mailInfo.To, "] success ! cur count :", mailPerSend.Count)
		}
	}

	if ret == 2 || mailPerSend.Count >= mailPerSend.Nums {
		fmt.Println("Timer send mail over")
		return
	}

	timer.SvrTimer.AddIntervalTimer(SEND_MAIL_INTERVAL, func() {
		doSendMail()
	}, false)
}

func TimerSendMailRequest(title string, content string, tablekey string, nums int) error {
	senderMap, err := getMailConfigFromRedis()
	if err != nil {
		fmt.Println("Can't read any sender mail config from redis")
		return errors.New("Can't read any sender mail config from redis")
	}
	recvAry, err := ReadMailListFromRedis(tablekey)
	if err != nil {
		fmt.Println("Invalid table key :", tablekey)
		return errors.New("Invalid table key ")
	}

	mailPerSend.Init(senderMap, recvAry)
	mailPerSend.Title = title
	mailPerSend.Content = content
	mailPerSend.Tablekey = tablekey
	mailPerSend.Nums = nums
	mailPerSend.Count = 0

	doSendMail()

	return nil
}
