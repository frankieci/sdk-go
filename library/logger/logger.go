package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gitlab.com/sdk-go/library/file"
)

type settings struct {
	Path       string `yaml:"path"`
	Name       string `yaml:"name"`
	Ext        string `yaml:"ext"`
	TimeFormat string `yaml:"time-format"`
}

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup(path, name string) {
	settings := &settings{Path: path, Name: name, Ext: "log", TimeFormat: "2006-01-02"}
	if strings.TrimSpace(path) == "" {
		settings.Name = "logs"
	}

	if strings.TrimSpace(name) == "" {
		settings.Name = "log"
	}

	setup(settings)
}

func setup(settings *settings) {
	dir := settings.Path
	fileName := fmt.Sprintf("%s-%s.%s",
		settings.Name,
		time.Now().Format(settings.TimeFormat),
		settings.Ext)

	logFile, err := file.MustOpen(fileName, dir)
	if err != nil {
		log.Fatalf("logging.setup err: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(mw, DefaultPrefix, log.LstdFlags)
}

func setPrefix(level Level) {
	_, f, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", levelFlags[level], filepath.Base(f), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
	printStack()
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	printStack()
	logger.Fatalln(v...)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	logger.Printf("ERROR STACK: %s\n", string(buf[:n]))
}
