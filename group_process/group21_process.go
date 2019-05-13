package group_process

import (
	"strings"
)

//编组21 无线电失效情报
func GetGroup21Info(message string) (data string, err error) {
	//fmt.Println("GetGroup21Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup21Info：resp", data)
	//}()

	messageArr := strings.Split(message, " ")
	// 最后双向联系的时间
	data += "最后双向联系时间" + messageArr[0][0:2]+"时"+messageArr[0][2:]+"分, "
	// 最后联系的频率
	data += "最后联系频率" + messageArr[1]+"MHz, "
	// 最后报告位置
	data += "最后报告位置" + messageArr[2]+", "
	// 最后报告时间
	data += "最后报告时间"+ messageArr[3][0:2]+"时"+messageArr[3][2:]+"分, "
	// 最后通信能力
	otherInfo := ""
	for _, value := range messageArr[4:]{
		otherInfo += " " + strings.ToLower(value)
	}

	otherInfo = strings.ReplaceAll(otherInfo,"mhz","MHz")
	data += "其他信息:" + otherInfo

	data += "\n"
	return
}
