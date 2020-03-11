package idcard

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomIdCard() string {
	// 行政区号
	addrCode := getAddrCode()
	// 生日日期
	birthdayCode := getBirthday()
	// 两位随机数
	twoRndCode := getTwoRandCode()
	// 性别的特殊数
	sexCode := getSexyCode()
	idCard := addrCode + birthdayCode + twoRndCode + sexCode
	// 校验位数
	verifyCode := getVerifyCode(idCard)
	idCard = idCard + verifyCode
	return idCard
}

// 随机获取一个行政区号,身份证前6位，即地址码
func getAddrCode() string {
	rand.Seed(time.Now().UnixNano()) // initialize global pseudo random generator
	return AddrSlice[rand.Intn(len(AddrSlice))]
}

// 随机生成一个生日日期 身份证7到14位，即出生年月日
func getBirthday() string {
	beginDate := "1950-01-01 00:00:00"
	EndDate := "1999-12-31 23:59:59"
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")

	// format timestamp
	beginTimestamp, _ := time.ParseInLocation(timeLayout, beginDate, loc)
	endTimestamp, _ := time.ParseInLocation(timeLayout, EndDate, loc)

	// 设置随机种子数
	rand.Seed(time.Now().UnixNano())
	finalTimestamp := rand.Int63n(endTimestamp.Unix()-beginTimestamp.Unix()+1) + beginTimestamp.Unix()
	finalDate := time.Unix(finalTimestamp, 0).Format("20060102")
	return finalDate
}

// 随机生成两个数字 身份证15到16位
func getTwoRandCode() string {
	rand.Seed(time.Now().UnixNano())
	firstNum := rand.Int63n(10)
	secondNum := rand.Int63n(10)
	return fmt.Sprintf("%v%v", firstNum, secondNum)
}

// 性别 身份证17位
func getSexyCode() string {
	// python 的算法如下， sex是1或2
	// idCode += str(random.randrange(self.sex, 9, 2))
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Int63n(8)+1)
}

// 校验位 身份证18位
func getVerifyCode(idCard string) string {
	verifyCode := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	Wi := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var codeSum int = 0
	for k, v := range idCard {
		tmpNum, _ := strconv.Atoi(fmt.Sprintf("%c", v))
		codeSum += tmpNum * Wi[k]
	}
	yushu := codeSum % 11
	return verifyCode[yushu]
}

type IdCardStruct struct {
	AddrNum    string
	AddrName   string
	BirthDay   string
	Sex        string
	VerifyCode string
}

// 检查身份证是否合法
func CheckIdCardValid(idCard string) error {
	// init idcard
	idCardByte := []byte(idCard)
	if len(idCardByte) != 18 {
		return errors.New(fmt.Sprintf("idcard len err! idCard len:%v", len(idCardByte)))
	}
	// init struct
	idCardRow := IdCardStruct{}
	// 检测行政区号 前6位
	if v, ok := AddrMap[string(idCardByte[0:6])]; !ok {
		return errors.New(fmt.Sprintf("addrCode err! addrCode:%v", string(idCardByte[0:6])))
	} else {
		idCardRow.AddrName = v
		idCardRow.AddrNum = string(idCardByte[0:6])
	}
	// 检测生日日期 身份证7到14位
	timeLayout := "20060102"
	loc, _ := time.LoadLocation("Local")
	// format timestamp
	birthDayTimestamp, _ := time.ParseInLocation(timeLayout, string(idCardByte[6:14]), loc)
	if birthDayTimestamp.Unix() < -2209017943 {
		return errors.New(fmt.Sprintf("birthday err! birthday:%v", string(idCardByte[6:14])))
	} else {
		idCardRow.BirthDay = birthDayTimestamp.Format("2006-01-02")
	}
	// 检测随机位 身份证15到16位,两个数字
	_, err := strconv.Atoi(string(idCardByte[14:15]))
	if err != nil {
		return errors.New(fmt.Sprintf("15 code must int! code:%v", string(idCardByte[14:15])))
	}
	_, err = strconv.Atoi(string(idCardByte[15:16]))
	if err != nil {
		return errors.New(fmt.Sprintf("16 code must int! code:%v", string(idCardByte[16:16])))
	}
	// 检测性别
	v, err := strconv.Atoi(string(idCardByte[16:17]))
	if err != nil || v == 0 {
		return errors.New(fmt.Sprintf("sex code must int! code:%v", string(idCardByte[14:15])))
	} else {
		if v%2 == 0 {
			idCardRow.Sex = "女"
		} else {
			idCardRow.Sex = "男"
		}
	}
	// 检测校验位
	verifyCode := getVerifyCode(string(idCardByte[0:17]))
	if string(idCardByte[17]) != verifyCode {
		return errors.New(fmt.Sprintf("18 verify code err! verify code should:%v", verifyCode))
	}

	// 输入返回
	//logrus.Infof("%+v", idCardRow)
	return nil
}

// 检测身份证年龄
func CheckIdCardAge(idCard string, minAge int64) error {
	// init idcard
	idCardByte := []byte(idCard)
	// 获取idcard上的生日
	timeLayout := "20060102"
	loc, _ := time.LoadLocation("Local")
	// format timestamp
	birthDayTimestamp, _ := time.ParseInLocation(timeLayout, string(idCardByte[6:14]), loc)
	// 对比现在的时间戳
	ageTimestamp := time.Now().Unix() - birthDayTimestamp.Unix()
	if ageTimestamp >= (minAge-1)*86400*365 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("idCard age less than minAge! idCard birthday:%v", string(idCardByte[6:14])))
	}
}
