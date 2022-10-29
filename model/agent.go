package model

type Agent struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ResponseAgents struct {
	Agents []Agent `json:"agents"`
	Count  int64   `json:""`
}

type Stats struct {
	Count int64 `json:"count"`
}

type MetricStruct struct {
	Metric string   `json:"metric"`
	Stats  Stats    `json:"stats"`
	Users  []string `json:"users"`
}
