package speech

//
type baiduAccount struct {
	AppKey    string
	SecretKey string
	Token     string
}

// new baidu speech
func NewBaiDu(appKey, secretKey string) (*baiduAccount, error) {
	var account = baiduAccount{}
	// get request token
	// set baiduAccount

	// return
	return &account, nil
}

func (account *baiduAccount) getToken() {

}

// document https://ai.baidu.com/ai-doc/SPEECH/Vk38lxily
// standard version
func (account *baiduAccount) StandardVersion() {

}
