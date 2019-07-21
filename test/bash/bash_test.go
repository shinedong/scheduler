package bash

import (
	"os"
	"os/exec"
	"testing"
)

const BashPath = "I:\\C++\\cygwin\\bin\\bash.exe"

func TestBash(t *testing.T) {

	command := exec.Command(BashPath, "-c", "pwd")
	file := os.Stdout
	//command.Run() // run command
	bytes, e := command.CombinedOutput()
	if e != nil {
		t.Fatalf("run command err:%v", e)
		return
	} else {
		t.Logf("执行成功=%s", string(bytes))
	}
	t.Log(file)
}
