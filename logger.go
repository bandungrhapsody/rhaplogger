package rhaplogger

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type LogModel struct {
	Level      string    `json:"level"`
	Timestamp  time.Time `json:"timestamp"`
	RequestURI string    `json:"request_uri"`
	StatusCode int       `json:"status_code"`
	Duration   int64     `json:"duration"`
	Message    string    `json:"message"`
	Version    string    `json:"version"`
	ErrorCode  string    `json:"error_code"`
}

type LogConfig struct {
	Filename string
	Stdout   bool
	Prefix   string
}

type RhapLogger struct {
	logger *log.Logger
}

func (*RhapLogger) getDefaultLogModel(level string) *LogModel {
	return &LogModel{
		Level:      level,
		Timestamp:  time.Now(),
		RequestURI: "-",
		StatusCode: 0,
		Duration:   0,
		Message:    "-",
		Version:    "-",
		ErrorCode:  "-",
	}
}

func NewRhapLogger(config *LogConfig) (rhapLogger *RhapLogger, err error) {
	var writers []io.Writer

	if config == nil {
		return &RhapLogger{
			logger: log.New(os.Stdout, "", 0),
		}, nil
	}

	if config.Filename != "" {
		var file *os.File

		file, err = os.OpenFile(config.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}

		writers = append(writers, file)
	}

	if config.Stdout {
		writers = append(writers, os.Stdout)
	}

	rhapLogger = &RhapLogger{logger: log.New(io.MultiWriter(writers...), config.Prefix, 0)}
	return
}

func (rl *RhapLogger) Logger() *log.Logger {
	return rl.logger
}

func (rl *RhapLogger) Print(model *LogModel) {
	data, _ := json.Marshal(model)

	rl.logger.Print(string(data))
}
