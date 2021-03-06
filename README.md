[![Go Report Card](https://goreportcard.com/badge/github.com/aellwein/slf4go-logrus-adaptor)](https://goreportcard.com/report/github.com/aellwein/slf4go-logrus-adaptor)
[![Coverage Status](https://img.shields.io/coveralls/github/aellwein/slf4go-logrus-adaptor/master.svg)](https://coveralls.io/github/aellwein/slf4go-logrus-adaptor?branch=master)
[![Build Status](https://img.shields.io/travis/aellwein/slf4go-logrus-adaptor/master.svg)](https://travis-ci.org/aellwein/slf4go-logrus-adaptor) 


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

# Logging parameters

This adaptor supports several parameters, available with ``SetLoggingParameters``:


 Parameter Key     | Value Type                        | Description
-------------------|-----------------------------------|----------------------------------
 "formatter"       | ``logrus.Formatter``              | Logrus Formatter to use
 "output"          | ``io.Writer``                     | Ouput writer as defined by Logrus
 "level"           | ``logrus.Level``                  | Log level to use
 "fields"          | ``logrus.Fields``                 | A map containing fields to be included in output
 "hooks"           | ``[]logrus.Hook``                 | List of hooks to be used with Logrus

# Development

* use ``go build ./...`` as usual.