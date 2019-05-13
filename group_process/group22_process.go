package group_process

import (
	"fmt"
	"strings"
)

//编组22 修订报文
// [A/B]

// A:需要修改的编组代号，B修改的数据
func GetGroup22Info(message string) (data string, err error) {
	//fmt.Println("GetGroup22Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup22Info：resp", data)
	//}()

	messageArr := strings.Split(message, "/")
	if len(messageArr) < 4 {
		fmt.Printf("编组22修订数据错误，请检查[Message: %s]\n",message)
		return
	}

	messageType := messageArr[0]
	tmp := ""
	switch messageType {
	case "8":
		tmp, _ = GetGroup8Info(messageArr[1])
	case "14":
		tmp, _ = GetGroup14Info(messageArr[1]+"/"+messageArr[2])
	case "18":
		tmp, _ = GetGroup18Info(messageArr[1])
	default:
		data = "此工具暂未实现对此报文的解析：" + message + "\n"
		return
	}

	data = "修改编组"+messageArr[0]+"的内容为：" + tmp +"\n"
	return
}