package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func UploadImg(c *gin.Context, picName string) (string, error) {
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3创建图片保存目录
	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, os.ModePerm)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}
	fileName := strconv.FormatInt(GetUnix(), 10) + extName //10表示十进制

	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}
