package feature

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"konecta/model"
	"net/http"
	"strings"
)

func (f Feature) GetDashboard() (model.DashboardReponse, error) {
	data := model.DashboardReponse{}

	token, err := f.Auth()

	if err != nil {
		return data, nil
	}

	url := "https://mqjl9s6vf4.execute-api.eu-west-1.amazonaws.com/prod/v1/hackday/private/event"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	a := fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", a)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
