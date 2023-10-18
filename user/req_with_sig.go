package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// To get signed Binance user data using Websocket data stream
// Need to follow a specific process to generate the signature.
// Signature will in turn construct the WebSocket connection URL.

// TODO: change it to .env file

const API_PUB = "OCwrKPU2rkk8gZFIjLkN77GUKz5lI8rBW3FdTe4JCWncABzzWYoVm5FDHkmJCFlE"
const API_PRV = "SvGRkfbByRaufyCqVCBohoFH3mqBMZiMzINg8VH25aXp6KhoMuyi6gaOAd1YfTj2"

type BinanceUserInfo map[string]string
type BinanceUserReq map[string]string

var SourceMap map[string]string = map[string]string{
	"test": "https://api.binance.com/api/v3",
}

var EndpointMap map[string]string = map[string]string{
	"userDataStream": "/userDataStream",
}

type BSigRequest struct {
	header BinanceUserInfo
	body   BinanceUserReq

	myUrl          *url.URL
	myUrlComponent url.Values // Addition to body, it has signature

	debug bool
}

func (sig *BSigRequest) New(source, endpoint string) {
	u, _ := url.Parse(fmt.Sprintf(
		"%s%s",
		SourceMap[source],
		EndpointMap[endpoint],
	))

	sig.header = map[string]string{}
	sig.body = map[string]string{}

	// Save it to the structure
	sig.myUrl = u
	sig.myUrlComponent = url.Values{}
}

func (sig *BSigRequest) NewWss() {
	sig.header = map[string]string{}
	sig.body = map[string]string{}

}

func (sig *BSigRequest) AddSecureKey(key string) {
	// NONE	        Endpoint can be accessed freely.
	// TRADE	    Endpoint requires sending a valid API-Key and signature.
	// MARGIN	    Endpoint requires sending a valid API-Key and signature.
	// USER_DATA	Endpoint requires sending a valid API-Key and signature.
	// USER_STREAM	Endpoint requires sending a valid API-Key.
	// MARKET_DATA	Endpoint requires sending a valid API-Key.
	sig.header["X-MBX-APIKEY"] = key
}

func (sig *BSigRequest) AddBody(key, value string) {
	sig.body[key] = value
}

func (sig *BSigRequest) AddRecvWindow(ms int) {
	if ms > 5000 && !sig.debug {
		sig.body["recvWindow"] = "5000"
		return
	}
	sig.body["recvWindow"] = strconv.Itoa(ms)
}

func (sig *BSigRequest) GenSignature() {
	// Add timestamp
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	sig.body["timestamp"] = timestamp

	// Move from map to url.Values{}
	for k, v := range sig.body {
		sig.myUrlComponent.Add(k, v)
	}

	// Generate the signature using HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(API_PRV))
	mac.Write([]byte(sig.myUrlComponent.Encode()))
	signature := hex.EncodeToString(mac.Sum(nil))

	if sig.debug {
		log.Printf("DEBUG: signature generated %s\n", signature)
	}

	sig.myUrlComponent.Add("signature", signature)
}

func (sig *BSigRequest) MakeGetRequest() {
	sig.myUrl.RawQuery = sig.myUrlComponent.Encode()

	if sig.debug {
		log.Printf("DEBUG: URL %s\n", sig.myUrl.String())
	}

	req, err := http.NewRequest("GET", sig.myUrl.String(), nil)
	if err != nil {
		log.Printf("ERROR: creating request: %v\n", err)
		return
	}

	for k, v := range sig.header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: sending request %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

//func (sig *BSigRequest) MakeWssPacket(toByte bool) {
//	sig.myUrl.R
//}
