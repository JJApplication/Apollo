package model

import "encoding/json"

type SystemMem struct {
	Percent   float64 `json:"percent"`
	Free      uint64  `json:"free"`
	Available uint64  `json:"available"`
	Used      uint64  `json:"used"`
	Cached    uint64  `json:"cached"`
}

func (s *SystemMem) JSON() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return b
}
