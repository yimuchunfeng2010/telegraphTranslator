package analysis_group

import (
	"fmt"
	"strings"
	"telegraphTranslator/global"
	"telegraphTranslator/tool"
)

// 编组17 落地机场及时间
// [AB  C]

// A:落地机场，B:落地时间，C：落地机场(A为ZZZZ时)
func AnalysisGroup17(message string) (data string, err error) {
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup17Info：", message)
		defer func(){
			fmt.Println("GetGroup17Info：resp", data)
		}()
	}
	meassageArr := strings.Split(message, " ")

	arrportName := "" // 落地机场
	arrTime := ""     // 落地时间
	// 第一部分为落地机场+飞行时间
	if len(message) >= 8 {
		arrportCode := meassageArr[0][0:4]

		if "ZZZZ" == arrportCode {
			arrportName = ""

		} else {
			// 落地名称
			arrportInfo, newErr := tool.GetAirportInfo(arrportCode)
			if newErr != nil {
				err = newErr
				return
			}

			arrportName = arrportInfo.AirPortName
		}

		// 获取落地时间
		arrTime = meassageArr[0][4:6] + ":" + meassageArr[0][6:]

	}

	if len(meassageArr) >= 2 && "" == arrportName {
		// 获取备降机场
		arrportInfo, newErr := tool.GetAirportInfo(meassageArr[1])
		if newErr != nil {
			err = newErr
			return
		}
		arrportName = arrportInfo.AirPortName
	}

	data = "落地机场: " + arrportName + ", 落地时间 " + arrTime

	data += "\n"
	return
}
