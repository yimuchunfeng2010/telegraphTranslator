package analysis_group

import (
	"errors"
	"fmt"
	"strings"
	"telegraphTranslator/global"
)

//编组22 修订报文
// [A/B]

// A:需要修改的编组代号，B修改的数据
func AnalysisGroup22(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup22Info：", message)
		defer func() {
			fmt.Println("GetGroup22Info：resp", data)
		}()
	}

	messageArr := strings.Split(message, "/")
	if len(messageArr) < 2 {
		errMsg := fmt.Sprintf("编组22修订数据错误，请检查[Message: %s]\n", message)
		err = errors.New(errMsg)
		return
	}

	messageType := messageArr[0]
	tmp := ""
	switch messageType {
	case "8":
		tmp, _ = AnalysisGroup8(messageArr[1])
	case "14":
		tmp, _ = AnalysisGroup14(messageArr[1] + "/" + messageArr[2])
	case "18":
		tmp, _ = AnalysisGroup18(messageArr[1])
	default:
		data = "此工具暂未实现对此报文的解析：" + message + "\n"
		return
	}

	data = "修改编组" + messageArr[0] + "的内容为：" + tmp + "\n"
	return
}
