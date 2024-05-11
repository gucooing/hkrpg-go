package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var LevelMap = map[int][]byte{
	DEBUG: []byte("DEBUG"),
	INFO:  []byte("INFO"),
	WARN:  []byte("WARN"),
	ERROR: []byte("ERROR"),
}

var (
	LeftBracket  = []byte("[")
	RightBracket = []byte("]")
	Space        = []byte(" ")
	Colon        = []byte(":")
	FuncBracket  = []byte("()")
	LineFeed     = []byte("\n")
)

var (
	RED     = []byte{27, 91, 51, 49, 109}
	GREEN   = []byte{27, 91, 51, 50, 109}
	YELLOW  = []byte{27, 91, 51, 51, 109}
	BLUE    = []byte{27, 91, 51, 52, 109}
	MAGENTA = []byte{27, 91, 51, 53, 109}
	CYAN    = []byte{27, 91, 51, 54, 109}
	WHITE   = []byte{27, 91, 51, 55, 109}
	RESET   = []byte{27, 91, 48, 109}
)

var LOG *Logger = nil

const (
	DefaultFileMaxSize = 10485760
	LogInfoChanSize    = 1000
	MaxWriteCacheNum   = 1000
)

type Logger struct {
	AppName       string
	Level         int
	TrackLine     bool
	TrackThread   bool
	EnableFile    bool
	FileMaxSize   int32
	DisableColor  bool
	EnableJson    bool
	File          *os.File
	LogInfoChan   chan *LogInfo
	WriteBuf      []byte
	WriteCacheNum int32
	CloseChan     chan struct{}
}

type LogInfo struct {
	Time        time.Time
	Level       int
	Msg         *[]byte
	FileName    string
	FuncName    string
	Line        int
	GoroutineId string
	ThreadId    string
}

func InitLogger(appName, logLevel string) {
	LOG = new(Logger)
	LOG.AppName = appName

	LOG.Level = LOG.getLogLevel(logLevel)
	LOG.TrackLine = true
	LOG.TrackThread = true
	LOG.EnableFile = true
	LOG.FileMaxSize = 10485760
	LOG.DisableColor = false
	LOG.EnableJson = false

	if LOG.FileMaxSize == 0 {
		LOG.FileMaxSize = DefaultFileMaxSize
	}
	LOG.File = nil
	LOG.LogInfoChan = make(chan *LogInfo, LogInfoChanSize)
	LOG.WriteBuf = make([]byte, 0)
	LOG.WriteCacheNum = 0
	LOG.CloseChan = make(chan struct{})
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		os.MkdirAll("./log", 0644)
	}
	go LOG.doLog()
}

func CloseLogger() {
	LOG.CloseChan <- struct{}{}
	<-LOG.CloseChan
}

func (l *Logger) doLog() {
	var logBuf bytes.Buffer
	timeBuf := make([]byte, 0, 64)
	exit := false
	exitCountDown := 0
	for {
		select {
		case <-l.CloseChan:
			exit = true
			exitCountDown = len(l.LogInfoChan)
		case logInfo := <-l.LogInfoChan:
			if !l.DisableColor {
				logBuf.Write(CYAN)
			}
			logBuf.Write(LeftBracket)
			logBuf.Write(logInfo.Time.AppendFormat(timeBuf, "2006-01-02 15:04:05.000"))
			logBuf.Write(RightBracket)
			if !l.DisableColor {
				logBuf.Write(RESET)
			}
			logBuf.Write(Space)

			if !l.DisableColor {
				switch logInfo.Level {
				case DEBUG:
					logBuf.Write(BLUE)
				case INFO:
					logBuf.Write(GREEN)
				case WARN:
					logBuf.Write(YELLOW)
				case ERROR:
					logBuf.Write(RED)
				}
			}
			logBuf.Write(LeftBracket)
			logBuf.Write(LevelMap[logInfo.Level])
			logBuf.Write(RightBracket)
			if !l.DisableColor {
				logBuf.Write(RESET)
			}
			logBuf.Write(Space)

			if !l.DisableColor && logInfo.Level == ERROR {
				logBuf.Write(RED)
				logBuf.Write(*logInfo.Msg)
				logBuf.Write(RESET)
			} else {
				logBuf.Write(*logInfo.Msg)
			}

			if l.TrackLine {
				logBuf.Write(Space)
				if !l.DisableColor {
					logBuf.Write(MAGENTA)
				}
				logBuf.Write(LeftBracket)
				logBuf.Write([]byte(logInfo.FileName))
				logBuf.Write(Colon)
				logBuf.Write([]byte(strconv.Itoa(logInfo.Line)))
				logBuf.Write(Space)
				logBuf.Write([]byte(logInfo.FuncName))
				logBuf.Write(FuncBracket)
				if l.TrackThread {
					logBuf.Write(Space)
					logBuf.Write([]byte("goroutine"))
					logBuf.Write(Colon)
					logBuf.Write([]byte(logInfo.GoroutineId))
					logBuf.Write(Space)
					logBuf.Write([]byte("thread"))
					logBuf.Write(Colon)
					logBuf.Write([]byte(logInfo.ThreadId))
				}
				logBuf.Write(RightBracket)
				if !l.DisableColor {
					logBuf.Write(RESET)
				}
			}

			logBuf.Write(LineFeed)

			logData := logBuf.Bytes()
			l.writeLog(logData)
			putBuf(logInfo.Msg)
			logInfoPool.Put(logInfo)
			logBuf.Reset()
			timeBuf = timeBuf[0:0]
			if exit {
				exitCountDown--
			}
		}
		if exit && exitCountDown == 0 {
			LOG.CloseChan <- struct{}{}
			return
		}
	}
}

func (l *Logger) writeLog(logData []byte) {
	l.WriteBuf = append(l.WriteBuf, logData...)
	l.WriteCacheNum++
	if len(l.LogInfoChan) != 0 && l.WriteCacheNum < MaxWriteCacheNum {
		return
	}
	l.writeLogConsole(l.WriteBuf)
	if l.EnableFile {
		l.writeLogFile(l.WriteBuf)
	}
	l.WriteBuf = l.WriteBuf[0:0]
	l.WriteCacheNum = 0
}

func (l *Logger) writeLogConsole(logData []byte) {
	_, _ = os.Stdout.Write(logData)
}

func (l *Logger) writeLogFile(logData []byte) {
	if l.File == nil {
		file, err := os.OpenFile("./log/"+l.AppName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"open new log file error: %v\n"+string(RESET), err))
			return
		}
		LOG.File = file
	}
	fileStat, err := l.File.Stat()
	if err != nil {
		_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"get log file stat error: %v\n"+string(RESET), err))
		return
	}
	if fileStat.Size() >= int64(l.FileMaxSize) {
		err = l.File.Close()
		if err != nil {
			_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"close old log file error: %v\n"+string(RESET), err))
			return
		}
		timeStr := time.Now().Format("20060102150405")
		err = os.Rename(l.File.Name(), l.File.Name()+"."+timeStr+".log")
		if err != nil {
			_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"rename old log file error: %v\n"+string(RESET), err))
			return
		}
		file, err := os.OpenFile("./log/"+l.AppName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"open new log file error: %v\n"+string(RESET), err))
			return
		}
		LOG.File = file
	}
	_, err = l.File.Write(logData)
	if err != nil {
		_, _ = os.Stdout.WriteString(fmt.Sprintf(string(RED)+"write log file error: %v\n"+string(RESET), err))
		return
	}
}

var bufPool = sync.Pool{New: func() any { return new([]byte) }}

func getBuf() *[]byte {
	p := bufPool.Get().(*[]byte)
	*p = (*p)[0:0]
	return p
}

func putBuf(p *[]byte) {
	if cap(*p) > 64<<10 {
		*p = nil
	}
	bufPool.Put(p)
}

var logInfoPool = sync.Pool{New: func() any { return new(LogInfo) }}

func Debug(msg string, param ...any) {
	if LOG.Level > DEBUG {
		return
	}
	logInfo := logInfoPool.Get().(*LogInfo)
	logInfo.Time = time.Now()
	logInfo.Level = DEBUG
	buf := getBuf()
	if LOG.EnableJson {
		jsonList := make([]any, 0)
		for _, obj := range param {
			data, _ := json.Marshal(obj)
			jsonList = append(jsonList, string(data))
		}
		param = jsonList
	}
	*buf = fmt.Appendf(*buf, msg, param...)
	logInfo.Msg = buf
	if LOG.TrackLine {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = LOG.getLineFunc()
	}
	if LOG.TrackThread {
		logInfo.GoroutineId = LOG.getGoroutineId()
		logInfo.ThreadId = LOG.getThreadId()
	}
	LOG.LogInfoChan <- logInfo
}

func Info(msg string, param ...any) {
	if LOG.Level > INFO {
		return
	}
	logInfo := logInfoPool.Get().(*LogInfo)
	logInfo.Time = time.Now()
	logInfo.Level = INFO
	buf := getBuf()
	if LOG.EnableJson {
		jsonList := make([]any, 0)
		for _, obj := range param {
			data, _ := json.Marshal(obj)
			jsonList = append(jsonList, string(data))
		}
		param = jsonList
	}
	*buf = fmt.Appendf(*buf, msg, param...)
	logInfo.Msg = buf
	if LOG.TrackLine {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = LOG.getLineFunc()
	}
	if LOG.TrackThread {
		logInfo.GoroutineId = LOG.getGoroutineId()
		logInfo.ThreadId = LOG.getThreadId()
	}
	LOG.LogInfoChan <- logInfo
}

func Warn(msg string, param ...any) {
	if LOG.Level > WARN {
		return
	}
	logInfo := logInfoPool.Get().(*LogInfo)
	logInfo.Time = time.Now()
	logInfo.Level = WARN
	buf := getBuf()
	if LOG.EnableJson {
		jsonList := make([]any, 0)
		for _, obj := range param {
			data, _ := json.Marshal(obj)
			jsonList = append(jsonList, string(data))
		}
		param = jsonList
	}
	*buf = fmt.Appendf(*buf, msg, param...)
	logInfo.Msg = buf
	if LOG.TrackLine {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = LOG.getLineFunc()
	}
	if LOG.TrackThread {
		logInfo.GoroutineId = LOG.getGoroutineId()
		logInfo.ThreadId = LOG.getThreadId()
	}
	LOG.LogInfoChan <- logInfo
}

func Error(msg string, param ...any) {
	if LOG.Level > ERROR {
		return
	}
	logInfo := logInfoPool.Get().(*LogInfo)
	logInfo.Time = time.Now()
	logInfo.Level = ERROR
	buf := getBuf()
	if LOG.EnableJson {
		jsonList := make([]any, 0)
		for _, obj := range param {
			data, _ := json.Marshal(obj)
			jsonList = append(jsonList, string(data))
		}
		param = jsonList
	}
	*buf = fmt.Appendf(*buf, msg, param...)
	logInfo.Msg = buf
	if LOG.TrackLine {
		logInfo.FileName, logInfo.Line, logInfo.FuncName = LOG.getLineFunc()
	}
	if LOG.TrackThread {
		logInfo.GoroutineId = LOG.getGoroutineId()
		logInfo.ThreadId = LOG.getThreadId()
	}
	LOG.LogInfoChan <- logInfo
}

func (l *Logger) getLogLevel(level string) int {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	default:
		panic(fmt.Sprintf("unknown log level: %v", level))
	}
}

func (l *Logger) getGoroutineId() (goroutineId string) {
	buf := make([]byte, 32)
	runtime.Stack(buf, false)
	buf = bytes.TrimPrefix(buf, []byte("goroutine "))
	buf = buf[:bytes.IndexByte(buf, ' ')]
	goroutineId = string(buf)
	return goroutineId
}

func (l *Logger) getLineFunc() (fileName string, line int, funcName string) {
	var pc uintptr
	var file string
	var ok bool
	pc, file, line, ok = runtime.Caller(2)
	if !ok {
		return "???", -1, "???"
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	split := strings.Split(funcName, ".")
	if len(split) != 0 {
		funcName = split[len(split)-1]
	}
	return fileName, line, funcName
}

func Stack() string {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}

func StackAll() string {
	buf := make([]byte, 1024*16)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}
