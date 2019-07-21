package bash

import (
	"context"
	"os/exec"
	"testing"
	"time"
)

type baseResult struct {
	err    error
	output []byte
}

func TestBatchBash(t *testing.T) {

	resultChan := make(chan *baseResult, 100)

	ctx, cancelFunc := context.WithCancel(context.TODO())

	go runCommand(ctx, resultChan)

	time.Sleep(time.Second * 1)

	cancelFunc()

	var res *baseResult
	res = <-resultChan

	t.Logf("执行结构:%v", string(res.output))
}

func runCommand(ctx context.Context, resultChan chan *baseResult) {
	var cmd *exec.Cmd = exec.CommandContext(ctx, BashPath, "-c", "sleep 2;echo hello;")
	bytes, e := cmd.CombinedOutput()
	resultChan <- &baseResult{
		err:    e,
		output: bytes,
	}
}
