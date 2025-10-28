package model

import "github.com/JJApplication/Apollo/utils/json"

type SystemCPU struct {
	Percent float64 `json:"percent"`
}

func (s *SystemCPU) JSON() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return b
}
