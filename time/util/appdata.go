package util

import (
	"runtime"
	"strings"
	"unicode"
)

// AppDataDir 返回操作系统指定的应用数据存储目录
// appName:应用程序申请的数据存储目录，POSIX风格的系统会在目录前添加一个,这是标准的做法
// appName如果为空或者只是一个"，"，表示请求当前目录，
func AppDataDir(appNmae string, roaming bool) string {
	return appDataDir(runtime.GOOS,appNmae,roaming)
}

func appDataDir(goos,appName string,roaming bool) string  {
	if appName == "" || appName == "." {
		return "."
	}
	// appName不能以"."开头，我们要加以判断
	if strings.HasPrefix(appName,"."){
		appName = appName[1:]
	}
	// unicode
	appNameUpper := string(unicode.ToUpper(rune(appName[0]))) + appName[1:]
	appNameLower := string(unicode.ToLower(rune(appName[0]))) + appName[1:]

}