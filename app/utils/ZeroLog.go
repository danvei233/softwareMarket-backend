package utils

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var log zerolog.Logger

func FormatTimestamp(i interface{}) string {

	return color.Sprintf("[%s]", i)

}
func FormatFieldName(i interface{}) string {
	return color.New(color.BgMagenta, color.FgWhite).Sprintf(" %s ", i)
}

func FormatFieldValue(i interface{}) string {
	return color.Sprintf(" %v", i)
}

func FormatMessage(i interface{}) string {
	return color.Sprintf(" ▶ %s ", i)
}
func FormatLevel(i interface{}) string {
	lvl := strings.ToUpper(fmt.Sprintf("%s", i))
	var bg color.Color
	var k color.Color
	switch lvl {
	case "DEBUG":
		bg = color.BgHiBlue // 蓝底
		k = color.FgLightBlue
	case "INFO":
		bg = color.BgGreen // 绿底
		k = color.FgGreen
	case "WARN", "WARNING":
		bg = color.BgYellow // 黄底
		k = color.FgYellow
	case "ERROR", "FATAL", "PANIC":
		bg = color.BgRed // 红底
		k = color.FgRed
	default:
		bg = color.BgGreen // 白底
		k = color.FgGreen
	}

	return color.New(k).Sprintf("\uE0B2") + color.New(bg).Sprintf("  %s  ", lvl) + color.New(k).Sprintf("\uE0B0")
}
func InitLog() {
	log = zerolog.New(zerolog.ConsoleWriter{
		Out:              os.Stdout,
		NoColor:          false,
		TimeFormat:       time.RFC3339,
		FormatLevel:      FormatLevel,
		FormatTimestamp:  FormatTimestamp,
		FormatMessage:    FormatMessage,
		FormatFieldName:  FormatFieldName,
		FormatFieldValue: FormatFieldValue,
	}).With().Timestamp().Logger()
}
func GetLog() zerolog.Logger {
	return log
}
