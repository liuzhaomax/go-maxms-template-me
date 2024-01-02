package core

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"strings"
)

var LoggerSet = wire.NewSet(wire.Struct(new(Logger), "*"), wire.Bind(new(ILogger), new(*Logger)))

// ILogger 主要用于中间件和handler
type ILogger interface {
	Succeed(string)
	Fail(Code, string, error)
	SucceedWithField(*gin.Context, string)
	FailWithField(*gin.Context, Code, string, error)
}

const callerLevel = 4 // 调用堆栈第n层

type Logger struct {
	Logger *logrus.Logger
}

func (l *Logger) Succeed(desc string) {
	l.Logger.Debug(FormatCaller(true, GetCallerFileAndLine(callerLevel)))
	l.Logger.Info(FormatInfo(desc))
}

func (l *Logger) Fail(code Code, desc string, err error) {
	l.Logger.Info(FormatCaller(false, GetCallerFileAndLine(callerLevel)))
	l.Logger.Info(FormatError(code, desc, err))
}

func (l *Logger) SucceedWithField(c *gin.Context, desc string) {
	l.Logger.WithField(strings.ToLower(TraceId), c.Request.Header.Get(TraceId)).Debug(FormatCaller(true, GetCallerFileAndLine(callerLevel)))
	l.Logger.WithField(strings.ToLower(TraceId), c.Request.Header.Get(TraceId)).Info(FormatInfo(desc))
}

func (l *Logger) FailWithField(c *gin.Context, code Code, desc string, err error) {
	l.Logger.WithField(strings.ToLower(TraceId), c.Request.Header.Get(TraceId)).Info(FormatCaller(false, GetCallerFileAndLine(callerLevel)))
	l.Logger.WithField(strings.ToLower(TraceId), c.Request.Header.Get(TraceId)).Info(FormatError(code, desc, err))
}
