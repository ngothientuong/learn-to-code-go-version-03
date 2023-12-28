package main

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type MyStruct struct {
	Log logr.Logger
}

func (logger *MyStruct) logMssg(k, v string) {
	logger.Log.WithValues(k, v).Info("This is an example log message.")
}

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	myStruct := &MyStruct{
		Log: zapr.NewLogger(logger),
	}

	myStruct.logMssg("key", "value")
}
