package analysis_group

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"telegraphTranslator/global"
)

// 编组15 航路
func AnalysisGroup15(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup15Info：", message)
		defer func() {
			fmt.Println("GetGroup15Info resp: ", data)
		}()
	}
	messageArr := strings.Split(message, "/")

	data, _ = GetCruiseInfo(messageArr[0])
	// todo 暂未处理航路
	// messageArr[1]

	data += "\n"
	return
}

// 解析巡航高度，速度和航路点
func GetCruiseInfo(message string) (data string, err error) {
	messageArr := strings.Split(message, " ")

	// 第一个元素是巡航速度与高度
	// 拆分为航速和高度
	speedHeight := messageArr[0]
	offset := 1
	for {
		if "A" <= speedHeight[offset:offset+1] && speedHeight[offset:offset+1] <= "Z" {
			break
		}
		offset++
	}

	speed, _ := GetSpeedInfo(speedHeight[:offset])
	height, _ := GetHeightInfo(speedHeight[offset:])

	// 第二个元素及以后皆为航路点
	point := "经过航路点"
	for _, value := range messageArr[1:] {
		point += value + " "
	}

	data = speed + ", " + height + ", " + point
	return
}

func GetSpeedInfo(message string) (data string, err error) {
	data = "巡航速度"
	unitType := message[0:1]
	switch unitType {
	case "K":
		speed, _ := strconv.ParseFloat(message[1:], 32)
		speedStr := strconv.FormatFloat(speed, 'f', -1, 32)
		data += speedStr + "千米每小时"
	case "N":
		speed, _ := strconv.ParseFloat(message[1:], 32)
		speedStr := strconv.FormatFloat(speed, 'f', -1, 32)
		data += speedStr + "海里每小时"
	case "M":
		speed, _ := strconv.ParseFloat(message[1:], 32)
		speed /= 100
		speedStr := strconv.FormatFloat(speed, 'f', -1, 32)
		data += speedStr + "马赫每小时"
	default:
		err = errors.New("非法速度类型")
		return

	}
	return
}

func GetHeightInfo(message string) (data string, err error) {
	data = "巡航高度"
	if "VFR" == message {
		data = "巡航高度不受管制的目视飞行规则"
		return
	}
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
