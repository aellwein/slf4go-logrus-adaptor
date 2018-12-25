package slf4go_logrus_adaptor

import (
	"errors"
	"fmt"
	"io"

	"github.com/aellwein/slf4go"
	log "github.com/sirupsen/logrus"
)

// facade for logrus
type loggerAdaptorLogrus struct {
	slf4go.LoggerAdaptor
	entry *log.Entry
}

func newLogrusLogger(name string, entry *log.Entry, lvl slf4go.LogLevel) *loggerAdaptorLogrus {
	result := &loggerAdaptorLogrus{}
	result.entry = entry.WithField("name", name)
	result.SetName(name)
	result.SetLevel(lvl)
	return result
}

func (lgr *loggerAdaptorLogrus) SetLevel(l slf4go.LogLevel) {
	lgr.LoggerAdaptor.SetLevel(l)
	switch l {
	case slf4go.LevelTrace:
		lgr.entry.Logger.Level = log.DebugLevel
	case slf4go.LevelDebug:
		lgr.entry.Logger.Level = log.DebugLevel
	case slf4go.LevelInfo:
		lgr.entry.Logger.Level = log.InfoLevel
	case slf4go.LevelWarn:
		lgr.entry.Logger.Level = log.WarnLevel
	case slf4go.LevelError:
		lgr.entry.Logger.Level = log.ErrorLevel
	case slf4go.LevelFatal:
		lgr.entry.Logger.Level = log.FatalLevel
	case slf4go.LevelPanic:
		lgr.entry.Logger.Level = log.PanicLevel
	}
}
func (lgr *loggerAdaptorLogrus) Trace(args ...interface{}) {
	// forward to Debug
	lgr.entry.Debug(args...)
}

func (lgr *loggerAdaptorLogrus) Tracef(format string, args ...interface{}) {
	// forward to Debug
	lgr.entry.Debugf(format, args...)
}

func (lgr *loggerAdaptorLogrus) Debug(args ...interface{}) {
	lgr.entry.Debugln(args...)
}

func (lgr *loggerAdaptorLogrus) Debugf(format string, args ...interface{}) {
	lgr.entry.Debugf(format, args...)
}

func (lgr *loggerAdaptorLogrus) Info(args ...interface{}) {
	lgr.entry.Infoln(args...)
}

func (lgr *loggerAdaptorLogrus) Infof(format string, args ...interface{}) {
	lgr.entry.Infof(format, args...)
}

func (lgr *loggerAdaptorLogrus) Warn(args ...interface{}) {
	lgr.entry.Warnln(args...)
}

func (lgr *loggerAdaptorLogrus) Warnf(format string, args ...interface{}) {
	lgr.entry.Warnf(format, args...)
}

func (lgr *loggerAdaptorLogrus) Error(args ...interface{}) {
	lgr.entry.Errorln(args...)
}

func (lgr *loggerAdaptorLogrus) Errorf(format string, args ...interface{}) {
	lgr.entry.Errorf(format, args...)
}

func (lgr *loggerAdaptorLogrus) Fatal(args ...interface{}) {
	lgr.entry.Fatalln(args...)
}

func (lgr *loggerAdaptorLogrus) Fatalf(format string, args ...interface{}) {
	lgr.entry.Fatalf(format, args...)
}

func (lgr *loggerAdaptorLogrus) Panic(args ...interface{}) {
	lgr.entry.Panicln(args...)
}

func (lgr *loggerAdaptorLogrus) Panicf(format string, args ...interface{}) {
	lgr.entry.Panicf(format, args...)
}

// internal LoggerFactory for logrus
type logrusLoggerFactory struct {
	level slf4go.LogLevel
	entry *log.Entry
}

func newLogrusLoggerFactory() slf4go.LoggerFactory {
	factory := &logrusLoggerFactory{level: slf4go.LevelInfo}
	factory.entry = log.NewEntry(log.New())
	return factory
}

func (factory *logrusLoggerFactory) GetLogger(name string) slf4go.Logger {
	return newLogrusLogger(name, factory.entry, factory.level)
}

func (factory *logrusLoggerFactory) SetDefaultLogLevel(l slf4go.LogLevel) {
	factory.level = l
}

func (factory *logrusLoggerFactory) GetDefaultLogLevel() slf4go.LogLevel {
	return factory.level
}

func (factory *logrusLoggerFactory) SetLoggingParameters(params slf4go.LoggingParameters) error {

	logger := log.New()
	var entry *log.Entry

	for k, v := range params {
		switch k {
		case "formatter":
			if fmter, ok := v.(log.Formatter); !ok {
				return errors.New("invalid type for parameter 'formatter', should be of type logrus.Formatter")
			} else {
				logger.Formatter = fmter
			}
		case "output":
			if writer, ok := v.(io.Writer); !ok {
				return errors.New("invalid type for parameter 'output', should be of type io.Writer")
			} else {
				logger.Out = writer
			}
		case "level":
			if lvl, ok := v.(log.Level); !ok {
				return errors.New("invalid type for parameter 'level', should be of type logrus.Level")
			} else {
				logger.Level = lvl
			}
		case "fields":
			if fields, ok := v.(log.Fields); !ok {
				return errors.New("invalid type for parameter 'fields', should be of type logrus.Fields")
			} else {
				entry = logger.WithFields(fields)
			}
		case "hooks":
			if hooks, ok := v.([]log.Hook); !ok {
				return errors.New("invalid type for parameter 'hooks', should be of type []logrus.Hook")
			} else {
				for _, hook := range hooks {
					logger.AddHook(hook)
				}
			}
		default:
			return fmt.Errorf("unsupported parameter: %v", k)
		}
	}
	if entry != nil {
		factory.entry = entry
	}
	return nil
}
