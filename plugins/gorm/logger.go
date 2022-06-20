package gorm

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/xiaohubai/go-layout/configs/global"
	"go.uber.org/zap"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type cfg struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}
type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: t.Interface, BeginAt: time.Now()}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), cfg{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), cfg{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	Recorder = traceRecorder{Interface: Default, BeginAt: time.Now()}
)

type _logger struct {
	cfg
	logger.Writer
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(writer logger.Writer, c cfg) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s\n"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s\n"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s\n"
	)

	if c.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s\n" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
	}

	return &_logger{
		Writer:       writer,
		cfg:          c,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// LogMode log mode
func (c *_logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *c
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (c *_logger) Info(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.Printf(ctx, "info", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (c *_logger) Warn(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.Printf(ctx, "warn", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (c *_logger) Error(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.Printf(ctx, "error", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (c *_logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && c.LogLevel >= logger.Error:
			sql, rows := fc()
			c.Printf(ctx, "error", fmt.Sprintf("fileLine:%s,elapsed:%f,row:%d,sql:%s,err:%s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql, err))
		case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", c.SlowThreshold)
			c.Printf(ctx, "warn", fmt.Sprintf("fileLine:%s,elapsed:%f,row:%d,sql:%s,slowLog:%s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql, slowLog))
		case c.LogLevel >= logger.Info:
			sql, rows := fc()
			c.Printf(ctx, "info", fmt.Sprintf("fileLine:%s,elapsed:%f,row:%d,sql:%s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	}
}

func (c *_logger) Printf(ctx context.Context, key string, data ...interface{}) {
	traceId := ctx.Value("trace_id")
	if traceId == nil {
		return
	}
	if global.Cfg.Mysql.LogZap {
		switch {
		case key == "error":
			global.Log.Error(traceId.(string), zap.Any("key", "sql"), zap.Any("msg", fmt.Sprintf("%s:%s", "sql", data)))
		default:
			global.Log.Info(traceId.(string), zap.Any("key", "sql"), zap.Any("msg", fmt.Sprintf("%s:%s", "sql", data)))
		}
	} else {
		c.Writer.Printf(traceId.(string), data...)
	}
}
