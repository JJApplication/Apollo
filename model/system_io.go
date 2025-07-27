package model

import "encoding/json"

type SystemIO struct {
	TotalReadBytes  uint64 `json:"totalReadBytes"`
	TotalWriteBytes uint64 `json:"totalWriteBytes"`
	TotalReadTime   uint64 `json:"totalReadTime"`
	TotalWriteTime  uint64 `json:"totalWriteTime"`
	TotalReadCount  uint64 `json:"totalReadCount"`
	TotalWriteCount uint64 `json:"totalWriteCount"`
	IOTime          uint64 `json:"ioTime"`
}

func (s *SystemIO) JSON() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return b
}
