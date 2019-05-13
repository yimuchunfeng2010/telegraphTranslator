package group_process

// 编组8，飞行规则与种类 [A|	B]
func GetGroup8Info(message string) (data string, err error) {
	//fmt.Println("GetGroup8Info：", message)
	//defer func() {
	//	fmt.Println("GetGroup8Info resp: ", data)
	//}()

	flightRule, _ := GetFlightRule(message[0:1])
	flightType, _ := GetFlightType(message[1:2])
	data =flightRule +"," + flightType +"\n"
	return
}

func GetFlightRule(message string)(data string , err error){
	var flightRule = map[string]string{
		"I": "仪表飞行",
		"V": "目视飞行",
		"Y": "先仪表飞行", // todo
		"Z": "先目视飞行", // todo
	}
	data = flightRule[message]
	return
}

func GetFlightType(message string)(data string , err error){
	var flightType = map[string]string{
		"G": "通用航空飞行",
		"M": "军用飞行",
		"N": "非定期航空运输飞行",
		"S": "定期航空运输飞行",
		"X": "其他飞行种类", // todo
	}
	data = flightType[message]
	return
}