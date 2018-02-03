package baidu

import (
	"io/ioutil"
	"encoding/base64"
	. "github.com/MillionHero-GO/config"
	"fmt"
	"net/url"
	"net/http"
	"strings"
	. "github.com/MillionHero-GO/utils"
	"encoding/json"
	"log"
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

func GetImageText(AccessToken string) bool {

	byteImg, _ := ioutil.ReadFile(BlockImg)
	//encode := make([]byte,500000)//byte初始化，encode base64后的图片
	imgBase64 := base64.StdEncoding.EncodeToString(byteImg)

	form := url.Values{}

	form.Add("access_token", AccessToken)
	form.Add("image", imgBase64)

	params := form.Encode()

	resp, err := http.Post(OcrApi, "application/x-www-form-urlencoded", strings.NewReader(params))

	defer resp.Body.Close()
	HandleError(err);
	body, _ := ioutil.ReadAll(resp.Body)
	ack := Image2Text{}

	err = json.Unmarshal(body, &ack)
	HandleError(err)

	if ack.ErrorCode > 0 {
		log.Println(ack.ErrorMsg)
		return false
	}
	//fmt.Printf("%#v", text)
	//解析问题和答案
	GetQA(ack)
	return true
}

func GetQA(ack Image2Text) {

	qa := QA{}
	//if ack.WordsResultNum <= 4 {
	//	qa.Question = ack.WordsResult[0].Words
	//} else if ack.WordsResultNum == 5 {
	//	qa.Question = ack.WordsResult[0].Words + ack.WordsResult[1].Words
	//} else {
	//	qa.Question = ack.WordsResult[0].Words + ack.WordsResult[1].Words + ack.WordsResult[2].Words
	//}

	for key, value := range ack.WordsResult {
		newMap[key] = value
	}

	fmt.Println(qa)

}
