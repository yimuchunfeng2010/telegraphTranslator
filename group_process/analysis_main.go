package analysis_group

import (
	"errors"
	"fmt"
	"strings"
	"telegraphTranslator/config"
)

// 解析电报主程序
func ProcessMessageMain(message string) (data string, err error) {
	fmt.Println("电报内容: ", message)
	// 按照"-"或者换行符进行切分
	newMessage := strings.ReplaceAll(message, "\n", "")
	messageArr := strings.Split(newMessage, "-")

	if 0 == len(messageArr) {
		fmt.Println("电报输入错误，请检查")
		return
	}

	// 编组3 解析电报类型
	group3Info, err := AnalysisGroup3(messageArr[0])
	if err != nil {
		return
	}
	// 对逻辑确认报特殊处理
	if strings.HasPrefix(messageArr[0], "LAM") {
		data = group3Info
		return
	}

	var messageRule config.MessageRule = config.MessageRule{
		"",
		[]int{0},
	}
	for _, rule := range config.Config.MessageRuleList {
		if rule.MessageType == messageArr[0] {
			messageRule = rule
		}
	}

	if messageRule.MessageType == "" {
		errMsg := "未找到正确的电报类型"
		err = errors.New(errMsg)
		return
	}

	data += group3Info
	index := 1
	blank := "    -"
	for _, value := range messageRule.Rule {
		processRsp := ""
		if index > len(messageArr) {
			break
		}
		// 按照编制类型进行解析
		switch value {
		case 3:
			// 编组3已处理
		case 5:
			processRsp, err = AnalysisGroup5(messageArr[index])
		case 7:
			processRsp, err = AnalysisGroup7(messageArr[index])
		case 8:
			processRsp, err = AnalysisGroup8(messageArr[index])
		case 9:
			processRsp, err = AnalysisGroup9(messageArr[index])
		case 10:
			processRsp, err = AnalysisGroup10(messageArr[index])
		case 13:
			processRsp, err = AnalysisGroup13(messageArr[index])
		case 14:
			processRsp, err = AnalysisGroup14(messageArr[index])
		case 15:
			processRsp, err = AnalysisGroup15(messageArr[index])
		case 16:
			processRsp, err = AnalysisGroup16(messageArr[index])
		case 17:
			processRsp, err = AnalysisGroup17(messageArr[index])
		case 18:
			processRsp, err = AnalysisGroup18(messageArr[index])
		case 19:
			processRsp, err = AnalysisGroup19(messageArr[index])
		case 20:
			processRsp, err = AnalysisGroup20(messageArr[index])
		case 21:
			processRsp, err = AnalysisGroup21(messageArr[index])
		case 22:
			processRsp, err = AnalysisGroup22(messageArr[index])
		default:
			errMsg := fmt.Sprintf("当前未对编组%d的内容进行解析\n", value)
			err = errors.New(errMsg)
			return
		}

		if err != nil {
			return
		}
		index++
		data += blank + processRsp
	}

	return
}
