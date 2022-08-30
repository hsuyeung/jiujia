package logconfig

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	LogDir   = "./logs/log-%Y%m%d.log" // 日志存放目录
	OneDay   = 24 * time.Hour          // 1 天
	OneMonth = 30 * OneDay             // 30 天
)

// RotateLog 切割 Info 级别的日志
func RotateLog() (w io.Writer, err error) {
	w, err = rotatelogs.New(
		LogDir,
		rotatelogs.WithMaxAge(OneMonth),
		rotatelogs.WithRotationTime(OneDay),
	)
	if err != nil {
		return nil, err
	}
	return w, nil
}
