package app

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	Log = zap.NewExample().Sugar()
	Log.Infow("log initialized")
}
