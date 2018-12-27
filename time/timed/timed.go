package main

import (
	"fmt"
	"os"
)
func _main()error{
	// 加载参数解析命令
	// 这个方法也会初始化日志并对应地配置
	loadConfig,_,err := loadConfig()
}
func main(){
	err := _main()
	if err != nil{
		fmt.Fprint(os.Stderr,"%v\n",err)
		os.Exit(1)
	}
}