package advancedExec

import (
	"fmt"
	"testing"
)

func TestShellCommandExecution(t *testing.T) {
	result,_:=ExecuteShellCommand("ls","-al")
	fmt.Println(string(result))
}