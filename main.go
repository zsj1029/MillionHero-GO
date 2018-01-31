package main

import (
	"fmt"
	"time"
	"runtime"
	"os/exec"
	"os"
	"github.com/MillionHero-GO/baidu"
	. "github.com/MillionHero-GO/utils"
)

func main() {


	if runtime.GOOS != "windows" {
		panic("程序只能运行在windows系统")
	}
	//创建临时目录
	err := os.MkdirAll("c:/screenshot/", 0777)
	HandleError(err)

	//杀死adb相关进程
	//var cmd *exec.Cmd
	//cmd = exec.Command("cmd","/C","taskkill /f /t /im adb.exe")
	//err = cmd.Run()
	//handleError(err)
	//启用adb进程
	cmd := exec.Command("adb","devices")
	out, err := cmd.Output()
	HandleError(err)
	fmt.Print(string(out))

	//百度auth_token
	token,err := baidu.GetAuth()
	HandleError(err)

	fmt.Printf("%#v",token)


	var quote string
	for true  {
		fmt.Print("按回车键开始识别问题...")
		fmt.Scanf("%s", &quote)
		start := float64(time.Now().UnixNano())
		screen_img()//安卓截屏

		cut_image()

		get_image_text();
		end := float64(time.Now().UnixNano())
		//fmt.Printf("处理时间:%v",math.Ceil(end-start)/100000000)
		useTime := (end-start)/1000000000
		fmt.Printf("处理时间：%.3f秒\n",useTime)

	}

}

func screen_img()  {
	//截图
	var cmd *exec.Cmd
	cmd = exec.Command("cmd","/C","adb shell /system/bin/screencap -p /sdcard/screenshot.png")
	err := cmd.Run()
	HandleError(err)
	//保存本地
	cmd = exec.Command("cmd","/C","adb pull /sdcard/screenshot.png c:/screenshot/screenshot.png")
	err = cmd.Run()
	HandleError(err)
}

func cut_image() {

}

func get_image_text()  {

	baidu.GetAuth()
}

