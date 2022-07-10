package customlogger

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type logger struct {
	echoLog   echo.Logger
	requestID string
}

func NewContextLogger(echoLog echo.Logger) *logger {
	return &logger{
		echoLog: echoLog,
	}
}

func (l *logger) SetRequestID(requestID string) {
	l.requestID = requestID
}

func (l *logger) Output() io.Writer {
	return l.echoLog.Output()
}
func (l *logger) SetOutput(w io.Writer) {
	l.echoLog.SetOutput(w)
}
func (l *logger) Prefix() string {
	return l.echoLog.Prefix()
}
func (l *logger) SetPrefix(p string) {
	l.echoLog.SetPrefix(p)
}
func (l *logger) Level() log.Lvl {
	return l.echoLog.Level()
}
func (l *logger) SetLevel(v log.Lvl) {
	l.echoLog.SetLevel(v)
}
func (l *logger) SetHeader(h string) {
	l.echoLog.SetHeader(h)
}

func (l *logger) Print(i ...interface{}) {
	l.echoLog.Print(i...)
}
func (l *logger) Printf(format string, i ...interface{}) {
	l.echoLog.Printf(format, i...)
}
func (l *logger) Printj(j log.JSON) {
	l.echoLog.Printj(j)
}
func (l *logger) Debug(i ...interface{}) {
	l.echoLog.Debug(i...)
}
func (l *logger) Debugf(format string, i ...interface{}) {
	l.echoLog.Debugf(format, i...)
}
func (l *logger) Debugj(j log.JSON) {
	l.echoLog.Debugj(j)
}
func (l *logger) Info(i ...interface{}) {
	l.echoLog.Infoj(log.JSON{
		"id":      l.requestID,
		"message": i[0],
		"args":    i[1:],
	})
}
func (l *logger) Infof(format string, i ...interface{}) {
	l.echoLog.Infof(format, i...)
}
func (l *logger) Infoj(j log.JSON) {
	l.echoLog.Infoj(j)
}
func (l *logger) Warn(i ...interface{}) {
	l.echoLog.Warnj(log.JSON{
		"id":      l.requestID,
		"message": i[0],
		"args":    i[1:],
	})
}
func (l *logger) Warnf(format string, i ...interface{}) {
	l.echoLog.Warnf(format, i...)
}
func (l *logger) Warnj(j log.JSON) {
	l.echoLog.Warnj(j)
}
func (l *logger) Error(i ...interface{}) {
	l.echoLog.Errorj(log.JSON{
		"id":      l.requestID,
		"message": i[0],
		"error":   i[1],
		"args":    i[2:],
	})
}
func (l *logger) Errorf(format string, i ...interface{}) {
	l.echoLog.Errorf(format, i...)
}
func (l *logger) Errorj(j log.JSON) {
	l.echoLog.Errorj(j)
}
func (l *logger) Fatal(i ...interface{}) {
	l.echoLog.Fatal(i...)
}
func (l *logger) Fatalj(j log.JSON) {
	l.echoLog.Fatalj(j)
}
func (l *logger) Fatalf(format string, i ...interface{}) {
	l.echoLog.Fatalf(format, i...)
}
func (l *logger) Panic(i ...interface{}) {
	l.echoLog.Panic(i...)
}
func (l *logger) Panicj(j log.JSON) {
	l.echoLog.Panicj(j)
}
func (l *logger) Panicf(format string, i ...interface{}) {
	l.echoLog.Panicf(format, i...)
}
