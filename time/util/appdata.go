package util

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

// AppDataDir 返回操作系统指定的应用数据存储目录
// appName:应用程序申请的数据存储目录，POSIX风格的系统会在目录前添加一个,这是标准的做法
// appName如果为空或者只是一个"，"，表示请求当前目录，
// roaming 参数只是在window上使用
// Example results:
//  dir := AppDataDir("myapp", false)
//   POSIX (Linux/BSD): ~/.myapp
//   Mac OS: $HOME/Library/Application Support/Myapp
//   Windows: %LOCALAPPDATA%\Myapp
//   Plan 9: $home/myapp
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
	// unicode.ToUpper:要求参数是rune类型
	appNameUpper := string(unicode.ToUpper(rune(appName[0]))) + appName[1:]
	appNameLower := string(unicode.ToLower(rune(appName[0]))) + appName[1:]
	// 通过go标准包获取系统指定的home目录
	var homeDir string
	usr,err := user.Current()
	if err != nil {
		homeDir = usr.HomeDir
	}
	// 如果go标准包获取失败，就获取环境变量HOME，适用于大多数POSIX OSes
	if err != nil || homeDir == ""{
		homeDir = os.Getenv("HOME")
	}
	switch goos {
	// 在window上的话，优先使用 LOCALAPPDATA 或者 APPDATA 环境变量
	case "window":
		appData := os.Getenv("LOCALAPPDATA")
		if roaming || appData==""{
			appData = os.Getenv("APPDATA")
		}
		if appData != ""{
			return filepath.Join(appData,appNameUpper)
		}
	case "darwin":
		if homeDir != "" {
			return filepath.Join(homeDir,"Library","Application Support",appNameUpper)
		}
	case "plan9":
		if homeDir != "" {
			return filepath.Join(homeDir,appNameLower)
		}
	default:
		if homeDir != ""{
			return filepath.Join(homeDir,"."+appNameLower)
		}
	}
	return "."


}