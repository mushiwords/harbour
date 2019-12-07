/**
 * 日志模块
 **/

package mylog

import (
	"fmt"
	"log"
)

/**
 * 日志级别
 **/
const (
	MYLOG_DEBUG_LEVEL = iota
	MYLOG_INFO_LEVEL
	MYLOG_WORN_LEVEL
	MYLOG_ERROR_LEVEL
	MYLOG_FAIL_LEVEL
)

type MyLog struct {
	level  int
	logger *log.Logger
	access_logger *log.Logger
	info_logger   *log.Logger
}

func (m MyLog) LogDebug(format string, outpara ...interface{}) {
	if m.level <= MYLOG_DEBUG_LEVEL && m.info_logger != nil {
		str := fmt.Sprintf("[DEBUG] "+format, outpara...)
		m.info_logger.Printf(str)
	}
}

func (m MyLog) LogInfo(format string, outpara ...interface{}) {
	if m.level <= MYLOG_INFO_LEVEL && m.info_logger != nil {
		str := fmt.Sprintf("[INFO] "+format, outpara...)
		m.info_logger.Printf(str)
	}
}

func (m MyLog) LogWorn(format string, outpara ...interface{}) {
	if m.level <= MYLOG_WORN_LEVEL && m.info_logger != nil {
		str := fmt.Sprintf(" [WORN] "+format, outpara...)
		m.info_logger.Printf(str)
	}
}

func (m MyLog) LogError(format string, outpara ...interface{}) {
	if m.level <= MYLOG_ERROR_LEVEL && m.info_logger != nil {
		str := fmt.Sprintf("[ERROR] "+format, outpara...)
		m.info_logger.Printf(str)
	}
}

func (m MyLog) LogAccess(format string, outpara ...interface{}) {
	if m.access_logger != nil {
		str := fmt.Sprintf("[ACCESS] "+format, outpara...)
		m.access_logger.Printf(str)
	}
}

/*****************************************************
 * Common Functions
 *****************************************************/
func newLogFile(fileName string, maxAge int) *log.Logger {
	if fileName == "" {
		return nil
	}
	return log.New(newlogger(fileName, maxAge), "", log.LstdFlags)
}

/*****************************************************
 * for LocalLogger
 *****************************************************/

var DefaultLogger MyLog

func Init(accLogFile, infoLogFile, logLevel string, maxAge int) error {
	switch logLevel {
	case "debug":
		DefaultLogger.level = MYLOG_DEBUG_LEVEL
	case "info":
		DefaultLogger.level = MYLOG_INFO_LEVEL
	case "worn":
		DefaultLogger.level = MYLOG_WORN_LEVEL
	case "error":
		DefaultLogger.level = MYLOG_ERROR_LEVEL
	default:
		return fmt.Errorf("%v", "Invalid LogLevel")
	}

	DefaultLogger.access_logger = newLogFile(accLogFile, maxAge)
	DefaultLogger.info_logger = newLogFile(infoLogFile, maxAge)
	return nil
}

func LogDebug(format string, outpara ...interface{}) {
	DefaultLogger.LogDebug(format, outpara...)
}

func LogInfo(format string, outpara ...interface{}) {
	DefaultLogger.LogInfo(format, outpara...)
}

func LogError(format string, outpara ...interface{}) {
	DefaultLogger.LogError(format, outpara...)
}

func LogAccess(format string, outpara ...interface{}) {
	DefaultLogger.LogAccess(format, outpara...)
}
