package main

type Response struct {
	Success  bool   `json:"success"`
	Msg      string `json:"msg"`
	Action   string `json:"action"`
	Pop      int    `json:"pop"`
	UserName string `json:"userName"`
	Location string `json:"location"`
}
