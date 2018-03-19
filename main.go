package main

import (
	"fmt"
	"time"
	"runtime"
	"os/exec"
	"os"
	"github.com/zsj1029/MillionHero-GO/baidu"
	. "github.com/zsj1029/MillionHero-GO/utils"
	. "github.com/zsj1029/MillionHero-GO/config"
	"strconv"
	"image"
	 _ "image/png"
	"image/png"
	"bytes"
	"io/ioutil"
)



func main() {
	if runtime.GOOS != "windows" {
		panic("程序只能运行在windows系统")
	}

	//创建临时目录
	err := os.MkdirAll("c:/screenshot/", 0777)
	HandleError(err)


	//杀死adb相关进程
	var cmd *exec.Cmd
	//cmd = exec.Command("cmd","/C","taskkill /f /t /im adb.exe")
	//err = cmd.Run()
	//HandleError(err)
	//启用adb进程
	cmd = exec.Command("adb","devices")
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
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		start := float64(time.Now().UnixNano())

		//screenImg2(i)//安卓截屏
		cutImage2(i)
		getAnswer(AccessToken)

		end := float64(time.Now().UnixNano())
		useTime := (end-start)/1000000000
		fmt.Printf("\n处理时间：%.3f秒\n",useTime)
		i++
	}

}
//原始2次io
func screenImg(i int)  {
	//截图
	start := float64(time.Now().UnixNano())
	var cmd *exec.Cmd
	cmd = exec.Command("cmd","/C","adb shell /system/bin/screencap -p /sdcard/screenshot.png")
	err := cmd.Run()
	HandleError(err)
	//保存本地
	cmd = exec.Command("cmd","/C","adb pull /sdcard/screenshot.png c:/screenshot/screenshot"+strconv.Itoa(i)+".png")
	err = cmd.Run()
	HandleError(err)
	end := float64(time.Now().UnixNano())
	useTime := (end-start)/1000000000
	fmt.Printf("\n截图时间：%.3f秒\n",useTime)
}
//直接截屏，省去文件io
func screenImg2(i int)  {
	start := float64(time.Now().UnixNano())
	cmd := exec.Command("adb", "shell", "screencap", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	x := bytes.Replace(out.Bytes(), []byte("\r\n"), []byte("\n"), -1)
	ioutil.WriteFile("c:/screenshot/screenshot"+strconv.Itoa(i)+".png",x,0666)
	end := float64(time.Now().UnixNano())
	useTime := (end-start)/1000000000
	fmt.Printf("\n截图时间：%.3f秒\n",useTime)
}
//文件读取切图
func cutImage(i int) {
	path := "c:/screenshot/screenshot"+strconv.Itoa(i)+".png";
	//打开图片
	file, err := os.Open(path)
	defer file.Close()
	HandleError(err)

	m, _, err := image.Decode(file)// 图片文件解码
	HandleError(err)

	img := m.(*image.NRGBA)
	newImg := img.SubImage(image.Rect(75,300, 1020,1220)).(*image.NRGBA)
	imgFile, err := os.Create(BlockImg)
	defer imgFile.Close()
	err = png.Encode(imgFile, newImg)
}
//截屏+文件流切图
func cutImage2(i int) {
	start := float64(time.Now().UnixNano())
	cmd := exec.Command("adb", "shell", "screencap", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
	x := bytes.Replace(out.Bytes(), []byte("\r\n"), []byte("\n"), -1)
	go func(i int) {
		ioutil.WriteFile("c:/screenshot/screenshot"+strconv.Itoa(i)+".png",x,0666)
	}(i)
	m, _, err := image.Decode(bytes.NewReader(x))// 图片文件解码
	HandleError(err)

	img := m.(*image.NRGBA)
	newImg := img.SubImage(image.Rect(75,300, 1020,1220)).(*image.NRGBA)
	imgFile, err := os.Create(BlockImg)
	defer imgFile.Close()
	err = png.Encode(imgFile, newImg)
	end := float64(time.Now().UnixNano())
	useTime := (end-start)/1000000000
	fmt.Printf("\n截图时间：%.3f秒\n",useTime)
}

func getAnswer(AccessToken string)  {

	qaText := baidu.GetImageText(AccessToken)

	qa := baidu.GetQA(&qaText)
	//baidu.Knowledge(&qa)
	baidu.SearchQ(&qa)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~数量统计~~~~~~~~~~~~~~~~~~~~~")
	for _,value := range qa.Answers{
		WaitGroup.Add(1)
		go baidu.SearchQA(qa.Question,value.Words)
	}
	WaitGroup.Wait()
}

