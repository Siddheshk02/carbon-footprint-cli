package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Body struct {
	Energy   int    `json:"consumption"`
	Location string `json:"location"`
}

func Traditional(en int, loc string) {
	var t Body
	t.Location = loc
	t.Energy = en
	jsonReq, err := json.Marshal(t)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://app.trycarbonapi.com/api/traditionalHydro"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		panic(err)
	}

	apiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI0IiwianRpIjoiYWRlODAxNzlmZWNjNDkyMDIwN2IyOWZjZWIxZDBhMWMzZjg1NjRkMzE3YjEwZDYyMzU3MjdmNjk1YjY4ZjRmYjU5NmUxNTYxNTcwODY4MjUiLCJpYXQiOjE2Nzc0MDg3NTEsIm5iZiI6MTY3NzQwODc1MSwiZXhwIjoxNzA4OTQ0NzUxLCJzdWIiOiIzNTc0Iiwic2NvcGVzIjpbXX0.jjixNRW7oUKTAp26Q-k5tTLmPkK9EkNo1uPk-feIYCElCtrqx_-AfslTqwsXGbKnX7LkzCq06BiISD3sVUKhJQ"

	req.Header.Set("Authorization", "Bearer "+apiKey)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

type Car struct {
	Distance int    `json:"distance"`
	Vehicle  string `json:"vehicle"`
}

func CarTravel(dis int, veh string) {

	var myStruct Car
	myStruct.Distance = dis
	myStruct.Vehicle = veh

	jsonReq, err := json.Marshal(myStruct)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://app.trycarbonapi.com/api/carTravel"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		panic(err)
	}

	apiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI0IiwianRpIjoiYWRlODAxNzlmZWNjNDkyMDIwN2IyOWZjZWIxZDBhMWMzZjg1NjRkMzE3YjEwZDYyMzU3MjdmNjk1YjY4ZjRmYjU5NmUxNTYxNTcwODY4MjUiLCJpYXQiOjE2Nzc0MDg3NTEsIm5iZiI6MTY3NzQwODc1MSwiZXhwIjoxNzA4OTQ0NzUxLCJzdWIiOiIzNTc0Iiwic2NvcGVzIjpbXX0.jjixNRW7oUKTAp26Q-k5tTLmPkK9EkNo1uPk-feIYCElCtrqx_-AfslTqwsXGbKnX7LkzCq06BiISD3sVUKhJQ"

	req.Header.Set("Authorization", "Bearer "+apiKey)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
