package main

import (
	"fmt"
	"group_process"
	"tool"
)

func main() {
	// 从文件中读取电报内容
	message := make(chan string, 100)
	go func (){
		err := tool.GetTeleMessage(message)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()


	// 解析电报

	for{
		line := <- message
		if 0 == len(line){
			continue
		}
		resp, err := group_process.ProcessMessageMain(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp)
	}


}
