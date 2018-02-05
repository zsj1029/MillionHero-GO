package baidu

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"regexp"
	"net/url"
	"strings"
	. "github.com/MillionHero-GO/config"
)

//直接搜索问题处理返回值
func SearchQ(qa *QA) {
	//start := float64(time.Now().UnixNano())
	fmt.Println(qa.Question)

	re := regexp.MustCompile(`^\d+[.]{0,1}`)//去除序号
	qa.Question = re.ReplaceAllString(qa.Question,"")
	//qa.Question = "妯娌"
	//fmt.Println("https://www.baidu.com/s?wd="+url.PathEscape(qa.Question))
	doc, err := goquery.NewDocument("http://www.baidu.com/s?wd="+url.PathEscape(qa.Question))
	if err != nil {
		log.Fatal(err)
	}
	//百度百科
	baiKe := doc.Find("#content_left .c-container").First().Text()
	if strings.Contains(baiKe,"百度百科") {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		baiKe = doc.Find("#content_left .c-container .c-span-last p").First().Text()
		fmt.Println("百科：" + strings.TrimSpace(baiKe))
	}
	//最佳答案
	re = regexp.MustCompile(`[更多关于].*`)//去除更多问题
	doc.Find("#content_left .c-container .c-abstract").Each(func(i int, s *goquery.Selection) {
		zuijia := strings.TrimSpace(s.Text())

		if strings.Contains(zuijia,"最佳答案"){
			zuijia = re.ReplaceAllString(zuijia,"")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(zuijia)
		}
	})

	leftContent := doc.Find("#content_left").First().Text()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~匹配统计")
	for _,value := range qa.Answers{
		countNum := strings.Count(leftContent,value.Words)
		fmt.Printf("|%-10s|----> %d\n",value.Words, countNum)
	}
	//end := float64(time.Now().UnixNano())
	//useTime := (end-start)/1000000000
	//fmt.Printf("\n搜索时间：%.3f秒\n",useTime)
}


//搜索问题加答案处理返回值
func SearchQA(q string,a string)  {
	doc, err := goquery.NewDocument("http://www.baidu.com/s?wd="+url.PathEscape(q+" "+a))
	if err != nil {
		log.Fatal(err)
	}

	result := doc.Find(".nums").First().Text()
	//re := regexp.MustCompile(`\d+[,]*`)//取出数字
	//result = re.FindString(result)
	strArr := strings.Split(result,"约")
	fmt.Printf("|%-10s|----> %-10s\n",a,strArr[1])
	//fmt.Println(strArr[1])
	WaitGroup.Done()
}
