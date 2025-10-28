package model

import "github.com/JJApplication/Apollo/utils/json"

type SystemNetwork struct {
	ByteRecv    uint64 `json:"byteRecv"`
	ByteSent    uint64 `json:"byteSent"`
	PacketsRecv uint64 `json:"packetsRecv"`
	PacketsSent uint64 `json:"packetsSent"`
}

func (s *SystemNetwork) JSON() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return b
}
