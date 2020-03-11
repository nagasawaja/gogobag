# golang baidu speech


base on baidu speech standard pattern.. <br>
just use it.don't worry token expire.. <br>
https://ai.baidu.com/ai-doc/SPEECH/Vk38lxily


## Contents
 - [Quick start](#quick-start)
 - [Converter tool](#converter-tool)
 
## Quick start
```sh
1.open speech_test.go
2.type appkey and secretKey
3.go test -v
```
**Correct Output**
```
$ go test -v
=== RUN   TestSpeech
--- PASS: TestSpeech (1.62s)
    speech_test.go:27: responseBaidu:我是猪，我是猪，我是猪。
    speech_test.go:40: responseBaidu:新年好。
PASS
ok      gogobag/speech  3.542s
```

## Converter tool
```
source:https://github.com/kn007/silk-v3-decoder
```