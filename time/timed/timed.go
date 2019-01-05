package main

import (
	"fmt"
	"os"
)
func _main()error{
	// 加载参数解析命令
	// 这个方法也会初始化日志并对应地配置
	loadedCfg,_,err := loadConfig()
	if err != nil {
		return fmt.Errorf("Could not load configuration file: %v", err)
	}
	defer func() {
		if logRotator != nil {
			logRotator.Close()
		}
	}()
	var proxy bool
	mode := "Store"
	if loadedCfg.StoreHost != ""{
		proxy = true
		mode = "Proxy"
	}
	log.Infof("Version : %v", version())
	log.Infof("Mode    : %v", mode)
	log.Infof("proxy    : %v", proxy)
	log.Infof("Network : %v", activeNetParams.Params.Name)
	log.Infof("Home dir: %v", loadedCfg.HomeDir)
	err= os.MkdirAll(loadedCfg.DataDir,0700)
	if err != nil {
		return err
	}
	// Generate the TLS cert and key file if both don't already
	if !fileExists(loadedCfg.HTTPSKey) && !fileExists(loadedCfg.HTTPSCert){
		llog.
	}
	return  nil
}
func main(){
	err := _main()
	if err != nil{
		fmt.Fprint(os.Stderr,"%v\n",err)
		os.Exit(1)
	}

}