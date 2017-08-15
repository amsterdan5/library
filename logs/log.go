package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//错误级别
const (
	L_INFO = iota
	L_WARING
	L_ERROR
	L_DEBUG
)

var logHandler *log.Logger

var logOption = newLogger()

var messLevel = [L_DEBUG + 1]string{"[I]", "[W]", "[E]", "[D]"}

const (
	L_FILENAME   = "log/record.log"
	L_LEVEL      = L_INFO | L_WARING | L_ERROR
	L_PREMISSION = 0666
	L_PREFIX     = ""
)

type Option struct {
	Filename       string
	Level          int32
	FilePermission os.FileMode
	Prefix         string
}

func GetOption() *Option {
	return logOption
}

func (o *Option) Init() {
	fileHandler, err := os.OpenFile(o.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, o.FilePermission)
	if err != nil {
		log.Println(err.Error())
	}
	logHandler = log.New(fileHandler, o.Prefix, log.Llongfile|log.Ldate|log.Ltime)
}

func newLogger() *Option {
	o := new(Option)
	o.Filename = L_FILENAME
	o.Level = L_LEVEL
	o.FilePermission = L_PREMISSION
	o.Prefix = L_PREFIX
	return o
}

// change log option
func (o *Option) SetLogger(jsonStr string) {
	err := json.Unmarshal([]byte(jsonStr), o)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("log", o)
}

func Info(message interface{}, v ...interface{}) {
	logOption.Info(message, v...)
}

func Waring(message interface{}, v ...interface{}) {
	logOption.Waring(message, v...)
}

func (o *Option) Info(message interface{}, v ...interface{}) {
	logOption.logMessage(L_INFO, message, v...)
}

func (o *Option) Waring(message interface{}, v ...interface{}) {
	logOption.logMessage(L_WARING, message, v...)
}

func (o *Option) logMessage(level int32, message interface{}, v ...interface{}) {
	mes := formatMess(message, v...)
	fmt.Println(mes)
	if level <= o.Level {
		logHandler.Println(messLevel[level] + mes)
	}
}

// format message
func formatMess(f interface{}, v ...interface{}) string {
	var message string
	switch f.(type) {
	case string:
		message = f.(string)
		if len(v) == 0 {
			return message
		}

		if strings.Contains(message, "%") && !strings.Contains(message, "%%") {

		} else {
			message += strings.Repeat(" %v", len(v))
		}
	default:
		message = fmt.Sprint(f)
		if len(v) == 0 {
			return message
		}
		message += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(message, v...)
}
