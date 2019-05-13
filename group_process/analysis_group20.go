package analysis_group

import (
	"fmt"
	"strings"
	"telegraphTranslator/global"
)

//编组20 搜寻和救援告警情报
func AnalysisGroup20(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup20Info：", message)
		defer func() {
			fmt.Println("GetGroup20Info：resp", data)
		}()
	}
	messageArr := strings.Split(message, " ")
	//营运人代号
	data += "营运人: " + messageArr[0] + ", "
	//最后联系单位 todo 未解析最后两位
	airportName, _ := GetAirportName(messageArr[1][:4])
	data += "最后联系单位：" + airportName + ", "
	//最后双向联系时间
	data += "最后联系时间: " + messageArr[2][:2] + "时" + messageArr[2][2:] + "分, "
	//最后联系频率
	data += "最后联系频率：" + messageArr[3] + "MHz, "
	//最后报告位置/时间
	data += "最后报告位置：" + messageArr[4][:len(messageArr[4])-4] + ", " +
		"时间" + messageArr[4][len(messageArr[4])-4:len(messageArr[4])-2] + "时" +
		messageArr[4][len(messageArr[4])-2:] + "分，"
	// 其他信息
	tmp := ""
	for _, value := range messageArr[5:] {
		if "NIL" == value {
			continue
		}
		tmp += " " + strings.ToLower(value)
	}
	data += "其他信息:" + tmp
	data += "\n"
	return
}
