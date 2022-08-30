package logconfig

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// CustomLogFormat 自定义日志格式
func CustomLogFormat() {
	logWriter, err := RotateLog()
	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: logWriter,
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.FatalLevel: logWriter,
		log.PanicLevel: logWriter,
	}, &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	log.SetReportCaller(false)
	log.AddHook(lfHook)
}
