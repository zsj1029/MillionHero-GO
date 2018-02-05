package config

import "sync"

const OcrApi = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic";

const ApiKey = "fGRChj2rPl5rPPekw05uGB5y";

const SecretKey = "MuLcPEZNnQ6mA3Wx5NPipxbNWXL2F9pR";

const TokenUrl = "https://aip.baidubce.com/oauth/2.0/token"

const TokenFile = "c:/screenshot/token.txt"

const BlockImg = "c:/screenshot/screenshot_block.png"

var WaitGroup sync.WaitGroup
