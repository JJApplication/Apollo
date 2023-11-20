/*
Project: Apollo register.go
Created: 2021/12/8 by Landers
*/

package uds

const (
	socketAddr = "/var/run/Apollo.sock"
)

// Register 注册一个Unix domain socket通信
func Register() {
	go listen()
}

func Listen() {
	go run()
}
