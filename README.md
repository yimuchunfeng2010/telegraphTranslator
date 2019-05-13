# 描述
    一个将MH/T 4007-2012标准电报内容解析成中文描述信息的工具
## 编译步骤
    go build telegraph_analysis.go
## 运行环境
    不依赖系统库环境，不同操作系统编译成相应的可执行文件运行即可
## 示例
    以windows环境为例，经过编译后生成telegraph_analysis.exe可执行文件
    1、启动
         .\telegraph_analysis.exe
    2、输入
        (1)、提前在input.txt文件文件中写入待解析的电报内容，或者运行过程中向input.txt文件写入电报信息
        (2)、运行过程中在终端中输入待解析的电报内容

## 效果图
![效果图](pic/pic.png)