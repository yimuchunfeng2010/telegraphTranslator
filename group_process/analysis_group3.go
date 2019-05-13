package analysis_group

import (
	"fmt"
	"telegraphTranslator/config"
	"telegraphTranslator/global"
)

// 编组3： 电报类型
func AnalysisGroup3(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo{
		fmt.Println("GetGroup3Info：", message)
		defer func() {
			fmt.Println("GetGroup3Info resp ", data)
		}()
	}


	if len(message) < 3 {
		fmt.Printf("电报类型数据错误：请检查[Message: %s]\n",message)
		return
	}
	messageType := message[0:3]
	for _, value := range config.Group3_info.Info {
		if value.MessageType == messageType {
			data = value.Desc + "\n"
		}
	}

	messageContent := message[3:]

	tmp := ""

	messageArr := make([]string, 0)
	for _, value := range messageContent {
		ch := fmt.Sprintf("%c", value)
		if "/" == ch {
			if "" != tmp {
				messageArr = append(messageArr, tmp)
			}

			tmp = ""
		} else if ch >= "A" && ch <= "Z" {
			if 0 != len(tmp) && tmp[0:1] >= "A" && tmp[0:1] <= "Z" {
				tmp += ch
			} else {
				if "" != tmp {
					messageArr = append(messageArr, tmp)
				}
				tmp = ch
			}

		} else if ch >= "0" && ch <= "9" {
			if 0 != len(tmp) && tmp[0:1] >= "0" && tmp[0:1] <= "9" {
				tmp += ch
			} else {
				if "" != tmp {
					messageArr = append(messageArr, tmp)
				}
				tmp = ch
			}
		}
	}
	messageArr = append(messageArr, tmp)

	// 解析报文
	if "LAM" == messageType {
		tmp, _ = ProLAMGroup3Message(messageArr)
	}

	if "" != tmp {
		data += "   -" + tmp + "\n"
	}

	return
}

func ProLAMGroup3Message(message []string) (data string, err error) {

	data = "空中交通服务单位" + message[1] + "收到了单位" + message[0] + "所发的" + message[2] + "电报" + ","
	data += "此电报为对单位" + message[3] + "发送给单位" + message[4] + "的" + message[5] + "号电报的确认"
	return
}
