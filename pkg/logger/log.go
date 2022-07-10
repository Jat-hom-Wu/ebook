package logger

import (
	// "log"
	"fmt"
	"log"
	"os"
	"runtime"
	"path/filepath"
)

type Level int

var (
	logging *log.Logger
	defaultPrefix string = ""
	DefaultCallerDepth = 2

	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func OpenLogFile(fileName, filePath string)(*os.File, error){
	err := os.MkdirAll(filePath, 0777)
	if err != nil{
		return nil, fmt.Errorf("Fail to open directory: %v\n", err)
	}
	f,err := os.OpenFile(filePath + "/" + fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil{
		return nil, fmt.Errorf("Fail to openfile: %v\n", err)
	}
	return f, nil
}

func SetUpLog(fileName, filePath string){
	f,err := OpenLogFile(fileName, filePath)
	if err != nil{
		log.Fatalf("logger set up failed:%v\n", err)
	}
	logging = log.New(f, defaultPrefix, log.LstdFlags)
}

func Info(v ...interface{}){
	setPrefix(INFO)
	logging.Println(v...)
}

func Debug(v ...interface{}){
	setPrefix(DEBUG)
	logging.Println(v...)
}

func Warning(v ...interface{}){
	setPrefix(WARNING)
	logging.Println(v...)
}

func Error(v ...interface{}){
	setPrefix(ERROR)
	logging.Println(v...)
}

func Fatal(v ...interface{}){
	setPrefix(FATAL)
	logging.Println(v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logging.SetPrefix(logPrefix)
}