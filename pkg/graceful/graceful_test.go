package graceful

import (
	"context"
	"log"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	count := 0
	var actualLog strings.Builder
	checkLog := ""
	expectedLog := "Stay in a loop\nStay in a loop\nSignal Sigterm...\nCancel Context\nBreak the loop\n"

	wait := make(chan struct{})
	outerCtx, cancelCtx := context.WithCancel(context.Background())
	action := Action{
		Start: func(c context.Context) {
			// max 10 seconds
			for i := 0; i < 10; i++ {
				exit := false
				select {
				case <-outerCtx.Done():
					actualLog.WriteString("Break the loop\n")
					checkLog += "Break the loop\n"
					exit = true
					wait <- struct{}{}
				case <-time.After(1 * time.Second):
					count += 1
					actualLog.WriteString("Stay in a loop\n")
					checkLog += "Stay in a loop\n"
				}

				if exit {
					break
				}
			}
		},
		Shutdown: func(c context.Context) {
			actualLog.WriteString("Cancel Context\n")
			cancelCtx()
		},
	}

	go Run(context.Background(),
		5*time.Second,
		map[string]Action{
			"PRINT LOOP": action,
		},
		StandardSignals...)

	// sleep for 2 seconds before doing sig int
	time.Sleep(2500 * time.Millisecond)

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Println(err)
	}

	// On a Unix-like system, pressing Ctrl+C on a keyboard sends a
	// SIGINT signal to the process of the program in execution.
	//
	// This example simulates that by sending a SIGINT signal to itself.
	if err := p.Signal(syscall.SIGTERM); err != nil {
		log.Println(err)
	}
	actualLog.WriteString("Signal Sigterm...\n")

	<-wait

	assert.Equal(t, expectedLog, actualLog.String())
}
