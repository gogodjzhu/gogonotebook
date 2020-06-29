package main

import (
	"flag"
	"fmt"
	"time"
)

var help *bool
var now *bool
var inputUnix *int64
var inputStr *string
var format *string

func init() {
	help = flag.Bool("h", false, "输出此帮助信息")
	now = flag.Bool("n", false, "当前时间模式")
	format = flag.String("f", "unix", "输出格式,unix=时间戳;str=yyyy-MM-dd HH:mm:ss;自定义格式")
	inputUnix = flag.Int64("iu", -1, "输入时间戳")
	inputStr = flag.String("is", "", "输入字符串")
}

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	var dt = time.Now()
	if !*now {
		if *inputStr != "" && *inputUnix != -1 {
			fmt.Println("inputStr 和 inputUnix 不能同时设置")
			return
		} else if *inputStr != "" {
			if d, err := time.ParseInLocation("2006-01-02 15:04:05", *inputStr, time.Local); err != nil {
				fmt.Println("时间格式化错误")
				return
			} else {
				dt = d
			}
		} else if *inputUnix != -1 {
			dt = time.Unix(*inputUnix, 0)
		} else {
			fmt.Println("请设置inputUnix or inputStr")
			return
		}
	}
	switch *format {
	case "unix":
		fmt.Println(dt.Unix())
	case "str":
		fmt.Println(dt.Format("2006-01-02 15:04:05"))
	default:
		fmt.Println(dt.Format(*format))
	}
	return

}
