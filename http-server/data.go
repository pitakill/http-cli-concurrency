package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://jsonplaceholder.cypress.io/users"
)

type person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func getDataFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if !(http.StatusOK <= response.StatusCode && response.StatusCode < http.StatusMultipleChoices) {
		return nil, errors.New("not a valid response from" + URL)
	}

	return ioutil.ReadAll(response.Body)
}

func getPeople() ([]person, error) {
	data, err := getDataFromURL(URL)
	if err != nil {
		return nil, err
	}

	people := make([]person, 0)
	if err = json.Unmarshal(data, &people); err != nil {
		return nil, err
	}

	return people, nil
}
