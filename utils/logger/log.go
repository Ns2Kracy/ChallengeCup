package logger

import "github.com/kataras/iris/v12"

func Info(args ...interface{}) {
	iris.New().Logger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	iris.New().Logger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	iris.New().Logger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	iris.New().Logger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	iris.New().Logger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	iris.New().Logger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	iris.New().Logger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	iris.New().Logger().Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	iris.New().Logger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	iris.New().Logger().Debugf(format, args...)
}
