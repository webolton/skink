package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/webolton/skink/internal/initializer"
	"github.com/webolton/skink/internal/skink"
)

func main() {
	// create context channel and invoke cancel for background routine TODO: can I do this in one step?
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// create channel to listen for os signals
	signalChan := make(chan os.Signal, 1)
	// configure signalChannel to listen for SIGTERM/SIGINT and SIGHUP
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	// initialize config
	config := &initializer.Config{}

	// initialize and call anonymous function to cancel the main context if a termination is sent
	// on the signal channel
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	// invoke and call go routine
	go func() {
		for { // start infinite loop
			select {
			case sig := <-signalChan: // while listening to our signalChannel, assign var to channel
				switch sig {
				case syscall.SIGHUP: // if SIGHUP, reinitialize program
					config.Initialize(os.Args)
				case os.Interrupt: // if an interupt, cancel context and exit with an error
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done(): // while listening to the context, if it is done exit
				os.Exit(1)
			}
		}
	}()

	// while routine is listening for signals, run the main skink program
	if err := skink.Run(ctx, os.Stdout, config); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
