package group_process

import (
	"fmt"
	"strings"
	"telegraphTranslator/tool"
)

// 编组16数据格式
// [AB  C]

// A:目的机场，B:预计飞行时间，C：目的地备降机场(可选todo待确认)
func GetGroup16Info(message string) (data string, err error) {
	//fmt.Println("GetGroup16Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup16Info resp ：", data)
	//}()
	meassageArr := strings.Split(message, " ")

	// 第一部分为目的机场+飞行时间(可选)
	airportName := ""
	if len(message) >= 4 {
		arrportCode := meassageArr[0][:4]

		// 获取机场名称
		arrportInfo, newErr := tool.FindAirportInfo(arrportCode)
		if newErr != nil {
			err = newErr
			return
		}

		airportName = arrportInfo.AirPortName
	}

	filghtTime := ""
	if len(message) >= 8 {

		// 获取飞行时间
		flag := false
		for _, value := range meassageArr[0][4:] {
			if "0" == fmt.Sprintf("%c", value) && false == flag {
				continue
			}
			filghtTime += fmt.Sprintf("%c", value)
			flag = true
		}

	}

	bakAirportName := ""
	if len(meassageArr) >= 2 {
		// 获取备降机场
		arrportInfo, newErr := tool.FindAirportInfo(meassageArr[1])
		if newErr != nil {
			err = newErr
			return
		}
		bakAirportName = arrportInfo.AirPortName
	}

	if "" != airportName {
		data += "目的地机场: " + airportName + ", "
	}

	if "" != filghtTime {
		data += "预计飞行时长: " + filghtTime + "分钟, "
	}

	if "" != bakAirportName {
		data += "备降机场: " + bakAirportName + ", "
	}

	data = data[:len(data)-2] + "\n"

	return
}
