package speech

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"sync"
	"time"
)

type responseTokenStruct struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
	Scope         string `json:"scope"`
	// below params exist in error request
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ExpiresAt        int64
}

// URL
const UrlToken = "https://aip.baidubce.com/oauth/2.0/token"

// lock
var tokenLock = sync.Mutex{}

func (account *BaiduAccountStruct) getToken() (string, error) {
	tokenLock.Lock()
	defer tokenLock.Unlock()
	// check token expire
	if time.Now().Unix() <= account.Token.ExpiresAt {
		// token valid
		return account.Token.AccessToken, nil
	}
	responseToken, err := account.getTokenRequest()
	if err != nil {
		return "", err
	}
	account.Token = responseToken
	account.Token.ExpiresAt = time.Now().Unix() + account.Token.ExpiresIn - 3600
	return account.Token.AccessToken, nil
}

// maybe refreshToken is need
func (account *BaiduAccountStruct) refreshToken() (string, error) {
	tokenLock.Lock()
	defer tokenLock.Unlock()
	return "", nil
}

func (account *BaiduAccountStruct) getTokenRequest() (*responseTokenStruct, error) {
	params := req.Param{
		"grant_type":    "client_credentials",
		"client_id":     account.AppKey,
		"client_secret": account.SecretKey,
	}
	response, err := req.Post(UrlToken, "", params)
	if err != nil {
		return nil, err
	}
	responseToken := responseTokenStruct{}
	err = response.ToJSON(&responseToken)
	if err != nil {
		return nil, err
	}
	if responseToken.Error != "" {
		// get token error
		return nil, errors.New(fmt.Sprintf("error:%s;error_description:%s",
			responseToken.Error, responseToken.ErrorDescription))
	}
	return &responseToken, nil
}
