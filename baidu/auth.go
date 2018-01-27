package baidu

import (
	"fmt"
	. "github.com/MillionHero-GO/config"
	. "github.com/MillionHero-GO/utils"
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
var filename = "c:/screenshot/token.txt"

func GetAuth() (Token,error) {

	var Token = Token{}
	//先从文件加载Token
	Token,err := LoadToken()
	if err == nil && Token.ExpiresIn > time.Now().Unix() {
		return Token,nil
	}

	url := fmt.Sprintf(Token_url+"?grant_type=%s&client_id=%s&client_secret=%s",
		"client_credentials",ApiKey,SecretKey)

	resp,err := http.Get(url)
	HandleError(err);

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	jsonStr := string(body)
	fmt.Println("jsonStr", jsonStr)
	err = json.Unmarshal(body, &Token)

	if err != nil {
		log.Println("json.Unmarshal failed", "err", err, "url", url, "method", "Get", "body", string(body))
	}

	Token.ExpiresIn += time.Now().Unix()

	tokenText,err := json.Marshal(&Token)

	ioutil.WriteFile(filename,tokenText,0666)

	return Token,err
}

func LoadToken()  (Token,error){

	var Token = &Token{}

	_,err :=os.Stat(filename);
	if err != nil {
		file,err := os.Create(filename)
		defer file.Close()
		return *Token,err
	}
	tokenText,_ := ioutil.ReadFile(filename)
	json.Unmarshal([]byte(tokenText),Token)

	if err != nil {
		log.Println("json.Unmarshal failed", "err", err)
	}
	return *Token,err
}

