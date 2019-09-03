package golog

import "testing"

var log = Log{
	Name: "testing",
}

func TestHello(t *testing.T) {
	want := "[testing][action] message"

	if got := log.Info("action", "message"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
