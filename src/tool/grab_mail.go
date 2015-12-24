package tool

import (
	"driver"
	"fmt"
	"io/ioutil"
	"regexp"
)

const (
	regular_mail = "([0-9]|[a-z]){6,15}@([a-z]|[0-9]){2,5}.com"
)

/*
func FindAll(reg *regexp.Regexp, src string) []byte {
	reg := regexp.MustCompile(regular)
	return reg.FindAllString(src, -1)
}
*/
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

func ReadFileMails(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Invalid file ! filename :", filename)
		return
	}
	mailAry := MailRegex.FindAllString(string(content), -1)
	PushMultiMailsToRedis(mailAry, "mail_set")
}
