package model

type Warn struct {
	Type    string `json:"type"`
	Date    string `json:"date"`
	TraceId string `json:"trace_id"`
}
