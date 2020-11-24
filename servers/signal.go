package servers

import (
	"os"
	"os/signal"
)

func notifyInterrupt() <-chan os.Signal {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	return sigchan
}
