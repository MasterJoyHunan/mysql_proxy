package logger

import (
	"log"

	"github.com/MasterJoyHunan/mysql_proxy/config"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Setup() {
	level, err := logrus.ParseLevel(config.LogConf.Level)
	if err != nil {
		log.Panic("日志level格式设置错误", err)
	}
	Logger.SetLevel(level)

	//设置日志格式
	//Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceQuote:   false,
		DisableQuote: true,
	}) // 默认格式 无需显示设置

	// 取消线程安全
	Logger.SetNoLock()

	// 打印堆栈信息
	//Logger.ReportCaller = true

	Logger.AddHook(&MyHook{})
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}
func Error(args ...interface{}) {
	Logger.Error(args...)
}
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Warningf(format string, args ...interface{}) {
	Logger.Warningf(format, args...)
}
func Warning(args ...interface{}) {
	Logger.Warning(args...)
}
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}
func Info(args ...interface{}) {
	Logger.Info(args...)
}
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}
func Tracef(format string, args ...interface{}) {
	Logger.Tracef(format, args...)
}
func Trace(args ...interface{}) {
	Logger.Trace(args...)
}
