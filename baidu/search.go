package baidu

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"regexp"
	"net/url"
	"strings"
	. "github.com/MillionHero-GO/config"
	"github.com/yanyiwu/gojieba"
	"github.com/axgle/mahonia"
)

func Knowledge(qa *QA) {
	question := qa.Question
	//中文分词，提取关键词
	var keywords []string
	words := gojieba.NewJieba().Tag(question)
	for _, v := range words {
		if strings.Contains(v, "/n") {
			keywords = append(keywords, strings.Split(v, "/")[0])
		}
	}
	fmt.Println(strings.Join(keywords, " "));
	//搜索题目
	searchURL := fmt.Sprintf("https://zhidao.baidu.com/search?word=%s", url.QueryEscape(strings.Join(keywords, " ")))
	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	s := doc.Find("#wgt-autoask")
	t := s.Find("#wgt-autoask > dl > dt").Text()
	d := s.Find("#wgt-autoask > dl > dd.dd.answer").Text()
	dec := mahonia.NewDecoder("gbk")
	t = dec.ConvertString(t)
	d = dec.ConvertString(d)
	//d = strings.Replace(d, "\n", "", -1)
	//d = strings.Replace(d, "推荐答案", "", -1)
	//d = strings.Replace(d, "[详细]", "", -1)
	fmt.Printf("知识图谱推荐答案：%s%s\n", t, d)
}

//直接搜索问题处理返回值
func SearchQ(qa *QA) {
	fmt.Println(qa.Question)


	re := regexp.MustCompile(`^\d+[.]{0,1}`)//去除序号
	qa.Question = re.ReplaceAllString(qa.Question,"")
	//qa.Question = "妯娌"
	//fmt.Println("https://www.baidu.com/s?wd="+url.PathEscape(qa.Question))
	doc, err := goquery.NewDocument("http://www.baidu.com/s?wd="+url.QueryEscape(qa.Question))
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

	//for _, v := range AttentionWord {
	//	if strings.Contains(qa.Question, v) {
	//		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//		fmt.Printf("请注意题干：%s...\n", v)
	//		break
	//	}
	//}
	leftContent := doc.Find("#content_left").First().Text()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~匹配统计~~~~~~~~~~~~~~~~~~~~~")
	for _,value := range qa.Answers{
		countNum := strings.Count(leftContent,value.Words)
		//length := len(value.Words)
		//fmt.Println(length)
		fmt.Printf("%s\t---->\t%d\n",value.Words, countNum)
	}
}


//搜索问题加答案处理返回值
func SearchQA(q string,a string)  {
	doc, err := goquery.NewDocument("http://www.baidu.com/s?wd="+url.QueryEscape(q+" "+a))
	if err != nil {
		log.Fatal(err)
	}

	result := doc.Find(".nums").First().Text()
	//re := regexp.MustCompile(`\d+[,]*`)//取出数字
	//result = re.FindString(result)
	strArr := strings.Split(result,"约")
	fmt.Printf("%s\t---->\t%s\n",a,strArr[1])
	//fmt.Println(strArr[1])
	WaitGroup.Done()
}
