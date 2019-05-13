package analysis_group

import (
	"errors"
	"fmt"
	"telegraphTranslator/global"
	"telegraphTranslator/tool"
)

// 编组13，获取起飞机场及时间
func AnalysisGroup13(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup13Info：", message)
		defer func() {
			fmt.Println("GetGroup13Info resp: ", data)
		}()
	}
	// 起飞机场
	if len(message) < 4 {
		errMsg := fmt.Sprintf("编组13起飞机场信息数据错误，请检查[Message：%s]\n", message)
		err = errors.New(errMsg)
		return
	}
	airportCode := message[0:4]

	airportName, err := GetAirportName(airportCode)
	if err != nil {
		return
	}

	// 起飞时间(可选)
	depTime := ""
	if len(message) >= 8 {
		// todo 待处理AFIL
		// 处理起飞时间
		depTime = "起飞时间 " + message[4:6] + ":" + message[6:]
	}

	if "" == depTime {
		data = "起飞机场: " + airportName + "\n"
	} else {
		data = "起飞机场: " + airportName + ", " + depTime + "\n"
	}

	return
}

func GetAirportName(airportCode string) (airportName string, err error) {

	resp, err := tool.GetAirportInfo(airportCode)
	if err != nil {
		return

	}

	if "" == resp.AirPortName {
		airportName = airportCode
	} else {
		airportName = resp.AirPortName
	}

	return
}
