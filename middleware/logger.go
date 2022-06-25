package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

// 自定义日志中间件
func Logger() gin.HandlerFunc {
	filePath := "log/log.log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err: ", err)
	}
	logger := logrus.New()

	logger.Out = scr
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spend := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/ 1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		// 拼接log格式
		entry := logger.WithFields(logrus.Fields{
			"HostName": hostName,
			"Status": statusCode,
			"SpendTime": spend,
			"Ip": clientIp,
			"Method": method,
			"Path": path,
			"Agent": userAgent,
		})
		// 系统内部有错误
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
