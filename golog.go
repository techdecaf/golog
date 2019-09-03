package golog

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Log logger
type Log struct {
	Name string
}

// LogMessage message structure
type LogMessage struct {
	Action  string
	Message interface{}
}

// Info displays information about the program durring runtime.
func (log *Log) Info(action string, text interface{}) string {
	return log.print(LogMessage{Action: action, Message: text}, false)
}

// Fatal logs to os.Stderr and exits 1
func (log *Log) Fatal(action string, err interface{}) string {
	return log.print(LogMessage{Action: action, Message: err}, true)
}

func (log *Log) print(m LogMessage, exit bool) string {
	msg := fmt.Sprintf("[%s][%s] %v", log.Name, m.Action, m.Message)

	printInfo := color.New(color.FgGreen)
	printError := color.New(color.FgRed)

	m.Action = strings.ToLower(m.Action)

	if exit {
		printError.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
	printInfo.Fprintln(os.Stdout, msg)

	return msg
}
