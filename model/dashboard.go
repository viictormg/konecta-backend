package model

type OnCreateHackathonEvents struct {
	CreateAt string `json:"createdAt"`
	Event    string `json:"event"`
	ID       string `json:"id"`
}

type Data struct {
	OnCreateHackathonEvents OnCreateHackathonEvents `json:"onCreateHackathonEvents"`
}

type Payload struct {
	Data Data `json:"data"`
}

type DashboardReponse struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}
