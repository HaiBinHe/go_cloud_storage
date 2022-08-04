package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Fields map[string]interface{}
type Logger struct {
	ctx context.Context
	newLogger *log.Logger
	fields Fields
	callers []string
}
//prefix: 每个日志行的开头，如果 flag = Lmsgprefix，则出现在每个日志行的结尾
//flag: 日志属性
func NewLogger(w io.Writer, prefix string, flag int) *Logger{
	logger:= log.New(w, prefix, flag)
	return &Logger{newLogger: logger}
}
//
func (l *Logger) WithFields(fields Fields) *Logger {
	if l.fields == nil {
		l.fields = make(Fields)
	}
	for k, v := range fields{
		l.fields[k] = v
	}
	return l
}
func (l *Logger) WithContext(ctx context.Context) *Logger {
	l.ctx = ctx
	return l
}
//skip:
//0 当前函数, 1 上一层函数,2 ...
func (l *Logger) WithCaller(skip int) *Logger {
	//函数指针，函数所在文件的名称或目录名称,函数所在文件的行号
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		//获取一个标识调用栈标识符pc对应的调用栈
		f := runtime.FuncForPC(pc)
		l.callers = []string{fmt.Sprintf("%s: %d %s",file, line, f.Name())}
	}
	return l
}
//WithCallersFrames：设置当前的整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth :=1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	//函数把当前go程调用栈上的调用栈标识符填入切片pc中，返回写入到pc中的项数
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next();more;frame, more = frames.Next(){
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if ! more {
			break
		}
	}
	l.callers = callers
	return l
}
//日志内容格式化
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields) + 4)
	data["level"] = level.String()
	data["message"] = message
	data["time"] = time.Now().Local().UnixNano()
	data["callers"] = l.callers
	if len(l.fields) > 0{
		for k, v := range l.fields{
			//判断fields中的k是否存在data中,不存在则赋值
			if _, ok := data[k];!ok{
				data[k] = v
			}
		}
	}
	return data
}
//输出
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}