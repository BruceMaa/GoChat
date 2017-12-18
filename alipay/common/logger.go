package common

import (
	"io"
	"log"
	"os"
	"sync"
)

type AlipayLogger struct {
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

func (gLog *AlipayLogger) Debug(format string, v ...interface{}) {
	debugOnce.Do(func() {
		gLog.DebugLogger = log.New(DebugHandle, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	})
	gLog.DebugLogger.Printf(format+"\n", v...)
}

func (gLog *AlipayLogger) Trace(format string, v ...interface{}) {
	traceOnce.Do(func() {
		gLog.TraceLogger = log.New(TraceHandle, "[TRACE] ", log.LstdFlags|log.Lshortfile)
	})
	gLog.TraceLogger.Printf(format+"\n", v...)
}

func (gLog *AlipayLogger) Info(format string, v ...interface{}) {
	infoOnce.Do(func() {
		gLog.InfoLogger = log.New(InfoHandle, "[INFO] ", log.LstdFlags|log.Lshortfile)
	})
	gLog.InfoLogger.Printf(format+"\n", v...)
}

func (gLog *AlipayLogger) Warning(format string, v ...interface{}) {
	warnOnce.Do(func() {
		gLog.WarnLogger = log.New(WarnHandle, "[WARNING] ", log.LstdFlags|log.Lshortfile)
	})
	gLog.WarnLogger.Printf(format+"\n", v...)
}

func (gLog *AlipayLogger) Error(format string, v ...interface{}) {
	errorOnce.Do(func() {
		gLog.ErrorLogger = log.New(ErrorHandle, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	})
	gLog.ErrorLogger.Printf(format+"\n", v...)
}
