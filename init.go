package slf4go_logrus_adaptor

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aellwein/slf4go"
	"os"
)

// initialize logger by including this package.
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr, could also be a file.
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	logger := log.New()
	// customize your root logger
	slf4go.SetLoggerFactory(NewLogrusLoggerFactory(logger))
}
