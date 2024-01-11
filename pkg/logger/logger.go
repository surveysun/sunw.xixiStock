package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
)

var Log *Logger

type Logger struct {
	file *os.File
}

func NewLogger(logFile string, isPrint bool) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if isPrint {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.SetOutput(file)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	Log = &Logger{
		file: file,
	}
}

func (l *Logger) SetLogLevel(level log.Level) {
	log.SetLevel(level)
}

func (l *Logger) Debug(message string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Debug(message)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Debugf(format, args...)
}

func (l *Logger) Info(message string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Info(message)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Infof(format, args...)
}

func (l *Logger) Error(message string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(message)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Errorf(format, args...)
}

func (l *Logger) Warn(message string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Warning(message)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Warningf(format, args...)
}

func (l *Logger) Trace(message string) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Trace(message)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Tracef(format, args...)
}

func (l *Logger) Close() {
	l.file.Close()
}
