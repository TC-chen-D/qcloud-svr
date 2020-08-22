package utils

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// Init log
func (c *Conf) InitLogs() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	hook := c.newLfsHook()
	log.AddHook(hook)
}

// Set the log split and storage mode
func (c *Conf) newLfsHook() log.Hook {
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s/%s", c.LogConf.LogPath, c.LogConf.LogFile),
		rotatelogs.WithRotationTime(time.Duration(c.LogConf.RotationTime)*time.Hour),
		rotatelogs.WithRotationCount(c.LogConf.RotationCount),
	)
	if err != nil {
		log.Fatalf("configuring log file fail %v", err)
	}

	level, err := log.ParseLevel(c.LogConf.LogLevel)
	if err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}
