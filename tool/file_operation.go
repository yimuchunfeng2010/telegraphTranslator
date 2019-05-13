package tool

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"telegraphTranslator/config"
)

type AirportInfo struct {
	Code             string // 机场代码
	AirPortName      string // 中文名称
	AirPortEnName    string // 英文名称
	AirPortCity      string //机场所在地区
	IntelligenceArea string // 情报区
}

type InfoType int

const (
	UnKown      InfoType = iota // value --> 0
	AirportCode                 // value --> 1

)

func GetAirportInfo(info string) (data AirportInfo, err error){
	cnt := 0
	for _ , airport := range config.AirportInfo.Info{
		if airport.Code == info || airport.AirPortCity == info || airport.IntelligenceArea == info {
			cnt++
			data = AirportInfo{
				Code:             airport.Code,
				AirPortName:      airport.AirPortName,
				AirPortEnName:    airport.AirPortEnName,
				AirPortCity:      airport.AirPortCity,
				IntelligenceArea: airport.IntelligenceArea,
			}
		}
	}
	if 0 == cnt {
		return
	}

	if 1 != cnt {
		return data, errors.New("机场信息不唯一")
	}
	return
}
// 入场机场编码/机场地区名
func FindAirportInfo(info string) (data AirportInfo, err error) {
	var cnfFileName string = "airport_info.json"
	filePtr, err := os.Open(cnfFileName)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var airportList []AirportInfo

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&airportList)
	if err != nil {
		errMsg := fmt.Sprintf("Decoder failed[Err: %s]", err.Error())
		fmt.Println(errMsg)
		return data, errors.New(errMsg)
	}

	// 解析查找机场编码
	var cnt = 0
	for _, arrport := range airportList {
		if arrport.Code == info || arrport.AirPortCity == info || arrport.IntelligenceArea == info {
			cnt++
			data = AirportInfo{
				Code:             arrport.Code,
				AirPortName:      arrport.AirPortName,
				AirPortEnName:    arrport.AirPortEnName,
				AirPortCity:      arrport.AirPortCity,
				IntelligenceArea: arrport.IntelligenceArea,
			}
		}
	}

	if 0 == cnt {
		return
	}

	if 1 != cnt {
		return data, errors.New("机场信息不唯一")
	}

	return
}

func InitAirportData() {
	f, err := os.Open("config/tmp.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读取机场名称文件
	f2, err := os.Open("config/tmp2.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	ld := bufio.NewReader(f2)

	codeNameMap := make(map[string]string, 0)
	for {
		line, err := ld.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}

		lineArr := strings.Split(line, ",")
		codeNameMap[lineArr[5][:len(lineArr[5])-2]] = lineArr[3]
		fmt.Printf("%s %s",lineArr[5], lineArr[3])
	}
	codeInfoArr := []AirportInfo{}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}

		dataArr := strings.Split(line, " ")
		IntelligenceArea := ""
		switch dataArr[0][:2] {
		case "ZB":
			IntelligenceArea = "北京情报区"
		case "ZG", "ZJ":
			IntelligenceArea = "广州情报区"
		case "ZH":
			IntelligenceArea = "武汉情报区"
		case "ZL":
			IntelligenceArea = "兰州情报区"
		case "ZP":
			IntelligenceArea = "昆明情报区"
		case "ZS":
			IntelligenceArea = "上海情报区"
		case "ZU":
			IntelligenceArea = "成都情报区"
		case "ZW":
			IntelligenceArea = "乌鲁木齐情报区"
		case "ZY":
			IntelligenceArea = "沈阳情报区"
		case "RC":
			IntelligenceArea = "中国台湾"
		case "VH":
			IntelligenceArea = "中国香港"
		case "VM":
			IntelligenceArea = "中国澳门"
		default:
			fmt.Println("Invalid Airport Code")
		}

		fmt.Printf("%s %s\n",dataArr[0], codeNameMap[dataArr[0]])
		codeInfoArr = append(codeInfoArr, AirportInfo{
			Code:             dataArr[0],
			AirPortName:      codeNameMap[dataArr[0]],
			AirPortEnName:    dataArr[1],
			AirPortCity:      dataArr[2][:len(dataArr[2])-2],
			IntelligenceArea: IntelligenceArea,
		})
	}

	// 创建文件
	filePtr, err := os.Create("config/airport_info.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	data, err := json.MarshalIndent(codeInfoArr, "", "  ")
	if err != nil {
		errMsg := fmt.Sprintf("Encoder failed", err.Error())
		fmt.Println(errMsg)
		return
	}
	filePtr.Write(data)
}

//
//func GetTeleMessage()(message []string , err error){
//	fptr, err := os.Open("input.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer fptr.Close()
//	ld := bufio.NewReader(fptr)
//
//	for {
//		line, readErr := ld.ReadString('\n') //以'\n'为结束符读入一行
//		if readErr != nil{
//			if io.EOF == readErr{
//				message = append(message, line)
//				err = nil
//			} else {
//				err = readErr
//			}
//			return
//		}
//
//
//		message = append(message, line[:len(line)-2])
//
//	}
//	return
//}


func GetTeleMessage(message chan string)(err error) {
	fi, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Open File Failed[Err:%s]", err.Error())
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			time.Sleep(time.Second)
			continue
		}

		line := fmt.Sprintf("%s",a)

		message <- line
	}
}