package main

import (
	"fmt"
	"github.com/sqeven/go-ai/vision"
	"github.com/sqeven/go-ai/vision/ocr"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

var client *ocr.OCRClient

func init() {
	client = ocr.NewOCRClient(APIKEY, APISECRET)
}

func main() {
	GeneralRecognizeBasic()
	GeneralRecognizeEnhanced()
}

func GeneralRecognizeBasic() {
	rs, err := client.GeneralRecognizeBasic(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),       //是否检测图像朝向，默认不检测
		ocr.DetectLanguage(),        //是否检测语言，默认不检测。
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
		ocr.WithProbability(),       //是否返回识别结果中每一行的置信度
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}

func GeneralRecognizeEnhanced() {

	resp, err := client.GeneralRecognizeEnhanced(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),
		ocr.DetectLanguage(),
		ocr.LanguageType("CHN_ENG"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())

}
