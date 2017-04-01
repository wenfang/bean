package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type logLevel int

const (
	debugLevel logLevel = iota
	infoLevel
	warningLevel
	errorLevel
	fatalLevel
)

func (l logLevel) String() string {
	switch l {
	case debugLevel:
		return "DEBUG"
	case infoLevel:
		return "INFO"
	case warningLevel:
		return "WARNING"
	case errorLevel:
		return "ERROR"
	case fatalLevel:
		return "FATAL"
	}
	return "NONAME"
}

type logger struct {
	sync.Mutex
	file *os.File

	level     logLevel
	calldepth int
}

func (l *logger) Debug(v ...interface{}) {
	l.output(debugLevel, l.calldepth, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.output(infoLevel, l.calldepth, v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.output(warningLevel, l.calldepth, v...)
}

func (l *logger) Error(v ...interface{}) {
	l.output(errorLevel, l.calldepth, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.output(fatalLevel, l.calldepth, v...)
}

func (l *logger) setLevel(name string) {
	switch strings.ToUpper(name) {
	case "DEBUG":
		l.level = debugLevel
	case "INFO":
		l.level = infoLevel
	case "WARNING":
		l.level = warningLevel
	case "ERROR":
		l.level = errorLevel
	case "FATAL":
		l.level = fatalLevel
	default:
		l.level = infoLevel
	}
}

func (l *logger) setFile(filename string) error {
	newfile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	oldfile := l.file
	l.Lock()
	l.file = newfile
	l.Unlock()

	if oldfile != os.Stdout && oldfile != os.Stderr {
		oldfile.Close()
	}
	return nil
}

func (l *logger) output(level logLevel, calldepth int, v ...interface{}) {
	if level < l.level || len(v) <= 0 {
		return
	}
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	var buf []byte
	timeFormat := time.Now().Format("2006-01-02 15:04:05.000")
	buf = append(buf, fmt.Sprintf("%-7s %s [%d] [%s:%d] msg=[%v]", level, timeFormat, pid, path.Base(file), line, v[0])...)
	for i := 1; i < len(v)-1; i += 2 {
		buf = append(buf, fmt.Sprintf(" %v=[%v]", v[i], v[i+1])...)
	}
	if len(v)%2 == 0 {
		buf = append(buf, fmt.Sprintf(" %v=[]", v[len(v)-1])...)
	}
	buf = append(buf, '\n')

	l.Lock()
	l.file.Write(buf)
	l.Unlock()
}

var out = &logger{level: infoLevel, file: os.Stdout, calldepth: 3}
var err = &logger{level: errorLevel, file: os.Stderr, calldepth: 3}
var pid = os.Getpid()

// Debug 输出
func Debug(v ...interface{}) {
	out.Debug(v...)
}

// Info 输出
func Info(v ...interface{}) {
	out.Info(v...)
}

// Warning 输出
func Warning(v ...interface{}) {
	out.Warning(v...)
}

// Error 输出
func Error(v ...interface{}) {
	out.Error(v...)
	err.Error(v...)
}

// Fatal 输出
func Fatal(v ...interface{}) {
	out.Fatal(v...)
	err.Error(v...)
}

// SetLevel 设置日志级别
func SetLevel(name string) {
	out.setLevel(strings.ToUpper(name))
}

// SetLogFile 设置日志输出文件
func SetLogFile(filename string) error {
	return out.setFile(filename)
}

// SetErrFile 设置错误日志输出文件
func SetErrFile(filename string) error {
	return err.setFile(filename)
}

// Init 初始化日志
func Init(config Config) {
	if config.LogFile != "" {
		if err := SetLogFile(config.LogFile); err != nil {
			Error("set log file error", "filename", config.LogFile, "error", err)
		}
	}
	if config.ErrFile != "" {
		if err := SetErrFile(config.ErrFile); err != nil {
			Error("set error file error", "filename", config.ErrFile, "error", err)
		}
	}
	if config.LogLevel != "" {
		SetLevel(config.LogLevel)
	}
}
