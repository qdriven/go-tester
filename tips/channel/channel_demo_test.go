package channel

import "testing"

func TestSendToChannelPanic(t *testing.T) {
	SendToChannelPanic()
}

func TestSendToChannel(t *testing.T) {
	SendToChannel()
}

func TestExecuteSum(t *testing.T) {
	ExecuteSum()
}

func TestExecFib(t *testing.T) {
	ExecFib(100)
}