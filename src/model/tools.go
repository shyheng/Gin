package model

import (
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

//获取当前日期
func GetDay() string {
	template := "20060101"
	return time.Now().Format(template)
}
