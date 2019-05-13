package group_process

import (
	"strings"
)

//编组22修订
// [A/B]

// A:需要修改的编组代号，B修改的数据
func GetGroup22Info(message string) (data string, err error) {
	//fmt.Println("GetGroup22Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup22Info：resp", data)
	//}()

	tmpArr := strings.Split(message, "/")
	messageType := tmpArr[0]
	tmp := ""
	switch messageType {
	case "8":
		tmp, _ = GetGroup8Info(tmpArr[1])
	case "14":
		tmp, _ = GetGroup14Info(tmpArr[1]+"/"+tmpArr[2])
	case "18":
		tmp, _ = GetGroup18Info(tmpArr[1])
	default:
		data = "此工具暂未实现对此报文的解析：" + message + "\n"
		return
	}

	data = "修改编组"+tmpArr[0]+"的内容为：" + tmp +"\n"
	return
}