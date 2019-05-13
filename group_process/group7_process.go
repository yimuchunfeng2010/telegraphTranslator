package group_process

import (
	"strings"
)

// 编组7 航空器识别标志
// [A]/[B  C]

// A表示航班，最后一位为字母时，表示补班
var ADataMap = map[string]string{
	"Z": "0",
	"Y": "1",
	"X": "2",
	"W": "3",
	"V": "4",
	"U": "5",
	"T": "6",
	"S": "7",
	"R": "8",
	"Q": "9",
}

func GetGroup7Info(message string) (data string, err error) {
	//fmt.Println("Group7Info：", message)
	//defer func(){
	//	fmt.Println("Group7Info resp ", data)
	//}()
	// 编组7以"/"为分隔符
	messageArr := strings.Split(message, "/")

	// 解析航班号
	flight := messageArr[0]
	lastChar := flight[len(flight)-1:]
	lastDigit := ""
	if lastChar >= "Q" && lastChar <= "Z" {
		lastDigit = ADataMap[lastChar]
		flight = "补班" + flight[:len(flight)-1]
		flight += lastDigit
	}

	additionalInfo := ""
	if len(messageArr) == 2 {
		// todo 目前只处理二次码场景
		switch messageArr[1][0:1] {
		case "A": // 数据项C为二次码
			additionalInfo = "二次码：" + messageArr[1][1:]
		}
		data = "航班: " + flight + ", " + additionalInfo + "\n"
	} else {
		data = "航班: " + flight + "\n"
	}

	return

}
