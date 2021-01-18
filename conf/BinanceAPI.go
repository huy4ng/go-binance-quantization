package conf

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/url"
	"time"
)

type BianceAPI struct {
	BaseURL    string
	BaseURL_V3 string
	PublicURL  string
	key        string
	secret     string
}

func (api *BianceAPI) Init(key string, secret string) {
	api.BaseURL = "https://www.binance.com/api/v1"
	api.BaseURL_V3 = "https://api.binance.com/api/v3"
	api.PublicURL = "https://www.binance.com/exchange/public/product"
	api.key = key
	api.secret = secret
}

func (api *BianceAPI) Ping() *http.Response {
	var path string = fmt.Sprintf("%s/ping", api.BaseURL_V3)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(path)
	if err != nil {
		panic(err)
	}
	return resp
}

func (api *BianceAPI) GetTicketPrice(market string) float32 {
	//path := fmt.Sprintf("%s/ticker/price", api.BaseURL_V3)
	//client := &http.Client{Timeout: 5 * time.Second}
	////request, err := http.NewRequest("GET", path, nil)
	//if err != nil{
	//	panic(err)
	//}
	//request.Header.Add()

	return 0
}

func (api *BianceAPI) get_ticker_24hour() {

}

func (api *BianceAPI) getNoSign(path string, temParams []byte) {
	var param interface{}
	err := json.Unmarshal(temParams, &param)
	if err != nil {
		panic(err)
	}
	//query := url.QueryEscape(param[])
}
