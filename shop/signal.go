package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// WaitForSignal creates a channel and wait until desired signal
// arriving. It's a block call so be sure you're using it correctly.
func WaitForSignal(callback func()) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	sig := <-sigCh
	log.Printf("signal arrived, signal: %s", sig.String())

	callback()
}
