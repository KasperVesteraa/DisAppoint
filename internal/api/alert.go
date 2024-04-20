package api

type Alert struct {
	Id       string `json:"id"`
	SendTime int    `json:"send_time"`
	Message  string `json:"message"`
}
