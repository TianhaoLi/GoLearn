package main

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writeList []LogWriter
}

//注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writeList = append(l.writeList,writer)
}

//将data写入日志
func (l *Logger) Log(data interface{}) {
	for _,writer := range l.writeList{
		writer.Write(data)
	}
}

//创建日志的实例
func NewLogger() *Logger {
	return &Logger{}
}
