package group_process

import (
	"strings"
)

//编组18 其他补充情报
// [AB  C]

// A:落地机场，B:落地时间，C：落地机场(A为ZZZZ时)
func GetGroup18Info(message string) (data string, err error) {
	//fmt.Println("GetGroup18Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup18Info：resp", data)
	//}()

	if "0" == message {
		data = "本航班无其他补充信息\n"
		return
	}

	// 重新对报文进行分组
	tmpArr := strings.Split(message, " ")
	tmp := ""
	newMessageArr := make([]string, 0)
	for _, value := range tmpArr {
		if " " == value {
			continue
		}
		if strings.Contains(value, "/") {
			newMessageArr = append(newMessageArr, tmp)
			tmp = value
		} else {
			tmp += " " + value
		}
	}

	// 处理最后一个分组
	newMessageArr = append(newMessageArr, tmp)
	for _, curMessage := range newMessageArr {
		messageArr := strings.Split(curMessage, "/")
		// 处理不同情报报文
		messageType := messageArr[0]
		if len(messageArr) < 2 {
			tmp = "类型不支持： " + messageArr[0] + ", "
			continue
		}
		tmp = ""
		switch messageType {
		case "DOF":
			tmp, _ = ProDOFMessage(messageArr[1])
		case "STS":
			tmp, _ = ProSTSMessage(messageArr[1])
		case "PBN":
			tmp, _ = ProPBNMessage(messageArr[1])
		case "NAV":
			tmp, _ = ProNAVMessage(messageArr[1])
		case "REG":
			tmp, _ = ProREGMessage(messageArr[1])
		case "EET":
			tmp, _ = ProEETMessage(messageArr[1])
		case "SEL":
			tmp, _ = ProSELMessage(messageArr[1])
		case "RMK":
			tmp, _ = ProRMKMessage(messageArr[1])
		case "RIF":
			tmp, _ = ProRIFMessage(messageArr[1])
		case "PER":
			tmp, _ = ProPERMessage(messageArr[1])
		case "OPR":
			tmp, _ = ProOPRMessage(messageArr[1])
		case "0":
			tmp = "无其他补充信息"
		default:
			tmp = "此类型暂未支持解析： " + messageType + "/" + messageArr[1]
		}

		data += tmp + ", "
	}
	data = data[:len(data)-2] + "\n"
	return
}

// 起飞日期
func ProDOFMessage(message string) (data string, err error) {
	// todo 暂时未做参数合法性校验
	data = "起飞日期：" + "20" + message[0:2] + "年" + message[2:4] + "月" + message[4:] + "日"
	return
}

func ProSTSMessage(message string) (data string, err error) {

	infoMap := map[string]string{
		"ALTRV":   "按照预留高度运行的飞行",
		"ATFMX":   "有关交通服务当局批准豁免交通流量管理措施的飞行",
		"FFR":     "灭火",
		"FLTCK":   "校验导航设施的飞行检测",
		"HAZMAT":  "运载有害材料的飞行",
		"HEAD":    "国家领导人性质的飞行",
		"HOSP":    "医疗当局公布的医疗飞行",
		"HUM":     "执行人道主义任务的飞行",
		"MARSA":   "军方负责管理的军用航空器最低安全飞行高度间隔飞行",
		"MEDEVAC": "与生命攸关的医疗紧急疏散",
		"NONRVSM": "不具备缩小垂直间隔能力的飞行准备在缩小垂直间隔空域运行",
		"SAR":     "从事搜寻与援救任务的飞行",
		"STATE":   "从事军队、海关或警察服务的飞行",
	}

	data = infoMap[message]
	return
}

func ProPBNMessage(message string) (data string, err error) {
	// todo 待解析
	data = "PBN能力为" + message
	return
}

func ProNAVMessage(message string) (data string, err error) {

	switch message {
	case "ABAS":
		data = "全球导航增强系统ABAS"
	default:
		data = "NAV能力类型不明确"

	}
	return
}

func ProREGMessage(message string) (data string, err error) {

	data = "航空器登记标志" + message
	return
}

func ProEETMessage(message string) (data string, err error) {

	messageArr := strings.Split(message, " ")
	data = "起飞预计"
	for _, value := range messageArr {
		location := value[:len(value)-4]
		timeHour := value[len(value)-4 : len(value)-2]
		timeMinute := value[len(value)-2:]
		data += "到达" + location + "用时" + timeHour + "小时" + timeMinute + "分, "
	}

	data = data[:len(data)-2]
	return
}

func ProSELMessage(message string) (data string, err error) {
	data = "航空器选呼编码" + message
	return
}

func ProRIFMessage(message string) (data string, err error) {
	data = "至修改后的目的地机场的航路详情" + message
	return
}

func ProRMKMessage(message string) (data string, err error) {
	messageArr := strings.Split(message, " ")
	switch message {
	case "ACAS II":
		data = "机上载有ACAS II防相撞设备"
	case "TCAS":
		data = "机上载有TCAS设备"
	case "POLAR":
		data = "极地飞行"
	case "APVD NONRVSM":
		data = "不具备RVSM能力的航空器获批在RVSM空域飞行"
	case "CHARTER":
		data = "包机"
	}

	// 处理返航和备降的情形
	flag := false
	switch messageArr[0] {
	case "RETURN":
		data = "返航信息: "
		flag = true
	case "ALTERNATE":
		data = "备降信息: "
		flag = true

	}
	if true == flag {
		for index, value := range messageArr {
			if 0 != index {
				data += " " + strings.ToLower(value)
			}
		}
	}
	if 0 == len(data) {
		data = "编组18其他补充信息: " + strings.ToLower(message)
		return
	}
	return
}

func ProPERMessage(message string) (data string, err error) {
	data = "航空器进近类别" + message
	return
}

func ProOPRMessage(message string) (data string, err error) {
	data = "航空器运行机构" + message
	return
}
