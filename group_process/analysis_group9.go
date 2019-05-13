package analysis_group

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"telegraphTranslator/global"
)

// 编组9，航空器数目、机型和尾流等级
//[A]B/C []表示可选
func AnalysisGroup9(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup9Info：", message)
		defer func() {
			fmt.Println("GetGroup9Info resp: ", data)
		}()
	}
	messageArr := strings.Split(message, "/")
	if len(messageArr) != 2 {
		errMsg := fmt.Sprintf("编组9数据格式错误,请检查[Message: %s]", message)
		err = errors.New(errMsg)
		return
	}

	// 开头1~2位可能为飞机数量
	var aircraftNum int64
	var offset int
	aircraftNum, parseErr := strconv.ParseInt(messageArr[0][0:2], 0, 10)
	if parseErr != nil {
		aircraftNum = -1
	}
	if -1 == aircraftNum {
		aircraftNum, parseErr = strconv.ParseInt(messageArr[0][0:1], 0, 10)
		if parseErr != nil {
			aircraftNum = -1
		}
	}
	if -1 != aircraftNum {
		data = "飞机数目：" + strconv.FormatInt(aircraftNum, 10) + ", "
		if aircraftNum >= 10 {
			offset = 2
		} else {
			offset = 1
		}
	} else {
		offset = 0
	}
	// 获取机型
	aircraftType := messageArr[0][offset:]

	// 尾流标志
	wakeflow := ""
	switch messageArr[1] {
	case "H":
		wakeflow = "重型机"
	case "M":
		wakeflow = "中型机"
	case "L":
		wakeflow = "轻型机"
	default:
		err = errors.New(fmt.Sprintf("无效尾流标志[Message: %s]", wakeflow))
		return
	}

	data += "机型" + aircraftType + "/" + wakeflow + "\n"
	return
}
