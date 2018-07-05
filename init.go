package slf4go_logrus_adaptor

import (
	"github.com/aellwein/slf4go"
)

// initialize logger by including this package.
func init() {
	slf4go.SetLoggerFactory(newLogrusLoggerFactory())
}
