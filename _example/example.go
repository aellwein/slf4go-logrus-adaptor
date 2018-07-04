package main

import (
	_ "github.com/aellwein/slf4go-logrus-adaptor"
	"github.com/aellwein/slf4go"
)

func main() {
	logger := slf4go.GetLogger("mylogger")

	logger.SetLevel(slf4go.LevelInfo)

	logger.Debug("don't see me!")
	logger.Info("It works!")
	logger.Warnf("The answer is %d", 42)
}
