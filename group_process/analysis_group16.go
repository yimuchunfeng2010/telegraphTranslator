package analysis_group

import (
	"fmt"
	"strings"
	"telegraphTranslator/global"
	"telegraphTranslator/tool"
)

// 编组16 目的地机场，预计飞行时间及备降机场
// [AB  C]

// A:目的机场，B:预计飞行时间，C：目的地备降机场(可选todo待确认)
func AnalysisGroup16(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup16Info：", message)
		defer func() {
			fmt.Println("GetGroup16Info resp ：", data)
		}()
	}
	meassageArr := strings.Split(message, " ")

	// 第一部分为目的机场+飞行时间(可选)
	airportName := ""
	if len(message) >= 4 {
		arrportCode := meassageArr[0][:4]

		// 获取机场名称
		arrportInfo, newErr := tool.GetAirportInfo(arrportCode)
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
		arrportInfo, newErr := tool.GetAirportInfo(meassageArr[1])
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
