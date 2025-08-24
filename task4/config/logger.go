package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Level uint8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = [...]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	mu     sync.Mutex
	level  Level
	logger *log.Logger
}

var (
	globalLogger *Logger
	once         sync.Once
)

func InitLogger(filePath string, level Level) {
	once.Do(func() {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("无法创建日志文件:", err)
		}
		globalLogger = &Logger{
			level:  level,
			logger: log.New(io.MultiWriter(file, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile),
		}

	})
}

func GetLogger() *Logger {
	if globalLogger == nil {
		InitLogger("app.log", INFO)
	}
	return globalLogger
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
	os.Exit(1)
}

func (l *Logger) log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.SetPrefix(levelNames[level] + " ")
	l.logger.Printf(format, v...)
}

func GinRouteLogger(l *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 记录路由信息
		routeInfo := fmt.Sprintf("Route: %s %s", c.Request.Method, c.Request.URL.Path)
		l.Debug(routeInfo)

		// 执行后续处理器
		c.Next()

		// 计算处理耗时
		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// 构建日志消息
		msg := fmt.Sprintf("[%s] %s %s %d %v", clientIP, method, path, status, latency)

		// 根据状态码记录不同级别的日志
		if status >= 500 {
			l.Error(msg)
		} else {
			l.Info(msg)
		}
	}
}
