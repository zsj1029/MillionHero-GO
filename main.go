package main

import (
	"fmt"
	"time"
	"runtime"
	"os/exec"
	"os"
	"github.com/MillionHero-GO/baidu"
	. "github.com/MillionHero-GO/utils"
	"strconv"
	"image"
	 _ "image/png"
	"image/png"
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
	AccessToken,err := baidu.GetAuth()
	HandleError(err)

	//fmt.Printf("%#v",AccessToken)


	var quote string
	var i = 1
	for true  {
		fmt.Print("按回车键开始识别问题...")
		fmt.Scanf("%s", &quote)
		start := float64(time.Now().UnixNano())

		//screenImg(i)//安卓截屏
		cutImage(99)
		getImageText(AccessToken)

		end := float64(time.Now().UnixNano())
		useTime := (end-start)/1000000000
		fmt.Printf("处理时间：%.3f秒\n",useTime)
		i++
	}

}

func screenImg(i int)  {
	//截图
	var cmd *exec.Cmd
	cmd = exec.Command("cmd","/C","adb shell /system/bin/screencap -p /sdcard/screenshot.png")
	err := cmd.Run()
	HandleError(err)
	//保存本地
	cmd = exec.Command("cmd","/C","adb pull /sdcard/screenshot.png c:/screenshot/screenshot"+strconv.Itoa(i)+".png")
	err = cmd.Run()
	HandleError(err)
}

func cutImage(i int) {
	path := "c:/screenshot/screenshot"+strconv.Itoa(i)+".png";
	//打开图片
	file, err := os.Open(path)
	defer file.Close()
	HandleError(err)

	m, _, err := image.Decode(file)// 图片文件解码
	HandleError(err)

	//fmt.Printf("%v",m.ColorModel())
	img := m.(*image.NRGBA)
	newImg := img.SubImage(image.Rect(75,300, 1020,1220)).(*image.NRGBA)
	imgfile, err := os.Create("c:/screenshot/screenshot_block.png")
	defer imgfile.Close()
	err = png.Encode(imgfile, newImg)
}

func getImageText(AccessToken string)  {

	fmt.Println(AccessToken)

}

