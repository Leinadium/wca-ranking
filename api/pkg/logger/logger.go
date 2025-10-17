package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func createStyle() *log.Styles {
	styles := log.DefaultStyles()

	// error
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Background(lipgloss.Color("#fd0000")).
		Foreground(lipgloss.Color("#200000"))
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#fd0000")).Bold(true)

	// warning
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Background(lipgloss.Color("#fdfd00")).
		Foreground(lipgloss.Color("#202000"))

	// info
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Background(lipgloss.Color("#0000fd")).
		Foreground(lipgloss.Color("#000020"))

	// debug
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Background(lipgloss.Color("#00fd00")).
		Foreground(lipgloss.Color("#002000"))

	return styles
}

type Level log.Level

const (
	DEBUG Level = Level(log.DebugLevel)
	INFO  Level = Level(log.InfoLevel)
	WARN  Level = Level(log.WarnLevel)
	ERROR Level = Level(log.ErrorLevel)
)

var defaultStyle *log.Styles = createStyle()

type Options struct {
	Level Level
}

func New(name string, opt *Options) *log.Logger {
	if opt == nil {
		opt = &Options{Level: INFO}
	}

	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:           log.Level(opt.Level),
		Prefix:          fmt.Sprintf("[%s]", name),
		TimeFormat:      time.DateTime,
		ReportTimestamp: true,
	})
	logger.SetStyles(defaultStyle)
	return logger
}
