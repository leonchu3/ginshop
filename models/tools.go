package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	. "ginshop/models/go_image"
	"html/template"
	"io"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
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

// int转string
func String(num int) string {
	n := strconv.Itoa(num)
	return n
}

func Sub(a int, b int) int {
	return a - b
}

// 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// 通过列获取值 (反射)
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

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
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName //10表示十进制

	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}

// 生成商品缩略图
func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)
	ThumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",")

	thumbnailSizeSlice := strings.Split(ThumbnailSize, ",")
	//static/upload/tao_400.png
	//static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		savepath := filename + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		w, _ := Int(thumbnailSizeSlice[i])
		err := ThumbnailF2F(filename, savepath, w, w)
		if err != nil {
			fmt.Println(err) //写个日志模块  处理日志
		}
	}

}
