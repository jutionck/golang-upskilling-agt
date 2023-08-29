package loggerutil

import (
	"os"

	"github.com/jutionck/golang-upskilling-agt/config"
	modelutil "github.com/jutionck/golang-upskilling-agt/utils/model_util"
	"github.com/sirupsen/logrus"
)

type LoggerUtil interface {
	InitialLoggerFile() error
	ReqLogInfo(requestLog modelutil.RequestLog)
	ReqLogError(requestLog modelutil.RequestLog)
	ResLogInfo(responseLog modelutil.ResponseLog)
	ResLogError(responseLog modelutil.ResponseLog)
}

type loggerUtil struct {
	cfg config.FileConfig
	log *logrus.Logger
}

// InitialLoggerFile implements LoggerUtil.
func (l *loggerUtil) InitialLoggerFile() error {
	file, err := os.OpenFile(l.cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	} else {
		l.log = logrus.New()
		l.log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		l.log.Out = file
	}
	return nil
}

// ReqLogError implements LoggerUtil.
func (l *loggerUtil) ReqLogError(requestLog modelutil.RequestLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Request Log",
	}).Error(requestLog)
}

// ReqLogInfo implements LoggerUtil.
func (l *loggerUtil) ReqLogInfo(requestLog modelutil.RequestLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Request Log",
	}).Info(requestLog)
}

// ResLogError implements LoggerUtil.
func (l *loggerUtil) ResLogError(responseLog modelutil.ResponseLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Response Log",
	}).Error(responseLog)
}

// ResLogInfo implements LoggerUtil.
func (l *loggerUtil) ResLogInfo(responseLog modelutil.ResponseLog) {
	l.log.WithFields(logrus.Fields{
		"message": "Response Log",
	}).Info(responseLog)
}

func NewLoggerUtil(cfg config.FileConfig) LoggerUtil {
	return &loggerUtil{cfg: cfg}
}
