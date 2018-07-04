[![Go Report Card](https://goreportcard.com/badge/github.com/aellwein/slf4go-logrus-adaptor)](https://goreportcard.com/report/github.com/aellwein/slf4go-logrus-adaptor)
[![Coverage Status](https://img.shields.io/coveralls/github/aellwein/slf4go-logrus-adaptor/master.svg)](https://coveralls.io/github/aellwein/slf4go-logrus-adaptor?branch=master)
[![Build Status](https://img.shields.io/travis/aellwein/slf4go-native-adaptor/master.svg)](https://travis-ci.org/aellwein/slf4go-native-adaptor) 


# Logrus adaptor for SLF4GO

This is a [Logrus](https://github.com/sirupsen/logrus) adaptor implementation for [SLF4GO](https://github.com/aellwein/slf4go).

An example usage is stupid simple:

```go

package main

import (
	"github.com/aellwein/slf4go"
	_ "github.com/aellwein/slf4go-logrus-adaptor"
)

func main() {
    logger := slf4go.GetLogger("mylogger")
    
    logger.SetLevel(slf4go.LevelInfo)
    
    logger.Debug("don't see me!")
    logger.Info("It works!")
    logger.Warnf("The answer is %d", 42)
}
```
Note the underscore in front of the import of the SLF4GO adaptor.

You can change the logger implementation anytime, without changing the facade you are using, only by changing 
the imported adaptor.

# Development

* Install [Dep](https://github.com/golang/dep) tool.
* Type ``dep ensure``, so that all vendored packages can be fetched.
* use ``go build ./...`` as usual.