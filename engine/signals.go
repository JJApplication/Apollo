/*
Project: dirichlet signals.go
Created: 2022/2/18 by Landers
*/

package engine

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	Stop      = syscall.SIGUSR2
	ForceStop = syscall.SIGINT
	Reload    = syscall.SIGUSR1
	APPName   = "[Dirichlet]"
)

// RegisterSignals 监听部分信号使用
func RegisterSignals(s *http.Server, sigChan chan os.Signal) {
	signal.Notify(sigChan, Stop, ForceStop, Reload)

	for sig := range sigChan {
		switch sig {
		case Stop:
			fmt.Println("\n" + APPName + " wait for closing all connections 🚥")
			if err := s.Shutdown(context.Background()); err != nil {
				fmt.Println("\n" + APPName + " server close failed ❌")
			}
		case ForceStop:
			fmt.Println("\n" + APPName + " force closed by user 🔔")
			if err := s.Close(); err != nil {
				fmt.Println("\n" + APPName + " server close failed ❌")
			}
		case Reload:
			// todo
			fmt.Println("\n" + APPName + " server reloaded ☘️")
		}
	}
}
