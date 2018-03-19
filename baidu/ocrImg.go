package baidu

import (
	"io/ioutil"
	"encoding/base64"
	. "github.com/zsj1029/MillionHero-GO/config"
	"net/url"
	"net/http"
	"strings"
	. "github.com/zsj1029/MillionHero-GO/utils"
	"encoding/json"
	"log"
	"regexp"
)

type Image2Text struct {
	LogID          int     `json:"log_id"`
	WordsResult    []Words `json:"words_result"`
	WordsResultNum int     `json:"words_result_num"`
	ErrorCode      int64   `json:"error_code"`
	ErrorMsg       string  `json:"error_msg"`
}

type QA struct {
	Question string
	Answers  []Words
}

type Words struct {
	Words string `json:"words"`
}

func GetImageText(AccessToken string) Image2Text {

	byteImg, _ := ioutil.ReadFile(BlockImg)
	//encode := make([]byte,500000)//byte初始化，encode base64后的图片
	imgBase64 := base64.StdEncoding.EncodeToString(byteImg)

	form := url.Values{}

	form.Add("access_token", AccessToken)
	form.Add("image", imgBase64)

	params := form.Encode()

	resp, err := http.Post(OcrApi, "application/x-www-form-urlencoded", strings.NewReader(params))

	HandleError(err);
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	ack := Image2Text{}

	err = json.Unmarshal(body, &ack)
	HandleError(err)

	if ack.ErrorCode > 0 {
		log.Println(ack.ErrorMsg)
	}
	//fmt.Printf("%#v", text)
	return ack
}

/**
	问题区分答案处理
 */
func GetQA(ack *Image2Text) QA {

	qa := QA{}
	findQues := false
	//ack.WordsResult[2].Words = "A："+ack.WordsResult[2].Words
	re := regexp.MustCompile(`\w+[:|：]+`)
	for _, value := range ack.WordsResult {
		if findQues == false {
			qa.Question += value.Words
			findQues,_ = regexp.MatchString(`\?|？`,qa.Question)
		}else {
			value.Words = strings.Replace(value.Words,"《","",1)
			value.Words = strings.Replace(value.Words,"》","",1)
			value.Words = re.ReplaceAllString(value.Words,"")
			//fmt.Println(value.Words)
			qa.Answers = append(qa.Answers,value)
		}
	}
	return qa

}
