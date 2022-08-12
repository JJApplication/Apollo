/*
Project: Apollo conn.go
Created: 2021/12/8 by Landers
*/

package uds

import (
	"fmt"
	"io"
	"net"
	"syscall"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
)

const MaxReadSize = 4096

func echo(c net.Conn) {
	logger.Logger.Info(fmt.Sprintf("client connected: [%s]", c.RemoteAddr().Network()))
	for {
		buf := make([]byte, MaxReadSize)
		cnt, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				logger.Logger.Info("read from UDS Client, client disconnect from server")
				break
			}
			logger.Logger.Info(fmt.Sprintf("read from UDS Client failed: %s", err.Error()))
			break
		}

		cmd := string(buf[:cnt])
		if supportCmds(cmd) {
			res := execCmd(cmd)
			c.Write(UDSResponse(res))
		} else {
			c.Write(UDSResponse(UDSRes{
				Error: ErrCmdNotFound,
				Data:  "",
			}))
		}
	}

	defer c.Close()
}

func removeSocket() {
	_ = syscall.Unlink(getSocket())
}

func getSocket() string {
	if config.ApolloConf.Server.Uds == "" {
		return socketAddr
	}
	return config.ApolloConf.Server.Uds
}

func listen() {
	removeSocket()
	addr, err := net.ResolveUnixAddr("unix", getSocket())

	l, err := net.ListenUnix("unix", addr)
	logger.Logger.Info(fmt.Sprintf("UDS Listen at: %s", getSocket()))
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
