package analysis_group

import (
	"errors"
	"fmt"
	"strings"
	"telegraphTranslator/global"
)

// 编组5，紧急情况说明
func AnalysisGroup5(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup5Info：", message)
		defer func() {
			fmt.Println("GetGroup5Info resp ", data)
		}()
	}
	messageArr := strings.Split(message, "/")

	if len(messageArr) < 3 {
		errMsg := fmt.Sprintf("编组5紧急情况说明数据错误，请检查[Mesage: %s]\n", message)
		err = errors.New(errMsg)
		return
	}
	riskLevel, _ := GetRiskLevel(messageArr[0])

	reporterInfo, _ := GetReporterInfo(messageArr[1])

	otherInfo := "其他信息: " + strings.ToLower(messageArr[2])

	data = riskLevel + ", " + reporterInfo + ", " + otherInfo + "\n"
	return
}

func GetRiskLevel(message string) (data string, err error) {
	tmp := ""
	switch message {
	case "INCERFA":
		tmp = "不明阶段"
	case "ALERFA":
		tmp = "告警阶段"
	case "DETRESFA":
		tmp = "遇险阶段"
	}

	data = "危险等级：" + tmp
	return
}

func GetReporterInfo(message string) (data string, err error) {
	// 前4位为机场编号

	airportCode := message[0:4]
	airportName, _ := GetAirportName(airportCode)

	// todo 未处理后四位编码 地区+单位内部编号

	data = "报告单位：" + airportName
	return
}
