package analysis_group

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"telegraphTranslator/global"
)

// 编组14，预计飞跃边界数据
func AnalysisGroup14(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup14Info：", message)
		defer func() {
			fmt.Println("GetGroup14Info resp: ", data)
		}()
	}
	messageArr := strings.Split(message, "/")
	if len(messageArr) < 2 {
		errMsg := fmt.Sprintf("编组14预计飞跃边界数据，请检查[Message: %s]\n", message)
		err = errors.New(errMsg)
		return
	}
	flightPoint := messageArr[0]
	data = "边界点" + flightPoint + ", "
	// 解析飞越时间，许可飞行层高度(可选)，预计飞行高度(可选)，飞越条件(与预计飞行高度配套)

	flightInfo := messageArr[1]

	flightOverTime := flightInfo[0:4]
	data += "预计飞越时间" + flightOverTime[0:2] + "时" + flightOverTime[2:] + "分" + ", "

	heightArr := make([]string, 0)
	tmp := ""
	for _, ch := range flightInfo[4:] {
		chTmp := fmt.Sprintf("%c", ch)
		if chTmp >= "A" && chTmp <= "Z" {
			if tmp != "" {
				heightArr = append(heightArr, chTmp)
			}
			tmp = chTmp
		} else {
			tmp += chTmp
		}
	}

	// 添加最后一个报文
	heightArr = append(heightArr, tmp)

	data1 := ""
	data2 := ""
	data3 := ""
	arrLen := len(heightArr)
	switch arrLen {
	case 1:
		data1, _ = AnalysisHeightInfo(heightArr[0])
	case 2:
		data1, _ = AnalysisHeightInfo(heightArr[0])
		data2, _ = AnalysisHeightInfo(heightArr[1])
	case 3:
		data1, _ = AnalysisHeightInfo(heightArr[0])
		data2, _ = AnalysisHeightInfo(heightArr[1])
		data3, _ = GetFlightOverCondition(heightArr[2])
	}

	if "" != data1 {
		data += "许可" + data1 + ", "
	}
	if "" != data2 {
		data += "预计在 " + data2
	}

	if "" != data3 {
		data += data3 + ", "
	}
	data = data[:len(data)-2]
	data += "\n"
	return
}

func GetFlightOverCondition(message string) (data string, err error) {
	switch message {
	case "A":
		data = "或其以上飞越边界点"
	case "B":
		data = "或其以下飞越边界点"
	default:
		data = "不合法字段"
	}
	return
}
func AnalysisHeightInfo(message string) (data string, err error) {

	unitType := message[0:1]
	switch unitType {
	case "M":
		height, _ := strconv.ParseFloat(message[1:], 32)
		height *= 10
		heightStr := strconv.FormatFloat(height, 'f', -1, 32)
		data += "海拔高度" + heightStr + "米"
	case "S":
		height, _ := strconv.ParseFloat(message[1:], 32)
		height *= 10
		heightStr := strconv.FormatFloat(height, 'f', -1, 32)
		data += "飞行层高度" + heightStr + "米"
	case "A":
		height, _ := strconv.ParseFloat(message[1:], 32)
		height *= 100
		heightStr := strconv.FormatFloat(height, 'f', -1, 32)
		data += "海拔高度" + heightStr + "英尺"
	case "F":
		height, _ := strconv.ParseFloat(message[1:], 32)
		height *= 100
		heightStr := strconv.FormatFloat(height, 'f', -1, 32)
		data += "飞行层高度" + heightStr + "英尺"
	default:
		err = errors.New("非法高度类型")
		return

	}
	return
}
