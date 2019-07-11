package app

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

var Log *zap.Logger

var Name = func() string {
	ret := os.Args[0]
	for _, x := range os.Args[1:] {
		if x[0] == '-' {
			break
		}
		ret = x
	}
	return filepath.Base(ret)
}()

func init() {
	var err error
	Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Log.Info("log initialized", zap.String("package", "app"))
}

func LoggerWith(fields ...zap.Field) *zap.Logger {
	println("LoggerWith")
	return Log.With(append(fields, zap.String("app", Name))...)
}
