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
        (1)、在同级目录下input.txt文件文件中写入待解析的电报内容，运行过程中也可以写入
        (2)、运行过程中在终端中输入待解析的电报内容
        注：(1)或(2)可任意选择
    3、退出
        在终端输入:exit\q\Q\quit，然后回车，程序退出

## 效果图
![效果图](file/pic.png)