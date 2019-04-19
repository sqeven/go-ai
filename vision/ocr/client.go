package ocr

import "github.com/sqeven/go-ai"

type OCRClient struct {
	*gosdk.Client
}

func NewOCRClient(apiKey, secretKey string) *OCRClient {
	return &OCRClient{
		Client: gosdk.NewClient(apiKey, secretKey),
	}
}
