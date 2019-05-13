package group_process

import (
	"errors"
	"fmt"
	"strings"
	"telegraphTranslator/config"
)

// 解析电报主程序
func ProcessMessageMain(message string) (data string, err error) {
	fmt.Println("Message: ", message)
	// 按照"-"或者换行符进行切分
	newMessage := strings.ReplaceAll(message, "\n", "")
	messageArr := strings.Split(newMessage, "-")

	// 解析电报类型
	group3Info, err := GetGroup3Info(messageArr[0])
	if err != nil {
		return
	}
	// 对逻辑确认报特殊处理
	if strings.HasPrefix(messageArr[0],"LAM"){
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
		switch value {
		case 3:
			// 编组3已处理
		case 5:
			processRsp, err = GetGroup5Info(messageArr[index])
		case 7:
			processRsp, err = GetGroup7Info(messageArr[index])
		case 8:
			processRsp, err = GetGroup8Info(messageArr[index])
		case 9:
			processRsp, err = GetGroup9Info(messageArr[index])
		case 10:
			processRsp, err = GetGroup10Info(messageArr[index])
		case 13:
			processRsp, err = GetGroup13Info(messageArr[index])
		case 14:
			processRsp, err = GetGroup14Info(messageArr[index])
		case 15:
			processRsp, err = GetGroup15Info(messageArr[index])
		case 16:
			processRsp, err = GetGroup16Info(messageArr[index])
		case 17:
			processRsp, err = GetGroup17Info(messageArr[index])
		case 18:
			processRsp, err = GetGroup18Info(messageArr[index])
		case 19:
			processRsp, err = GetGroup19Info(messageArr[index])
		case 20:
			processRsp, err = GetGroup20Info(messageArr[index])
		case 21:
			processRsp, err = GetGroup21Info(messageArr[index])
		case 22:
			processRsp, err = GetGroup22Info(messageArr[index])
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
