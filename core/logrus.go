package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type MyLog struct {
}

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLoggers() {
	logrus.SetLevel(logrus.DebugLevel) // 日志等级
	logrus.SetReportCaller(true)       //开启返回函数名和行号
	logrus.SetFormatter(MyLog{})       // 格式
	logrus.AddHook(&MyHook{
		logPath: "logs",
	})
}

type MyHook struct {
	file    *os.File //当前打开的日志文件
	errFile *os.File //错误的日志文件
	date    string   // 当前日志时间
	logPath string   // 日志目录
	mu      sync.Mutex
}

func (my *MyHook) Fire(entry *logrus.Entry) error {
	// 1.写入到文件
	// 2.按时间分片
	// 3.错误的日志单独存放
	my.mu.Lock()
	defer my.mu.Unlock()
	msg, _ := entry.String()
	date := entry.Time.Format("2006-01-02-15-04")
	if my.date != date {
		// 换时间,换文件对象
		my.rotateFiles(date)
		my.date = date
	}
	if entry.Level <= logrus.ErrorLevel {
		my.errFile.Write([]byte(msg))
	}

	my.file.Write([]byte(msg))
	return nil

}

func (my *MyHook) rotateFiles(timer string) error {
	if my.file != nil {
		my.file.Close()
	}
	if my.file == nil {
		// 创建目录
		logDir := fmt.Sprintf("%s/%s", my.logPath, timer)
		os.MkdirAll(logDir, 0666)
		logPath := fmt.Sprintf("%s/info.log", logDir)
		file, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		my.file = file

		errLogPath := fmt.Sprintf("%s/err.log", logDir)
		errFile, _ := os.OpenFile(errLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		my.errFile = errFile
	}
	return nil
}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
