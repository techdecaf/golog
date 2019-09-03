package golog

import "testing"

var log = Log{
	Name: "testing",
}

func TestLogInfo(t *testing.T) {
	want := "[testing][action] message"

	if got := log.Info("action", "message"); got != want {
		t.Errorf("log.Info() = %q, want %q", got, want)
	}
}

func TestLogPrint(t *testing.T) {
	want := "[testing][log.Print] exit"
	msg := LogMessage{
		Action:  "log.Print",
		Message: "exit",
	}
	if got := log.Print(msg); got != want {
		t.Errorf("log.Print() = %q, want %q", got, want)
	}
}
