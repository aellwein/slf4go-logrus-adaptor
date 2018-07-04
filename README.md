# Logrus adaptor for slf4go

This is a [Logrus](https://github.com/sirupsen/logrus) adaptor implementation for [slf4go](https://github.com/aellwein/slf4go).

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