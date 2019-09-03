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
	Message string
}

// Info displays information about the program durring runtime.
func (log *Log) Info(action, text string) {
	log.print(LogMessage{Action: action, Message: text}, false)
}

// Fatal logs to os.Stderr and exits 1
func (log *Log) Fatal(action, text string) {
	log.print(LogMessage{Action: action, Message: text}, true)
}

func (log *Log) print(m LogMessage, exit bool) {
	msg := fmt.Sprintf("[%s][%s] %s", log.Name, m.Action, m.Message)

	printInfo := color.New(color.FgGreen)
	printError := color.New(color.FgRed)

	m.Action = strings.ToLower(m.Action)

	if exit {
		printError.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
	printInfo.Fprintln(os.Stdout, msg)
}
