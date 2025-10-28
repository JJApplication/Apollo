package model

import "github.com/JJApplication/Apollo/utils/json"

type SystemLoad struct {
	Minute1  float64 `json:"minute1"`
	Minute5  float64 `json:"minute5"`
	Minute15 float64 `json:"minute15"`
}

func (s *SystemLoad) JSON() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return b
}
