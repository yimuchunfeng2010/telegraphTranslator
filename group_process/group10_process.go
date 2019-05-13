package group_process

import (
	"strings"
)

// 编组10，获取起飞机场及时间
func GetGroup10Info(message string) (data string, err error) {
	//fmt.Println("GetGroup10Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup10Info resp: ", data)
	//}()

	messageArr := strings.Split(message, "/")

	data1, err := GetEquipmentInfo(messageArr[0])
	if err != nil {
		return
	}

	data2, err := GetCapactiyInfo(messageArr[1])
	if err != nil {
		return
	}

	data2 = data2[:len(data2)-2]
	data = data1 + data2 + "\n"
	return
}

func GetEquipmentInfo(message string) (data string, err error) {

	aDataMap := map[string]string{
		"N":  "航空器未载有无线电通信、导航、进近设备或此类设备不工作",
		"S":  "航空器未载有标准的通信、导航、进近设备并可工作",
		"A":  "GBAS 着陆系统",
		"B":  "LPV(星基增强系统的垂直引导进近程序)",
		"C":  "罗兰C",
		"D":  "测距仪",
		"E1": "飞行管理计算机、航路点位置报告、航空器通信寻址与报告系统",
		"E2": "数据链飞行情报服务、航空器通信寻址与报告服务",
		"E3": "起飞前放行、航空器通信寻址与报告系统",
		"F":  "自动定向仪",
		"G":  "全球导航卫星系统",
		"H":  "高频、无线电话",
		"I":  "惯性导航",
		"J1": "管制员驾驶员数据链通信、航空电信网、甚高频数据链模式2",
		"J2": "管制员驾驶员数据链通信、FANS1/A、高频数据链",
		"J3": "管制员驾驶员数据链通信、FANS1/A、甚高频数据链模式4",
		"J4": "管制员驾驶员数据链通信、FANS1/A、甚高频数据链模式2",
		"J5": "管制员驾驶员数据链通信、FANS1/A、卫星通信(国际海事卫星组织)",
		"J6": "管制员驾驶员数据链通信、FANS1/A、卫星通信(多功能运输卫星)",
		"J7": "管制员驾驶员数据链通信、FANS1/A、卫星通信(铱星)",
		"K":  "微波着陆系统",
		"L":  "仪表着陆系统",
		"M1": "空中交通管制无线电话、卫星通信(国际海事卫星组织)",
		"M2": "空中交通管制无线电话(多功能运输卫星)",
		"M3": "空中交通管制无线电话(铱星)",
		"O":  "全向信标台",
		"R":  "获得PBN批准",
		"T":  "塔康",
		"U":  "特高频无线电话",
		"V":  "甚高频无线电话",
		"W":  "获得缩小垂直间隔批准",
		"X":  "获得最低导航性能规范批准",
		"Y":  "有8.33 kHz 频道间距能力的甚高频",
		"Z":  "携带的其他设备或能力",
	}

	for key, value := range aDataMap {
		if true == strings.Contains(message, key) {
			data += value + ", "
		}
	}
	return
}

func GetCapactiyInfo(message string) (data string, err error) {
	bDataMap := map[string]string{
		"N":  "没有应答机",
		"A":  "应答机A模式",
		"C":  "应答机A模式和应答机C模式",
		"S":  "应答机S模式，具有气压高度和航空器识别的能力",
		"P":  "应答机S模式，具有气压高度和航空器识别，但没有航空器识别的能力",
		"I":  "应答机S模式，具有航空器识别，但无气压高度发射信号的能力",
		"X":  "应答机S模式，没有航空器识别和气压高度能力",
		"E":  "应答机S模式，具有航空器识别，气压高度发射信号和超长电文(ADS-B)能力",
		"H":  "应答机S模式，具有航空器识别，气压高度发射信号和增强的监视能力",
		"L":  "应答机S模式，具有航空器识别，气压高度发射信号、超长电文(ADS-B)和增强的监视能力",
		"B1": "具有专用1090 MHz广播式自动相关监视 发送 能力的广播式自动相关监视",
		"B2": "具有专用1090 MHz广播式自动相关监视 发送 和接收 能力的广播式自动相关监视",
		"U1": "使用UAT广播式自动相关监视 发送 能力",
		"U2": "使用UAT广播式自动相关监视 发送 和 接收 能力",
		"V1": "使用VDL 模式4 广播式自动相关监视 发送 能力",
		"V2": "使用VDL 模式4 广播式自动相关监视 发送 和 接收 能力",
		"D1": "具有 FANS 1/A 能力的契约式自动相关监视",
		"G1": "具有航空电信网能力的契约式自动相关监视",
	}
	for key, value := range bDataMap {
		if true == strings.Contains(message, key) {
			data += value + ", "
		}
	}
	return
}
