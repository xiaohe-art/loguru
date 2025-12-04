package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L zerolog.Logger

// Init 初始化日志系统
// logPath 参数示例："/var/log/myapp/app.log"
func Init(isDev bool, logPath string) {

	// --- 创建日志目录 ---
	if logPath != "" {
		if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
			panic("无法创建日志目录: " + err.Error())
		}
	}

	// --- 文件日志 writer（自动切割） ---
	fileWriter := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    20, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	}

	// --- 控制台输出 ---
	var consoleWriter io.Writer
	if isDev {
		consoleWriter = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		}
	} else {
		consoleWriter = os.Stdout
	}

	// --- 输出合并：控制台 + 文件 ---
	multiWriter := io.MultiWriter(consoleWriter, fileWriter)

	// --- 初始化 Logger ---
	if isDev {
		L = zerolog.New(multiWriter).With().Timestamp().Logger()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		L = zerolog.New(multiWriter).With().Timestamp().Caller().Logger()
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Logger = L
}

//
// -------- loguru 风格 API 封装 --------
//

func Debug(v any, args ...any) {
	message := toString(v)
	if len(args) > 0 {
		for _, arg := range args {
			message += " " + toString(arg)
		}
	}
	L.Debug().Msg(message)
}

func Info(v any, args ...any) {
	message := toString(v)
	if len(args) > 0 {
		for _, arg := range args {
			message += " " + toString(arg)
		}
	}
	L.Info().Msg(message)
}

func Warn(v any, args ...any) {
	message := toString(v)
	if len(args) > 0 {
		for _, arg := range args {
			message += " " + toString(arg)
		}
	}
	L.Warn().Msg(message)
}

func Error(v any, args ...any) {
	message := toString(v)
	if len(args) > 0 {
		for _, arg := range args {
			message += " " + toString(arg)
		}
	}
	L.Error().Msg(message)
}

func Fatal(v any, args ...any) {
	message := toString(v)
	if len(args) > 0 {
		for _, arg := range args {
			message += " " + toString(arg)
		}
	}
	L.Fatal().Msg(message)
}
