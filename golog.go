package golog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Log logger
type Log struct {
	Name string
}

// LogMessage message structure
type LogMessage struct {
	Action   string
	Message  interface{}
	ExitCode int
}

// Info displays information about the program durring runtime.
func (log *Log) Info(action string, text interface{}) string {
	return log.Print(LogMessage{Action: action, Message: text})
}

// Fatal logs to os.Stderr and exits 1
func (log *Log) Fatal(action string, err interface{}) string {
	return log.Print(LogMessage{Action: action, Message: err, ExitCode: 1})
}

// Print is a helper function for log.Info and log.Fatal,
// but can be called directly if you want more control over the exit code.
func (log *Log) Print(m LogMessage) string {
	msg := fmt.Sprintf("[%s][%s] %v", log.Name, m.Action, m.Message)

	printInfo := color.New(color.FgGreen)
	printError := color.New(color.FgRed)

	if m.ExitCode > 0 {
		printError.Fprintln(os.Stderr, msg)
		os.Exit(m.ExitCode)
	}
	printInfo.Fprintln(os.Stdout, msg)

	return msg
}
