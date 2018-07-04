package slf4go_logrus_adaptor

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aellwein/slf4go"
	"os"
	"testing"
)

func TestGetLogrusLogger(t *testing.T) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr, could also be a file.
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	// use defined logger factory
	l := log.New()
	l.WriterLevel(log.DebugLevel)

	slf4go.SetLoggerFactory(NewLogrusLoggerFactory(l))

	// do test
	logger := slf4go.GetLogger("test")
	logger.SetLevel(slf4go.LevelTrace)
	logger.Debug("are you prety?", true)
	logger.Debugf("are you prety? %t", true)
	logger.Info("how old are you? ", nil)
	logger.Infof("i'm %010d", 18)
	logger.Warn("you aren't honest! ")
	logger.Warnf("haha%02d", 1000, nil)
	logger.Trace("set level!!!!!!!")
	logger.SetLevel(slf4go.LevelWarn)
	logger.Trace("set level???")
	logger.Info("this should net appear.")
	logger.Error("what?")
	logger.Errorf("what?..$%s$", "XD")
	logger.Fatalf("import cycle not allowed! %s", "shit...")
	logger.Fatal("never reach here")
}

func TestLogrusPanic(t *testing.T) {
	logger := slf4go.GetLogger("test")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	logger.Panic("this causes panic!")
}
