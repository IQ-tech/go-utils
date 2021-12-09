package utils

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"
)

func AtInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}

func EndAsErr(err error, message string, infow io.Writer, errorw io.Writer) {
	if err == nil {
		return
	}

	fmt.Fprintln(errorw, "Error:", err)
	fmt.Fprintln(infow, message)
	time.Sleep(time.Millisecond * 50) // needed for printing all messages before exiting
	os.Exit(1)
}
