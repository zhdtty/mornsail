package tool

import (
	"driver"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	regular_mail = "([0-9]|[a-z]){6,15}@([a-z]|[0-9]){2,5}.com"
)

var MailRegex *regexp.Regexp = regexp.MustCompile(regular_mail)

func PushSingleMailToRedis(mailAdr string, tableKey string) {
	defer func() {
		msg := recover()
		if msg != nil {
			panic(msg)
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("sadd", tableKey, mailAdr)
	if err != nil {
		panic(err)
	}
}

func PushMultiMailsToRedis(mails []string, tableKey string) {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Panic error :", msg)
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	args := make([]interface{}, len(mails)+1)
	args[0] = tableKey
	for i, v := range mails {
		args[i+1] = v
	}
	_, err := conn.Do("sadd", args...)
	if err != nil {
		panic(err)
	}
}

func ReadFileMails(filename string, tablekey string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, errors.New("Invalid file")
	}
	mailAry := MailRegex.FindAllString(string(content), -1)
	if len(mailAry) <= 0 {
		return []string{}, errors.New("Filename not get any mails")
	}
	PushMultiMailsToRedis(mailAry, tablekey)
	return mailAry, nil
}

var SMTP_163_ADDRESS string = "smtp.163.com:25"

func ReadFileSendMail(filename string, tablekey string) {
	defer func() {
		msg := recover()
		if msg != nil {
			panic(msg)
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Invalid file")
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, v := range lines {
		mails := strings.Split(v, "----")
		if len(mails) != 2 {
			continue
		}
		_, _ = conn.Do("hset", REDIS_SEND_MAIL_CONFIG_TABLE, mails[0], mails[1]+";"+SMTP_163_ADDRESS)
	}
}
