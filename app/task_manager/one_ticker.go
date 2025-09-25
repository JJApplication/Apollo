package task_manager

type OneTicker struct {
	Ch         chan bool `json:",omitempty"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Des        string    `json:"des"`
	Stopped    bool      `json:"stopped"`
	CreateTime int64     `json:"create_time"`
	Duration   int       `json:"duration"`
	LastRun    int64     `json:"lastRun"`
}

type OneTickerRes struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Des        string `json:"des"`
	Stopped    bool   `json:"stopped"`
	CreateTime int64  `json:"create_time"`
	Duration   int    `json:"duration"`
	LastRun    int64  `json:"lastRun"`
}

func (tc *OneTicker) Start() (uuid string, err error) {
	tc.Stopped = false
	return tc.UUID, nil
}

func (tc *OneTicker) Stop() (uuid string, err error) {
	tc.Stopped = true
	tc.Ch <- true
	return tc.UUID, nil
}
