package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

//初始化
func init() {
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   "micro.log", //文件名称
			MaxSize:    50,          //单位MB
			MaxAge:     7,           //日志保留时间,(单位天)
			MaxBackups: 10,          //最大备份
			LocalTime:  true,        //本地时区
			Compress:   true,        //是否启用压缩
		})
	//编码
	encoder := zap.NewProductionEncoderConfig()
	//时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		// 编码器
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		//日志级别
		zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))
	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
