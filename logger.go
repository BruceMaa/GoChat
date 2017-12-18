package main

import (
	"io"
	"log"
	"os"
	"sync"
	"fmt"
)

type GoChatLogger struct {
	DebugLogger *log.Logger
	TraceLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

var (
	DebugHandle io.Writer = os.Stdout
	TraceHandle io.Writer = os.Stdout
	InfoHandle  io.Writer = os.Stdout
	WarnHandle  io.Writer = os.Stdout
	ErrorHandle io.Writer = os.Stderr
)

var debugOnce sync.Once
var traceOnce sync.Once
var infoOnce sync.Once
var warnOnce sync.Once
var errorOnce sync.Once

var Logger = &GoChatLogger{}

func (gLog *GoChatLogger) Debugf(format string, v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debugf")
		gLog.DebugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.DebugLogger.Printf(format, v...)
}

func (gLog *GoChatLogger) Debug(v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debug")
		gLog.DebugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.DebugLogger.Println(v...)
}

func (gLog *GoChatLogger) Tracef(format string, v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Tracef")
		gLog.TraceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.TraceLogger.Printf(format, v...)
}

func (gLog *GoChatLogger) Trace(v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Trace")
		gLog.TraceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.TraceLogger.Println(v...)
}

func (gLog *GoChatLogger) Infof(format string, v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Infof")
		gLog.InfoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.InfoLogger.Printf(format, v...)
}

func (gLog *GoChatLogger) Info(v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Info")
		gLog.InfoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.InfoLogger.Println(v...)
}

func (gLog *GoChatLogger) Warningf(format string, v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init WarningF")
		gLog.WarnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.WarnLogger.Printf(format, v...)
}

func (gLog *GoChatLogger) Warning(v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init Warning")
		gLog.WarnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.WarnLogger.Println(v...)
}

func (gLog *GoChatLogger) Errorf(format string, v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Errorf")
		gLog.ErrorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.ErrorLogger.Printf(format, v...)
}

func (gLog *GoChatLogger) Error(v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Error")
		gLog.ErrorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	gLog.ErrorLogger.Println(v...)
}
