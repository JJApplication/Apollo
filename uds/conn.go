/*
Project: dirichlet conn.go
Created: 2021/12/8 by Landers
*/

package uds

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/landers1037/dirichlet/logger"
)

func echo(c net.Conn) {
	logger.Logger.Info(fmt.Sprintf("client connected: [%s]", c.RemoteAddr().Network()))
	for {
		buf := make([]byte, 1024)
		cnt, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				logger.Logger.Info("read from UDS Client, client disconnect from server")
				break
			}
			logger.Logger.Info(fmt.Sprintf("read from UDS Client failed %s", err.Error()))
			break
		}

		cmd := string(buf[:cnt])
		if supportCmds(cmd) {
			res := execCmd(cmd)
			c.Write([]byte(res))
		} else {
			c.Write([]byte("cmd not support"))
		}
	}

	defer c.Close()
}

func removeSocket() {
	_ = os.Remove(sockerAddr)
}

func listen() {
	removeSocket()
	addr, err := net.ResolveUnixAddr("unix", sockerAddr)

	l, err := net.ListenUnix("unix", addr)
	logger.Logger.Info(fmt.Sprintf("UDS Listen at: %s", sockerAddr))
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("UDS Listen failed: %s", err.Error()))
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		go echo(conn)
	}
}
