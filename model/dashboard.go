package model

import "time"

type AutoGenerated struct {
	Version    string        `json:"version"`
	ID         string        `json:"id"`
	DetailType string        `json:"detail-type"`
	Source     string        `json:"source"`
	Account    string        `json:"account"`
	Time       time.Time     `json:"time"`
	Region     string        `json:"region"`
	Resources  []interface{} `json:"resources"`
	Detail     struct {
		Events []struct {
			Version    string        `json:"version"`
			ID         string        `json:"id"`
			DetailType string        `json:"detail-type"`
			Source     string        `json:"source"`
			Account    string        `json:"account"`
			Time       int           `json:"time"`
			Region     string        `json:"region"`
			Resources  []interface{} `json:"resources"`
			Detail     struct {
				TopicName string `json:"topicName"`
				Version   string `json:"version"`
				EventBody struct {
					Service struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Users       []struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"users"`
					} `json:"service"`
					Data struct {
						Metrics []struct {
							Metric string `json:"metric"`
							Stats  struct {
								Count int `json:"count"`
							} `json:"stats"`
							UsersID   []string `json:"usersId,omitempty"`
							Qualifier string   `json:"qualifier,omitempty"`
						} `json:"metrics"`
					} `json:"data"`
				} `json:"eventBody"`
			} `json:"detail"`
		} `json:"events"`
	} `json:"detail"`
}

type OnCreateHackathonEvents struct {
	CreateAt   string        `json:"createdAt"`
	Event      AutoGenerated `json:"eventStruct,omitempty"`
	EventSring string        `json:"event"`
	ID         string        `json:"id"`
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
