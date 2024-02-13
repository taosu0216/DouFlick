package log

import (
	"favoritesvr/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumber "gopkg.in/natefinch/lumberjack.v2"
)

var (
	log   *zap.Logger
	sugar *zap.SugaredLogger
)

func InitLog() {
	var coreArr []zapcore.Core

	// 获取编码器
	// NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	// 指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//显示完整文件路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	// error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	// info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if config.GetGlobalConfig().LogConfig.Level == "debug" {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		} else {
			return lev < zap.ErrorLevel && lev >= zap.InfoLevel
		}
	})

	// info文件writeSyncer
	logConfig := config.GetGlobalConfig().LogConfig
	infoFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   logConfig.LogPath + "info_" + logConfig.FileName, // 日志文件存放目录，
		MaxSize:    logConfig.MaxSize,                                // 文件大小限制,单位MB
		MaxBackups: logConfig.MaxBackups,                             // 最大保留日志文件数量
		MaxAge:     logConfig.MaxAge,                                 // 日志文件保留天数
		Compress:   false,
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   logConfig.LogPath + "error_" + logConfig.FileName, // 日志文件存放目录
		MaxSize:    logConfig.MaxSize,                                 // 文件大小限制,单位MB
		MaxBackups: logConfig.MaxBackups,                              // 最大保留日志文件数量
		MaxAge:     logConfig.MaxAge,                                  // 日志文件保留天数
		Compress:   false,                                             // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	// zap.AddCaller()为显示文件名和行号，可省略
	log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(1))
	sugar = log.Sugar()
}

func Infof(s string, v ...interface{}) {
	sugar.Infof(s, v...)
}

func Infow(s string, v ...interface{}) {
	sugar.Infow(s, v...)
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Debugf(s string, v ...interface{}) {
	sugar.Debugf(s, v...)
}

func Debugw(s string, v ...interface{}) {
	sugar.Debugw(s, v...)
}

func Debug(v ...interface{}) {
	sugar.Debug(v...)
}

func Errorf(s string, v ...interface{}) {
	sugar.Errorf(s, v...)
}

func Errorw(s string, v ...interface{}) {
	sugar.Errorw(s, v...)
}

func Error(v ...interface{}) {
	sugar.Error(v...)
}

func Fatalf(s string, v ...interface{}) {
	sugar.Fatalf(s, v...)
}

func Fatalw(s string, v ...interface{}) {
	sugar.Fatalw(s, v...)
}

func Fatal(v ...interface{}) {
	sugar.Error(v...)
}

func Sync() {
	log.Sync()
}
