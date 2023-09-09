/*
   Create: 2023/7/31
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package status_manager

type Status struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AppType     string `json:"app_type"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Data        struct {
		PID      string `json:"pid"`
		Port     []int  `json:"port"`
		Language string `json:"language"`
		Cpu      string `json:"cpu"`
		Mem      string `json:"mem"`
		Lifetime string `json:"lifetime"`
	} `json:"data"`
	Children []Status `json:"children"`
}

type StatusTree struct {
	UpdateTime string `json:"update_time"`
}
