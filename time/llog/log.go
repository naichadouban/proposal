package llog

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
	logFlags = log.Ldate|log.Lshortfile|log.LUTC
)
var llog *log.Logger
//日志flag
//const (
//	Ldate         = 1 << iota     //日期示例： 2009/01/23
//	Ltime                         //时间示例: 01:23:23
//	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
//	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
//	Lshortfile                    //文件和行号: d.go:23.
//	LUTC                          //日期时间转为0时区的
//	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
//)
// 定义日志对象，包含四个级别
type Logger struct{
	level int
	err *log.Logger
	warn *log.Logger
	info *log.Logger
	debug *log.Logger
}
func NewLogger(logFile string,prefix string ,level int)*Logger{
	f,err := os.Open(logFile)
	if err != nil {
		log.Panicf("open file %v failed,%v",logFile,err)
	}
	return &Logger{
		level:level,
		err:log.New(bufio.NewWriter(f),prefix+"-[E]",logFlags),
		warn:log.New(bufio.NewWriter(f),prefix+"-[W]",logFlags),
		info:log.New(bufio.NewWriter(f),prefix+"-[I]",logFlags),
		debug:log.New(bufio.NewWriter(f),prefix+"-[D]",logFlags),
	}
}
// 设置日志级别
func (ll *Logger) SetLevel(level int){
	ll.level = level
}
func (ll *Logger) Error(format string,v ...interface{}){
	if ll.level < LevelError{
		return
	}
	msg := fmt.Sprintf("[err]"+format,v...)
	ll.err.Println(msg)
}
