package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetData() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// MD5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// string转int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// string转int
func String(num int) string {
	n := strconv.Itoa(num)
	return n
}
