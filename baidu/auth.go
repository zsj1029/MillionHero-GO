package baidu

import (
	"fmt"
	. "github.com/zsj1029/MillionHero-GO/config"
	. "github.com/zsj1029/MillionHero-GO/utils"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"time"
	"os"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Error       string `json:"error"`
	ErrDesc     string `json:"error_description"`
	TS          int64
}

//获取百度token
func GetAuth() (AccessToken string,err error) {

	//先从文件加载Token
	Token,err := LoadToken()
	if err == nil && Token.ExpiresIn > time.Now().Unix() {
		return Token.AccessToken,nil
	}

	url := fmt.Sprintf(TokenUrl+"?grant_type=%s&client_id=%s&client_secret=%s",
		"client_credentials",ApiKey,SecretKey)

	resp,err := http.Get(url)
	HandleError(err);

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//jsonStr := string(body)
	//fmt.Println("jsonStr", jsonStr)
	err = json.Unmarshal(body, &Token)

	if err != nil {
		log.Println("json.Unmarshal failed", "err", err, "url", url, "method", "Get", "body", string(body))
	}

	Token.ExpiresIn += time.Now().Unix()

	tokenText,err := json.Marshal(&Token)

	ioutil.WriteFile(TokenFile,tokenText,0666)

	return Token.AccessToken,err
}

//文件加载token
func LoadToken()  (Token,error){

	var Token = &Token{}

	_,err :=os.Stat(TokenFile);
	if err != nil {
		file,err := os.Create(TokenFile)
		defer file.Close()
		return *Token,err
	}
	tokenText,_ := ioutil.ReadFile(TokenFile)
	json.Unmarshal([]byte(tokenText),Token)

	if err != nil {
		log.Println("json.Unmarshal failed", "err", err)
	}
	return *Token,err
}

