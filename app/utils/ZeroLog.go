package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func GetLog() zerolog.Logger {
	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	}

	// 时间戳：黑底白字
	out.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("\033[40;37m %s \033[0m", i)
	}

	// 日志级别：不同级别不同底色，黑字
	out.FormatLevel = func(i interface{}) string {
		lvl := strings.ToUpper(fmt.Sprintf("%s", i))
		var bg string
		switch lvl {
		case "DEBUG":
			bg = "\033[44m" // 蓝底
		case "INFO":
			bg = "\033[42m" // 绿底
		case "WARN", "WARNING":
			bg = "\033[43m" // 黄底
		case "ERROR", "FATAL", "PANIC":
			bg = "\033[41m" // 红底
		default:
			bg = "\033[47m" // 白底
		}
		return fmt.Sprintf("%s\033[30m %s \033[0m", bg, lvl)
	}

	// 模块名（或者包名）：紫底白字，可换成你喜欢的字段
	out.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\033[45;37m %s \033[0m", i)
	}

	// 字段值：保持和背景分离，直接黑字
	out.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf(" %v", i)
	}

	// 消息前加箭头，白底黑字
	out.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("\033[47;30m ▶ %s \033[0m", i)
	}

	return zerolog.New(out).With().Timestamp().Logger()
}
