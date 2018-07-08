package slf4go_logrus_adaptor

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/aellwein/slf4go"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := slf4go.GetLogger("test")

	var levels = []slf4go.LogLevel{
		slf4go.LevelPanic,
		slf4go.LevelFatal,
		slf4go.LevelError,
		slf4go.LevelWarn,
		slf4go.LevelInfo,
		slf4go.LevelDebug,
		slf4go.LevelTrace,
	}

	for _, i := range levels {
		logger.SetLevel(i)
		logger.Trace("Trace")
		logger.Tracef("Tracef: %v", logger)
		logger.Debug("Debug")
		logger.Debugf("Debugf: %s", "debug mode")
		logger.Info("Info")
		logger.Infof("Infof: %v", slf4go.GetLoggerFactory())
		logger.Warn("Warn")
		logger.Warnf("Warnf: %d", 42)
		logger.Error("Error")
		logger.Errorf("Errorf: %v", errors.New("some error"))
	}
}

func TestLoggerFatal(t *testing.T) {
	mockExit := func(int) {
		panic("mockExit called")
	}
	patch := monkey.Patch(os.Exit, mockExit)
	defer patch.Unpatch()

	logger := slf4go.GetLogger("test")
	underTest := func() {
		logger.Fatal("fatality!")
	}

	assert.Panics(t, underTest, "should panic")
}

func TestLoggerFatalf(t *testing.T) {
	mockExit := func(int) {
		panic("mockExit called")
	}
	patch := monkey.Patch(os.Exit, mockExit)
	defer patch.Unpatch()

	logger := slf4go.GetLogger("test")
	underTest := func() {
		logger.Fatalf("fatality: %d", 42)
	}

	assert.Panics(t, underTest, "should panic")
}

func TestLoggerPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	logger := slf4go.GetLogger("test")
	logger.Panic("this is expected to cause panic")
}

func TestLoggerPanicf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	logger := slf4go.GetLogger("test")
	logger.Panicf("this is expected to cause panic: %d", 42)
}

type ParamTester struct {
	name      string
	goodValue interface{}
	badValue  interface{}
}

type mockHook struct {
}

func (*mockHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (*mockHook) Fire(*logrus.Entry) error {
	return nil
}

func TestSetLoggingParametersAutomatically(t *testing.T) {
	var params = []ParamTester{
		{
			name:      "formatter",
			goodValue: new(logrus.JSONFormatter),
			badValue:  1337,
		},
		{
			name:      "fields",
			goodValue: logrus.Fields{"foo": "bar"},
			badValue:  1337,
		},
		{
			name:      "output",
			goodValue: os.Stdout,
			badValue:  1337,
		},
		{
			name:      "level",
			goodValue: logrus.ErrorLevel,
			badValue:  1337,
		},
		{
			name:      "hooks",
			goodValue: []logrus.Hook{&mockHook{}},
			badValue:  1337,
		},
	}

	for _, p := range params {
		goodTest := func(pt *ParamTester) {
			prm := slf4go.LoggingParameters{pt.name: pt.goodValue}
			err := slf4go.GetLoggerFactory().SetLoggingParameters(prm)
			if err != nil {
				panic(err)
			}
		}
		badTest := func(pt *ParamTester) {
			prm := slf4go.LoggingParameters{pt.name: pt.badValue}
			err := slf4go.GetLoggerFactory().SetLoggingParameters(prm)
			if err == nil {
				panic(err)
			}
		}
		assert.NotPanics(t, func() { goodTest(&p) })
		assert.NotPanics(t, func() { badTest(&p) })
	}

	// also test an unknown parameter
	assert.Panics(t, func() {
		if err := slf4go.GetLoggerFactory().SetLoggingParameters(
			slf4go.LoggingParameters{"xyzunknown": "blah"}); err != nil {
			panic(err)
		}
	})

	// ...and the branch with no params
	assert.Nil(t, slf4go.GetLoggerFactory().SetLoggingParameters(slf4go.LoggingParameters{}))
	assert.Nil(t, slf4go.GetLoggerFactory().SetLoggingParameters(nil))
}

func TestLoggingExtendedWithParams(t *testing.T) {
	slf4go.GetLoggerFactory().SetLoggingParameters(
		slf4go.LoggingParameters{
			"fields":    logrus.Fields{"foo": "bar"},
			"output":    os.Stderr,
			"formatter": &logrus.JSONFormatter{},
			"level":     logrus.DebugLevel,
			"hooks": []logrus.Hook{
				&mockHook{},
			},
		},
	)
	logger := slf4go.GetLogger("test")
	logger.Infof("logging using custom fields: %v", t)
}

func TestLogrusLoggerFactory_GetDefaultLogLevel(t *testing.T) {
	assert.Equal(t, slf4go.GetLoggerFactory().GetDefaultLogLevel(), slf4go.LevelInfo)
}

func TestLogrusLoggerFactory_SetDefaultLogLevel(t *testing.T) {
	slf4go.GetLoggerFactory().SetDefaultLogLevel(slf4go.LevelTrace)
	logger := slf4go.GetLogger("test")
	assert.True(t, logger.IsTraceEnabled())
}
