package model

type SystemInfo struct {
	Kernel   string `json:"kernel"`
	Platform string `json:"platform"`
	Family   string `json:"family"`
	Version  string `json:"version"`

	// CPU

	// Memory

	// Disk IO

	// NetWork
}
