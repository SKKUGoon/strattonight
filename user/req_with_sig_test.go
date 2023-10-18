package user_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

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
	sig    string

	myUrl          *url.URL
	myUrlComponent url.Values

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
	if ms > 5000 {
		sig.body["recvWindow"] = "5000"
		return
	}
	sig.body["recvWindow"] = strconv.Itoa(ms)
}

func (sig *BSigRequest) GenSignature() string {
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
		fmt.Printf("DEBUG: signature generated %s\n", signature)
	}

	sig.myUrlComponent.Add("signature", signature)
	return signature
}

func (sig *BSigRequest) GenSignatureABC() string {
	// Add timestamp
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	sig.body["timestamp"] = timestamp

	// Generate signature in alphabetical order
	// Alphabetically sorted parameters are needed to create websocket requests
	keys := make([]string, 0, len(sig.body))
	for k := range sig.body {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedParams := []string{}
	for _, k := range keys {
		if k == "signature" {
			continue
		}
		sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", k, sig.body[k]))
	}

	// Generate the signature using HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(API_PRV))
	mac.Write([]byte(strings.Join(sortedParams, "&")))
	signature := hex.EncodeToString(mac.Sum(nil))
	return signature
}

func (sig *BSigRequest) MakeRequest() {
	sig.myUrl.RawQuery = sig.myUrlComponent.Encode()

	if sig.debug {
		fmt.Printf("DEBUG: URL %s\n", sig.myUrl.String())
	}
}

func TestBSigRequest_MakeRequest(t *testing.T) {
	// Define test data
	var s BSigRequest
	s.New("test", "userDataStream")
	s.debug = true

	// Add body
	s.AddBody("symbol", "BTCUSDT")
	s.AddBody("side", "BUY")
	s.AddBody("type", "LIMIT")
	s.AddBody("quantity", "1")

	s.AddRecvWindow(99999)

	// Create secure signature
	s.GenSignature()

	// Create request
	s.AddSecureKey(API_PUB)
	s.MakeRequest()
}
