package main

import (
	"fmt"
	"telegraphTranslator/group_process"
	"telegraphTranslator/tool"
)

func main() {
	// 从文件中读取电报
	message := make(chan string, 100)
	go func() {
		err := tool.GetTeleMessageFromFile(message)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 从用户命令行输入
	go func() {
		err := tool.GetTeleMessageFromUserInput(message)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 解析电报
	for {
		line := <-message
		if 0 == len(line) {
			continue
		}

		if "quit" == line || "q" == line || "Q" == line || "exit" == line {
			fmt.Println("程序退出")
			break
		}
		resp, err := analysis_group.ProcessMessageMain(line)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(resp)
	}

}
