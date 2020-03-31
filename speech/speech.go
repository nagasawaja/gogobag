package speech

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"io/ioutil"
	"strconv"
)

// baidu account struct
type BaiduAccountStruct struct {
	AppKey    string
	SecretKey string
	Token     *responseTokenStruct
}

// send baidu speech request struct
type ParamsSpeechRequestStruct struct {
	FileName string
	Format   string
	Rate     int64
}

// baidu sppech response struct
type responseBaiduStruct struct {
	CorpusNo string   `json:"corpus_no,omitempty"`
	ErrNo    int64    `json:"err_no"`
	ErrMsg   string   `json:"err_msg"`
	Sn       string   `json:"sn"`
	Result   []string `json:"result"`
}

// URL
const UrlStandardPattern = "http://vop.baidu.com/server_api"

// new baidu speech
func NewBaiDu(appKey, secretKey string) (*BaiduAccountStruct, error) {
	var account = BaiduAccountStruct{
		AppKey:    appKey,
		SecretKey: secretKey,
		Token:     &responseTokenStruct{},
	}
	// return
	return &account, nil
}

// document https://ai.baidu.com/ai-doc/SPEECH/Vk38lxily
// standard version
func (account *BaiduAccountStruct) StandardPattern(params *ParamsSpeechRequestStruct) (*responseBaiduStruct, error) {
	// read file
	fileByte, err := ioutil.ReadFile(params.FileName)
	if err != nil {
		return nil, err
	}
	// build request
	contentType := "audio/" + params.Format + "; rate=" + strconv.FormatInt(params.Rate, 10)
	// send request
	response, err := account.sendRequest(UrlStandardPattern, contentType, fileByte)
	if err != nil {
		return nil, err
	}
	baiduResponse := responseBaiduStruct{}
	err = response.ToJSON(&baiduResponse)
	if err != nil {
		return nil, err
	}
	if baiduResponse.ErrNo != 0 {
		// error
		if baiduResponse.ErrNo == 3302 {
			// set token invalid
			account.Token.ExpiresAt = 1
		}
		return nil, errors.New("baiduResponse err_msg:" + baiduResponse.ErrMsg)
	}
	return &baiduResponse, nil
}

func (account *BaiduAccountStruct) sendRequest(url string, contentType string, fileByte []byte) (*req.Resp, error) {
	header := req.Header{
		"Content-Type": contentType,
	}
	token, err := account.getToken()
	if err != nil {
		return nil, errors.New("get token err:" + err.Error())
	}
	url = fmt.Sprintf(url + "?cuid=qweqwe&token=" + token)
	response, err := req.Post(url, header, fileByte)
	if err != nil {
		return nil, err
	}
	return response, nil
}
