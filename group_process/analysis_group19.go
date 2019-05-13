package analysis_group

import (
	"fmt"
	"strconv"
	"strings"
	"telegraphTranslator/global"
)

//编组19 补充情报
// [AB  C]

// A:落地机场，B:落地时间，C：落地机场(A为ZZZZ时)
func AnalysisGroup19(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup19Info：", message)
		defer func() {
			fmt.Println("GetGroup19Info：resp", data)
		}()
	}

	messageArr := strings.Split(message, " ")

	for _, value := range messageArr {
		tmp, _ := GetAditionalInfo(value)
		data += tmp + ", "
	}
	data = data[:len(data)-2]
	data += "\n"
	return
}

func GetAditionalInfo(message string) (data string, err error) {
	tmp := ""
	messageArr := strings.Split(message, "/")
	messageType := messageArr[0]
	messageContent := messageArr[1]
	switch messageType {
	case "E":
		tmp = "续航时间" + messageContent[0:2] + "小时" + messageContent[2:] + "分钟"
	case "P":
		tmp = "机上共有" + messageContent + "人"
	case "R":
		tmp, _ = GetRadioFrequency(messageContent)
	case "S":
		tmp, _ = GetSaveInfo(messageContent)
	case "J":
		tmp, _ = GetOther1Info(messageContent)
	case "D":
		tmp, _ = GetOther2Info(messageContent)
	case "A":
		tmp, _ = GetOther3Info(messageContent)
	case "N":
		tmp, _ = GetOther4Info(messageContent)
	case "C":
		tmp, _ = GetCaptainName(messageContent)
	default:
		tmp = "无法解析此信息 " + message

	}
	data = tmp
	return
}

func GetRadioFrequency(message string) (data string, err error) {
	for _, ch := range message {
		chTmp := fmt.Sprintf("%c", ch)
		switch chTmp {
		case "U":
			data += "有特高频243.0MHz频率, "
		case "V":
			data += "有甚高频121.5MHz频率, "
		case "Z":
			data += "有紧急示位信标, "
		default:
			data += "无法解析此标志" + message
		}
	}

	data = data[:len(data)-2]
	return
}

func GetSaveInfo(message string) (data string, err error) {
	for _, ch := range message {
		chTmp := fmt.Sprintf("%c", ch)
		switch chTmp {
		case "P":
			data += "有极地救生设备"
		case "D":
			data += "有沙漠救生设备"
		case "M":
			data += "有海上救生设备"
		case "J":
			data += "有丛林救生设备"
		default:
			data += "无法解析此标志" + message
		}
	}
	return
}

func GetOther1Info(message string) (data string, err error) {
	for _, ch := range message {
		chTmp := fmt.Sprintf("%c", ch)
		switch chTmp {
		case "L":
			data += "救生衣配备有灯光"
		case "F":
			data += "救生衣配备有荧光素"
		case "U":
			data += "救生衣配备特高频243.0频率电台"
		case "V":
			data += "救生衣配备甚高频121.5频率电台"
		default:
			data += "无法解析此标志: " + message
		}
	}
	return
}
func GetOther2Info(message string) (data string, err error) {
	// 相关救生设备信息
	// 2位数字表示救生艇数目
	// 3位数字所有救生艇可载总人数
	// C救生艇有棚子
	// 英文单词表示救生艇颜色

	messageArr := strings.Split(message, " ")
	for _, value := range messageArr {
		_, transErr := strconv.ParseFloat(value, 64)
		if nil != transErr {
			// 字符由字母组成
			switch value {
			case "C":
				data += "救生艇有棚子" + ","
			default:
				data += "救生艇颜色: " + value + ", "
			}
		} else {
			if len(value) == 2 {
				data += "救生艇数目: " + value + ", "
			}
			if len(value) == 3 {
				data += "救生艇可载总人数: " + value + ", "
			}
		}
	}
	data = "暂为解析此电报 " + message
	return
}

func GetOther3Info(message string) (data string, err error) {
	data = "航空器颜色: " + strings.ToLower(message)
	return
}

func GetOther4Info(message string) (data string, err error) {
	data = "其他救生设备：" + message
	return
}

func GetCaptainName(message string) (data string, err error) {
	data = "机长: " + strings.ToLower(message)
	return
}
