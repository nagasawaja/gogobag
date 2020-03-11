package speech

import (
	"testing"
)

func TestSpeech(t *testing.T) {
	appKey := ""
	secretKey := ""
	baiduAccount, err := NewBaiDu(appKey, secretKey)
	if err != nil {
		t.Logf("NewBaiDu err:%s", err.Error())
		return
	}

	// m4a
	params := ParamsSpeechRequestStruct{
		FileName: "test_data/woshizhuwoshizhuwoshizhu.m4a",
		Format:   "m4a",
		Rate:     16000,
	}
	responseBaidu, err := baiduAccount.StandardPattern(&params)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("responseBaidu:%+v", responseBaidu.Result[0])

	// silk
	params = ParamsSpeechRequestStruct{
		FileName: "test_data/xinnianhao.pcm",
		Format:   "pcm",
		Rate:     16000,
	}
	responseBaidu, err = baiduAccount.StandardPattern(&params)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("responseBaidu:%+v", responseBaidu.Result[0])
}
