package config

import "sync"

const OcrApi = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic";

const ApiKey = "fGRChj2rPl5rPPekw05uGB5y";

const SecretKey = "MuLcPEZNnQ6mA3Wx5NPipxbNWXL2F9pR";

const TokenUrl = "https://aip.baidubce.com/oauth/2.0/token"

const TokenFile = "c:/screenshot/token.txt"

const BlockImg = "c:/screenshot/screenshot_block.png"

var WaitGroup sync.WaitGroup

var AttentionWord = []string{"是错", "错误", "没有", "不是", "不能", "不对", "不属于", "不可以", "不正确", "不提供", "不包含", "不包括", "不存在", "不经过", "不可能", "不匹配"}

