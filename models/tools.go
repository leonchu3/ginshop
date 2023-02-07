package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetData() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
