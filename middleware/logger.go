package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Log() gin.HandlerFunc{
	filePath := "log/log"

	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		fmt.Println("err:",err)
	}
	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7 * 24 * time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		)

	writerMap := lfshook.WriterMap{
		logrus.InfoLevel: 	logWriter,
		logrus.FatalLevel:	logWriter,
		logrus.DebugLevel: 	logWriter,
		logrus.WarnLevel:	logWriter,
		logrus.ErrorLevel: 	logWriter,
		logrus.PanicLevel:	logWriter,
	}

	hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostname, err := os.Hostname()
		if err != nil{
			hostname = "unkown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0{
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.URL

		entry := logger.WithFields(logrus.Fields{
			"Hostname": 	hostname,
			"Status":		statusCode,
			"SpendTime":	spendTime,
			"Ip":			clientIp,
			"Method":		method,
			"Path": 		path,
			"Agent": 		userAgent,
			"DataSize": 	dataSize,

		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500{
			entry.Error()
		}	else if statusCode >= 400 {
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}
