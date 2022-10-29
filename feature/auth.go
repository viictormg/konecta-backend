package feature

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"konecta/model"
	"net/http"
)

func (f Feature) Auth() (model.RespnseShape, error) {
	data := model.RespnseShape{}

	url := "https://hackday-22-prod.auth.eu-west-1.amazoncognito.com/oauth2/token?grant_type=client_credentials"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return data, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic MTI5Y2lvZzZpaDRvam5sYWFmOGQzcmxvc3U6MTNydmliNzdjOGRhbWV2bmphdjNtZ2c1c2pkMTFpbzBucGs0MjJqY3RnamxhbmpqdnZsZQ==")
	req.Header.Add("Cookie", "XSRF-TOKEN=248aec81-5567-4f01-baee-1c1ac3ec7d25")

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

	json.Unmarshal(body, &data)

	return data, nil

}
