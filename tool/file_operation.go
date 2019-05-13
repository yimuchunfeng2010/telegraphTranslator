package tool

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"telegraphTranslator/config"
	"time"
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

func GetAirportInfo(info string) (data AirportInfo, err error) {
	cnt := 0
	for _, airport := range config.AirportInfo.Info {
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
func GetTeleMessageFromUserInput(message chan string) (err error) {
	for {
		var userCmd string
		fmt.Scanf("%s", &userCmd)
		message <- userCmd
	}
	return
}

func GetTeleMessageFromFile(message chan string) (err error) {
	for {
		// 等待输入文件创建完成
		if true == IsExist("input.txt") {
			break
			time.Sleep(time.Second)

		}
	}

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

		line := fmt.Sprintf("%s", a)

		message <- line
	}
}

// 判断文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
