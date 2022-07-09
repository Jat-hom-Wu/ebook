package logger

import (
	"log"
	"testing"
	"time"
)

func TestOpenFile(t *testing.T){
	_, err := OpenLogFile("test","/home/lighthouse/ebook/log")
	if err != nil{
		t.Errorf("err: %v", err)
	}
}

func TestIsLogSync(t * testing.T){
	for i:=0; i < 1000; i++{
		time.Sleep(1 * time.Second)
		log.Printf("test,%v", i)
	}
}

func TestLogCtl(t *testing.T){
	SetUpLog("test2","/home/lighthouse/ebook/log")
	Info("hello")
	Error("nonono")
}